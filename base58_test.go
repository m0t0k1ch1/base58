package base58

import (
	"encoding/hex"
	"reflect"
	"testing"
)

type encodeTestCase struct {
	in  string
	out string
	err error
}

type decodeTestCase struct {
	in  string
	out string
	err error
}

var (
	encodeTestCasesBitcoin = []encodeTestCase{
		{"", "", nil},
		{"00", "1", nil},
		{"00010966776006953d5567439e5e39f86a0d273beed61967f6", "16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM", nil},
	}
	decodeTestCasesBitcoin = []decodeTestCase{
		{"", "", nil},
		{"0", "", ErrInvalidChar},
		{"1", "00", nil},
		{"16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM", "00010966776006953d5567439e5e39f86a0d273beed61967f6", nil},
	}
)

func TestEncodeToString(t *testing.T) {
	for _, testCase := range encodeTestCasesBitcoin {
		inBytes, err := hex.DecodeString(testCase.in)
		if err != nil {
			t.Fatal("can not decode input string")
		}

		b58 := NewBitcoinBase58()

		out, err := b58.EncodeToString(inBytes)
		if !reflect.DeepEqual(err, testCase.err) {
			t.Errorf(`
invalid error
- got: %v
- expected: %v
`, err, testCase.err)
		}
		if out != testCase.out {
			t.Errorf(`
invalid encoding
- got: %s
- expected: %s
`, out, testCase.out)
		}
	}
}

func TestDecodeString(t *testing.T) {
	for _, testCase := range decodeTestCasesBitcoin {
		b58 := NewBitcoinBase58()

		outBytes, err := b58.DecodeString(testCase.in)
		if !reflect.DeepEqual(err, testCase.err) {
			t.Errorf(`
invalid error
- got: %v
- expected: %v
`, err, testCase.err)
		}

		out := hex.EncodeToString(outBytes)
		if out != testCase.out {
			t.Errorf(`
invalid decoding
- got: %s
- expected: %s
`, out, testCase.out)
		}
	}
}
