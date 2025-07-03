package cat

import (
	"context"
	crand "crypto/rand"
	"fmt"
	"maps"
	"math/big"
	mrand "math/rand"
	"slices"
	"sync"

	"github.com/ppopth/p2p-broadcast/pb"
	"github.com/ppopth/p2p-broadcast/pubsub"

	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/peer"
)

var log = logging.Logger("cat")

type MsgIdFunc func([]byte) string

type CatOption func(*CatRouter) error

// NewCat returns a new CatRouter object.
func NewCat(idFunc MsgIdFunc, opts ...CatOption) (*CatRouter, error) {
	c := &CatRouter{
		idFunc: idFunc,

		peers:  make(map[peer.ID]pubsub.TopicSendFunc),
		chunks: make(map[string][]Chunk),
	}

	err := WithCatParams(DefaultCatParams())(c)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// DefaultCatParams returns the default CAT parameters as a config.
func DefaultCatParams() CatParams {
	p := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
	p.Add(p, big.NewInt(297))

	params := CatParams{
		ChunkSize:      1024, // 1KB
		MaxCoefficient: big.NewInt(255),
		Prime:          p, // 2^256+297 (the first prime having more than 256 bits)
		Fanout:         40,
	}
	params.ElementsPerChunk = 8 * params.ChunkSize / (p.BitLen() - 1)

	return params
}

// WithCatParams is a router option that allows a custom config to be set when instantiating the CAT router.
func WithCatParams(params CatParams) CatOption {
	return func(c *CatRouter) error {
		if (8*params.ChunkSize)%params.ElementsPerChunk != 0 {
			return fmt.Errorf("8*ElementsPerChunk (%d) must divide ChunkSize (%d)",
				8*params.ElementsPerChunk, params.ChunkSize)
		}
		if params.Prime.BitLen()-1 > (8*params.ChunkSize)/params.ElementsPerChunk {
			return fmt.Errorf("The length of Prime must be greater than (8*ChunkSize)/ElementsPerChunk (%d)",
				(8*params.ChunkSize)/params.ElementsPerChunk)
		}
		c.params = params
		return nil
	}
}

type Chunk struct {
	Data   []byte
	Coeffs []*big.Int
}

// CatRouter is a router that implements the CAT protocol.
type CatRouter struct {
	lk sync.Mutex

	params CatParams
	idFunc MsgIdFunc

	peers  map[peer.ID]pubsub.TopicSendFunc
	chunks map[string][]Chunk
}

type CatParams struct {
	// Chunk size in bytes. When the message is larger than this size, it will be chunked, so
	// every chunk except the last one will be of this size.
	ChunkSize int
	// The number of field elements per chunk.
	ElementsPerChunk int
	// Max coeficient used in linear combinations.
	MaxCoefficient *big.Int
	// Prime number of the prime field used in linear combinations.
	Prime *big.Int
	// The number of chunks sent out in total, when publish.
	Fanout int
}

func (c *CatRouter) Publish(buf []byte) error {
	if len(buf)%c.params.ChunkSize != 0 {
		return fmt.Errorf("the size of the message (%d) must be a multiple of the chunk size (%d)",
			len(buf), c.params.ChunkSize)
	}

	mid := c.idFunc(buf)

	// Divide into chunks
	var chunks [][]byte
	for len(buf) > 0 {
		chunk := buf[:c.params.ChunkSize]
		chunks = append(chunks, chunk)
		buf = buf[c.params.ChunkSize:]
	}

	c.lk.Lock()
	sendFuncs := slices.Collect(maps.Values(c.peers))
	// Shuffle the peers
	mrand.Shuffle(len(sendFuncs), func(i, j int) {
		sendFuncs[i], sendFuncs[j] = sendFuncs[j], sendFuncs[i]
	})
	c.lk.Unlock()

	sendFuncIndex := 0
	// Do a round robin on peers to send the chunks
	for i := 0; i < c.params.Fanout; i++ {
		// do linear combination for each peer
		combined, err := c.combine(chunks)
		if err != nil {
			log.Warnf("error publishing; %s", err)
			// Skipping to the next peer
			continue
		}
		// send the combined chunk to that peer
		chunk := &pb.CatRpc_Chunk{
			MessageID: &mid,
			Data:      combined.Data,
		}
		for _, coeff := range combined.Coeffs {
			chunk.Coefficients = append(chunk.Coefficients, coeff.Bytes())
		}
		sendFuncs[sendFuncIndex](&pb.TopicRpc{
			Cat: &pb.CatRpc{
				Chunks: []*pb.CatRpc_Chunk{chunk},
			},
		})

		sendFuncIndex++
		sendFuncIndex %= len(sendFuncs)
	}
	return nil
}

// combine does linear combination on chunks and return the combined chunk and its coefficients.
func (c *CatRouter) combine(chunks [][]byte) (Chunk, error) {
	var coeffs []*big.Int
	acc := make([]*big.Int, c.params.ElementsPerChunk)
	for i := range acc {
		acc[i] = new(big.Int)
	}
	bitsPerElement := 8 * c.params.ChunkSize / c.params.ElementsPerChunk

	for _, chunk := range chunks {
		v, err := splitBitsToBigInts(chunk, bitsPerElement)
		if err != nil {
			// There is supposed to be no errors
			panic(err)
		}
		if len(v) != len(acc) {
			// There is supposed to be no errors
			panic("splitBitsToBigInts returned a wrong vector")
		}
		// Randomize a coefficient
		coeff, err := crand.Int(crand.Reader, c.params.MaxCoefficient)
		if err != nil {
			return Chunk{}, err
		}
		coeffs = append(coeffs, coeff)
		for i, elem := range v {
			acc[i].Mod(new(big.Int).Add(acc[i], new(big.Int).Mul(coeff, elem)), c.params.Prime) // acc[i] = (acc[i] + coeff * elem) % Prime
		}
	}

	combined, err := bigIntsToBytes(acc, bitsPerElement)
	if err != nil {
		// There is supposed to be no errors
		panic(err)
	}
	chunk := Chunk{
		Data:   combined,
		Coeffs: coeffs,
	}
	return chunk, nil
}

func (c *CatRouter) Next(ctx context.Context) ([]byte, error) {
	return nil, nil
}

func (c *CatRouter) AddPeer(p peer.ID, sendFunc pubsub.TopicSendFunc) {
	c.lk.Lock()
	defer c.lk.Unlock()

	c.peers[p] = sendFunc
}

func (c *CatRouter) RemovePeer(p peer.ID) {
	c.lk.Lock()
	defer c.lk.Unlock()

	delete(c.peers, p)
}

func (c *CatRouter) HandleIncomingRPC(p peer.ID, trpc *pb.TopicRpc) {
	c.lk.Lock()
	defer c.lk.Unlock()

	catRpc := trpc.GetCat()
	// Remember chunks
	for _, chunk := range catRpc.GetChunks() {
		mid := chunk.GetMessageID()
		data := chunk.GetData()
		var coeffs []*big.Int
		for _, buf := range chunk.GetCoefficients() {
			coeffs = append(coeffs, new(big.Int).SetBytes(buf))
		}
		if data != nil {
			c.chunks[mid] = append(c.chunks[mid], Chunk{
				Data:   data,
				Coeffs: coeffs,
			})
		}
	}
}

func (c *CatRouter) Close() error {
	return nil
}

// splitBitsToBigInts splits a byte slice into big.Ints each with exactly k bits. (credit to ChatGPT)
func splitBitsToBigInts(data []byte, k int) ([]*big.Int, error) {
	if k <= 0 {
		return nil, fmt.Errorf("k must be positive")
	}

	totalBits := len(data) * 8
	if totalBits%k != 0 {
		return nil, fmt.Errorf("total number of bits (%d) not divisible by k (%d)", totalBits, k)
	}

	numChunks := totalBits / k
	result := make([]*big.Int, 0, numChunks)

	bitIndex := 0
	for i := 0; i < numChunks; i++ {
		val := big.NewInt(0)
		for j := 0; j < k; j++ {
			byteIndex := bitIndex / 8
			bitOffset := 7 - (bitIndex % 8) // Big-endian bit order

			bit := (data[byteIndex] >> bitOffset) & 1
			val.Lsh(val, 1)
			if bit == 1 {
				val.Or(val, big.NewInt(1))
			}

			bitIndex++
		}
		result = append(result, val)
	}

	return result, nil
}

// bigIntsToBytes is an inverse of splitBitsToBigInts. (credit to ChatGPT)
func bigIntsToBytes(chunks []*big.Int, k int) ([]byte, error) {
	if k <= 0 {
		return nil, fmt.Errorf("k must be positive")
	}

	totalBits := len(chunks) * k
	if totalBits%8 != 0 {
		return nil, fmt.Errorf("total number of bits (%d) is not divisible by 8", totalBits)
	}

	numBytes := totalBits / 8
	result := make([]byte, numBytes)

	bitIndex := 0
	for _, chunk := range chunks {
		for i := k - 1; i >= 0; i-- {
			bit := chunk.Bit(i)

			byteIndex := bitIndex / 8
			bitOffset := 7 - (bitIndex % 8)

			if bit == 1 {
				result[byteIndex] |= 1 << bitOffset
			}
			bitIndex++
		}
	}

	return result, nil
}

// invertMatrix computes the inverse of an n x n matrix over F_p using Gaussian elimination. (credit to ChatGPT)
func invertMatrix(A [][]*big.Int, p *big.Int) ([][]*big.Int, error) {
	n := len(A)
	inv := make([][]*big.Int, n)
	for i := range inv {
		inv[i] = make([]*big.Int, n)
		for j := range inv[i] {
			if i == j {
				inv[i][j] = big.NewInt(1)
			} else {
				inv[i][j] = big.NewInt(0)
			}
		}
	}

	// Make a deep copy of A to work on
	B := make([][]*big.Int, n)
	for i := range A {
		B[i] = make([]*big.Int, n)
		for j := range A[i] {
			B[i][j] = new(big.Int).Set(A[i][j])
		}
	}

	for i := 0; i < n; i++ {
		// Find pivot
		invPivot := new(big.Int).ModInverse(B[i][i], p)
		if invPivot == nil {
			return nil, fmt.Errorf("matrix not invertible")
		}
		for j := 0; j < n; j++ {
			B[i][j].Mul(B[i][j], invPivot).Mod(B[i][j], p)
			inv[i][j].Mul(inv[i][j], invPivot).Mod(inv[i][j], p)
		}
		// Eliminate other rows
		for k := 0; k < n; k++ {
			if k == i {
				continue
			}
			factor := new(big.Int).Set(B[k][i])
			for j := 0; j < n; j++ {
				tmp := new(big.Int).Mul(factor, B[i][j])
				B[k][j].Sub(B[k][j], tmp).Mod(B[k][j], p)

				tmp2 := new(big.Int).Mul(factor, inv[i][j])
				inv[k][j].Sub(inv[k][j], tmp2).Mod(inv[k][j], p)
			}
		}
	}
	return inv, nil
}

// recoverVectors solves V = A⁻¹ * R, where A is the coefficient matrix and R the combined vectors. (credit to ChatGPT)
func recoverVectors(A [][]*big.Int, R [][]*big.Int, p *big.Int) ([][]*big.Int, error) {
	n := len(A)
	m := len(R[0])
	Ainv, err := invertMatrix(A, p)
	if err != nil {
		return nil, err
	}

	V := make([][]*big.Int, n)
	for i := range V {
		V[i] = make([]*big.Int, m)
		for j := 0; j < m; j++ {
			V[i][j] = big.NewInt(0)
			for k := 0; k < n; k++ {
				tmp := new(big.Int).Mul(Ainv[i][k], R[k][j])
				V[i][j].Add(V[i][j], tmp).Mod(V[i][j], p)
			}
		}
	}
	return V, nil
}
