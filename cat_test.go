package cat

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"net/netip"
	"slices"
	"testing"
	"time"

	"github.com/ppopth/p2p-broadcast/host"
	"github.com/ppopth/p2p-broadcast/pubsub"
)

func hashSha256(buf []byte) string {
	h := sha256.New()
	h.Write(buf)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func TestCatPublish(t *testing.T) {
	hosts := getHosts(t, 5)
	psubs := getPubsubs(t, hosts)

	var topics []*pubsub.Topic
	var subs []*pubsub.Subscription
	var routers []*CatRouter

	for _, ps := range psubs {
		c, err := NewCat(hashSha256,
			WithCatParams(CatParams{
				ChunkSize:        8,
				ElementsPerChunk: 2,
				MaxCoefficient:   big.NewInt(255),
				Prime:            big.NewInt(4_294_967_311),
				Fanout:           16,
			}),
		)
		if err != nil {
			t.Fatal(err)
		}
		tp, err := ps.Join("foobar", c)
		if err != nil {
			t.Fatal(err)
		}
		sub, err := tp.Subscribe()
		if err != nil {
			t.Fatal(err)
		}
		topics = append(topics, tp)
		subs = append(subs, sub)
		routers = append(routers, c)
	}

	for _, h := range hosts[1:] {
		if err := hosts[0].Connect(context.Background(), h.LocalAddr()); err != nil {
			t.Fatal(err)
		}
	}

	time.Sleep(200 * time.Millisecond)
	// Publish a message
	msg := []byte("cold-bird-jump-fog-grid-sand-pen")
	mid := hashSha256(msg)
	if err := topics[0].Publish(msg); err != nil {
		t.Fatal(err)
	}
	time.Sleep(200 * time.Millisecond)
	for i, router := range routers[1:] {
		if len(router.chunks) != 1 {
			t.Fatalf("a router %d should have received one message id; received %d", i, len(router.chunks))
		}
		if _, ok := router.chunks[mid]; !ok {
			t.Fatalf("a router %d didn't receive the expected messag id", i)
		}
		if len(router.chunks[mid]) != 4 {
			t.Fatalf("a router %d should have received %d chunks; received %d", i, 4, len(router.chunks[mid]))
		}
		bitsPerElement := 8 * router.params.ChunkSize / router.params.ElementsPerChunk
		/*
		   To recover the original vectors v₁, ..., vₙ from n random linear combinations r₁, ..., rₙ over a prime field F_p, we assume
		   the random linear combinations are linearly independent, i.e., the matrix of coefficients used to produce the combinations
		   is invertible. So:

		   Let V be an n × m matrix where each row is a vector vᵢ.

		   Let A be an n × n matrix of random coefficients.

		   Let R = A · V, where each row of R is rᵢ.

		   To recover V, compute A⁻¹ · R.
		*/
		var A [][]*big.Int
		var R [][]*big.Int
		for _, chunk := range router.chunks[mid] {
			r, err := splitBitsToBigInts(chunk.Data, bitsPerElement)
			if err != nil {
				t.Fatal(err)
			}
			A = append(A, chunk.Coeffs)
			R = append(R, r)
		}
		V, err := recoverVectors(A, R, router.params.Prime)
		if err != nil {
			t.Fatal(err)
		}

		var buf []byte
		for _, v := range V {
			chunk, err := bigIntsToBytes(v, bitsPerElement)
			if err != nil {
				t.Fatal(err)
			}
			buf = slices.Concat(buf, chunk)
		}
		if !slices.Equal(buf, msg) {
			t.Fatalf("a router %d should have received the recovered message of %s; received %s instead", i, string(msg), string(buf))
		}
	}
}

// (credit to ChatGPT)
func TestSplitAndCombineBits(t *testing.T) {
	tests := []struct {
		name     string
		hexData  string   // hex string of input bytes
		k        int      // bit chunk size
		expected []string // expected big.Ints in decimal
	}{
		{
			name:     "Simple 4-bit split",
			hexData:  "abcdef",
			k:        4,
			expected: []string{"10", "11", "12", "13", "14", "15"},
		},
		{
			name:     "8-bit split",
			hexData:  "010203",
			k:        8,
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "12-bit split",
			hexData:  "abcd12",
			k:        12,
			expected: []string{"2748", "3346"},
		},
		{
			name:     "Empty slice",
			hexData:  "",
			k:        8,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, _ := hex.DecodeString(tt.hexData)

			// Split
			chunks, err := splitBitsToBigInts(data, tt.k)
			if err != nil {
				t.Fatalf("splitBitsToBigInts failed: %v", err)
			}

			if len(chunks) != len(tt.expected) {
				t.Fatalf("expected %d chunks, got %d", len(tt.expected), len(chunks))
			}

			for i, exp := range tt.expected {
				if chunks[i].String() != exp {
					t.Errorf("chunk %d: expected %s, got %s", i, exp, chunks[i].String())
				}
			}

			// Recombine
			recombined, err := bigIntsToBytes(chunks, tt.k)
			if err != nil {
				t.Fatalf("bigIntsToBytes failed: %v", err)
			}

			if hex.EncodeToString(recombined) != tt.hexData {
				t.Errorf("recombined bytes don't match original: expected %s, got %s", tt.hexData, hex.EncodeToString(recombined))
			}
		})
	}
}

// (credit to ChatGPT)
func TestSplitBitsToBigInts_Errors(t *testing.T) {
	// k <= 0
	_, err := splitBitsToBigInts([]byte{0xFF}, -1)
	if err == nil {
		t.Error("expected error for k <= 0")
	}
	_, err = splitBitsToBigInts([]byte{0xFF}, 0)
	if err == nil {
		t.Error("expected error for k <= 0")
	}

	// data bits not divisible by k
	_, err = splitBitsToBigInts([]byte{0xFF}, 6)
	if err == nil {
		t.Error("expected error when data bits not divisible by k")
	}
}

// (credit to ChatGPT)
func TestBigIntsToBytes_Errors(t *testing.T) {
	// k <= 0
	_, err := bigIntsToBytes([]*big.Int{big.NewInt(1)}, 0)
	if err == nil {
		t.Error("expected error for k <= 0")
	}

	// total bits not divisible by 8
	_, err = bigIntsToBytes([]*big.Int{big.NewInt(1), big.NewInt(2)}, 5) // 10 bits
	if err == nil {
		t.Error("expected error for non-byte-aligned total bits")
	}
}

func getHosts(t *testing.T, n int) []*host.Host {
	var hs []*host.Host

	for i := 0; i < n; i++ {
		h, err := host.NewHost(host.WithAddrPort(netip.MustParseAddrPort("127.0.0.1:0")))
		if err != nil {
			t.Fatal(err)
		}
		hs = append(hs, h)
	}
	return hs
}

func getPubsubs(t *testing.T, hs []*host.Host) []*pubsub.PubSub {
	var psubs []*pubsub.PubSub
	for _, h := range hs {
		ps, err := pubsub.NewPubSub(h)
		if err != nil {
			t.Fatal(err)
		}
		psubs = append(psubs, ps)
	}
	return psubs
}

// Test the matrix operations

var p = big.NewInt(101) // small prime field for test consistency

func identityMatrix(n int, p *big.Int) [][]*big.Int {
	I := make([][]*big.Int, n)
	for i := 0; i < n; i++ {
		I[i] = make([]*big.Int, n)
		for j := 0; j < n; j++ {
			if i == j {
				I[i][j] = big.NewInt(1)
			} else {
				I[i][j] = big.NewInt(0)
			}
		}
	}
	return I
}

func matrixMultiply(A, B [][]*big.Int, p *big.Int) [][]*big.Int {
	n, m := len(A), len(B[0])
	k := len(B)
	C := make([][]*big.Int, n)
	for i := range C {
		C[i] = make([]*big.Int, m)
		for j := 0; j < m; j++ {
			sum := big.NewInt(0)
			for l := 0; l < k; l++ {
				tmp := new(big.Int).Mul(A[i][l], B[l][j])
				sum.Add(sum, tmp)
			}
			C[i][j] = new(big.Int).Mod(sum, p)
		}
	}
	return C
}

func matricesEqual(A, B [][]*big.Int) bool {
	for i := range A {
		for j := range A[i] {
			if A[i][j].Cmp(B[i][j]) != 0 {
				return false
			}
		}
	}
	return true
}

func generateRandomMatrix(n, m int, p *big.Int) [][]*big.Int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	M := make([][]*big.Int, n)
	for i := range M {
		M[i] = make([]*big.Int, m)
		for j := range M[i] {
			M[i][j] = new(big.Int).Rand(rng, p)
		}
	}
	return M
}

// (credit to ChatGPT)
func TestInvertMatrixIdentity(t *testing.T) {
	n := 4
	I := identityMatrix(n, p)

	Ainv, err := invertMatrix(I, p)
	if err != nil {
		t.Fatalf("Failed to invert identity matrix: %v", err)
	}
	if !matricesEqual(Ainv, I) {
		t.Errorf("Inverse of identity matrix should be identity")
	}
}

// (credit to ChatGPT)
func TestInvertMatrixSmall(t *testing.T) {
	// 1x1 invertible
	A := [][]*big.Int{{big.NewInt(3)}}
	Ainv, err := invertMatrix(A, p)
	if err != nil {
		t.Errorf("1x1 matrix invert failed: %v", err)
	}
	expect := new(big.Int).ModInverse(big.NewInt(3), p)
	if Ainv[0][0].Cmp(expect) != 0 {
		t.Errorf("Expected %v got %v", expect, Ainv[0][0])
	}

	// 2x2 invertible
	A = [][]*big.Int{
		{big.NewInt(4), big.NewInt(7)},
		{big.NewInt(2), big.NewInt(6)},
	}
	Ainv, err = invertMatrix(A, p)
	if err != nil {
		t.Fatalf("2x2 matrix invert failed: %v", err)
	}
	// Check A * Ainv == I
	I := matrixMultiply(A, Ainv, p)
	if !matricesEqual(I, identityMatrix(2, p)) {
		t.Errorf("2x2 invert failed: A * Ainv != I")
	}
}

// (credit to ChatGPT)
func TestInvertMatrixSingular(t *testing.T) {
	// Singular matrix (rows are linearly dependent)
	A := [][]*big.Int{
		{big.NewInt(1), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(4)},
	}
	_, err := invertMatrix(A, p)
	if err == nil {
		t.Errorf("Expected error for singular matrix")
	}
}

// (credit to ChatGPT)
func TestRecoverVectorsIdentity(t *testing.T) {
	// A is identity, R should equal V
	n := 3
	V := [][]*big.Int{
		{big.NewInt(1), big.NewInt(2)},
		{big.NewInt(3), big.NewInt(4)},
		{big.NewInt(5), big.NewInt(6)},
	}
	A := identityMatrix(n, p)
	R := matrixMultiply(A, V, p)

	Vrec, err := recoverVectors(A, R, p)
	if err != nil {
		t.Fatalf("recoverVectors failed: %v", err)
	}
	if !matricesEqual(V, Vrec) {
		t.Errorf("Expected recovered = original, got %v", Vrec)
	}
}

// (credit to ChatGPT)
func TestRecoverVectorsKnown(t *testing.T) {
	// A: 2x2, V: 2x2
	A := [][]*big.Int{
		{big.NewInt(2), big.NewInt(3)},
		{big.NewInt(1), big.NewInt(4)},
	}
	V := [][]*big.Int{
		{big.NewInt(5), big.NewInt(6)},
		{big.NewInt(7), big.NewInt(8)},
	}
	R := matrixMultiply(A, V, p)
	Vrec, err := recoverVectors(A, R, p)
	if err != nil {
		t.Fatalf("recoverVectors failed: %v", err)
	}
	if !matricesEqual(V, Vrec) {
		t.Errorf("Expected recovered = original, got %v", Vrec)
	}
}

// (credit to ChatGPT)
func TestRecoverVectorsSingular(t *testing.T) {
	// A is singular (non-invertible), should fail
	A := [][]*big.Int{
		{big.NewInt(1), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(4)},
	}
	V := [][]*big.Int{
		{big.NewInt(3)},
		{big.NewInt(5)},
	}
	R := matrixMultiply(A, V, p)
	_, err := recoverVectors(A, R, p)
	if err == nil {
		t.Errorf("Expected error for singular matrix in recoverVectors")
	}
}

// (credit to ChatGPT)
func TestRecoverVectorsRandom(t *testing.T) {
	n, m := 5, 3
	V := generateRandomMatrix(n, m, p)

	var A [][]*big.Int
	var err error
	for {
		A = generateRandomMatrix(n, n, p)
		_, err = invertMatrix(A, p)
		if err == nil {
			break
		}
	}
	R := matrixMultiply(A, V, p)

	Vrec, err := recoverVectors(A, R, p)
	if err != nil {
		t.Fatalf("recoverVectors failed: %v", err)
	}
	if !matricesEqual(V, Vrec) {
		t.Errorf("Recovered matrix did not match original")
	}
}
