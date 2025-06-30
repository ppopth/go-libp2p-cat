package cat

import (
	"context"
	"crypto/rand"
	"fmt"
	"maps"
	"math/big"
	"sync"

	"github.com/ppopth/p2p-broadcast/pb"
	"github.com/ppopth/p2p-broadcast/pubsub"

	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/peer"
)

var log = logging.Logger("cat")

type CatOption func(*CatRouter) error

// NewCat returns a new CatRouter object.
func NewCat(opts ...CatOption) (*CatRouter, error) {
	c := &CatRouter{}

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
		MaxCoefficient: *big.NewInt(255),
		Prime:          *p, // 2^256+297 (the first prime having more than 256 bits)
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

// CatRouter is a router that implements the CAT protocol.
type CatRouter struct {
	lk sync.Mutex

	params CatParams
	peers  map[peer.ID]pubsub.TopicSendFunc
}

type CatParams struct {
	// Chunk size in bytes. When the message is larger than this size, it will be chunked, so
	// every chunk except the last one will be of this size.
	ChunkSize int
	// The number of field elements per chunk.
	ElementsPerChunk int
	// Max coeficient used in linear combinations.
	MaxCoefficient big.Int
	// Prime number of the prime field used in linear combinations.
	Prime big.Int
}

func (c *CatRouter) Publish(buf []byte) error {
	if len(buf)%c.params.ChunkSize != 0 {
		return fmt.Errorf("the size of the message (%d) must be a multiple of the chunk size (%d)",
			len(buf), c.params.ChunkSize)
	}

	// Divide into chunks
	var chunks [][]byte
	for len(buf) > 0 {
		chunk := buf[:c.params.ChunkSize]
		chunks = append(chunks, chunk)
		buf = buf[c.params.ChunkSize:]
	}

	c.lk.Lock()
	peers := maps.Clone(c.peers)
	c.lk.Unlock()

	for pid, sendFunc := range peers {
		// do linear combination for each peer
		chunked, coeffs, err := c.combine(chunks)
		if err != nil {
			log.Warnf("error publishing to %s; %s", pid, err)
			// Skipping to the next peer
			continue
		}
		// send the combined chunk to that peer
		chunk := &pb.CatRpc_Chunk{
			Data: chunked,
		}
		for _, coeff := range coeffs {
			chunk.Coefficients = append(chunk.Coefficients, coeff.Bytes())
		}
		sendFunc(&pb.TopicRpc{
			Cat: &pb.CatRpc{
				Chunks: []*pb.CatRpc_Chunk{chunk},
			},
		})
	}
	return nil
}

// combine does linear combination on chunks and return the combined chunk and its coefficients.
func (c *CatRouter) combine(chunks [][]byte) ([]byte, []*big.Int, error) {
	var coeffs []*big.Int
	acc := make([]*big.Int, c.params.ElementsPerChunk)
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
		coeff, err := rand.Int(rand.Reader, &c.params.MaxCoefficient)
		if err != nil {
			return nil, nil, err
		}
		coeffs = append(coeffs, coeff)
		for i, elem := range v {
			acc[i].Add(acc[i], new(big.Int).Mul(coeff, elem)) // acc[i] += coeff * elem
		}
	}

	combined, err := bigIntsToBytes(acc, bitsPerElement)
	if err != nil {
		// There is supposed to be no errors
		panic(err)
	}
	return combined, coeffs, nil
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
