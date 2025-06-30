package cat

import (
	"encoding/hex"
	"math/big"
	"testing"
)

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
