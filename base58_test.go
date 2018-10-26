package base58

import (
	"encoding/hex"
	"testing"
)

func TestEncodeToString(t *testing.T) {
	testCases := []struct {
		in  string
		out string
		err error
	}{
		{
			"",
			"",
			nil,
		},
		{
			"00",
			"1",
			nil,
		},
		{
			"00010966776006953d5567439e5e39f86a0d273beed61967f6",
			"16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM",
			nil,
		},
	}

	b58 := NewBitcoinBase58()

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			inBytes, err := hex.DecodeString(tc.in)
			if err != nil {
				t.Fatal("can not decode input string")
			}

			out, err := b58.EncodeToString(inBytes)
			if err != tc.err {
				t.Errorf("expected: %v, actual: %v", tc.err, err)
			}
			if out != tc.out {
				t.Errorf("expected: %s, actual: %s", tc.out, out)
			}
		})
	}
}

func TestDecodeString(t *testing.T) {
	testCases := []struct {
		in  string
		out string
		err error
	}{
		{
			"",
			"",
			nil,
		},
		{
			"1",
			"00",
			nil,
		},
		{
			"16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM",
			"00010966776006953d5567439e5e39f86a0d273beed61967f6",
			nil,
		},
		{
			"0",
			"",
			ErrInvalidChar,
		},
	}

	b58 := NewBitcoinBase58()

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			outBytes, err := b58.DecodeString(tc.in)
			if err != tc.err {
				t.Errorf("expected: %v, actual: %v", tc.err, err)
			}

			out := hex.EncodeToString(outBytes)
			if out != tc.out {
				t.Errorf("expected: %s, actual: %s", tc.out, out)
			}
		})
	}
}
