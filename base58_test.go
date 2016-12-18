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

var encodeTestCases = []encodeTestCase{
	{"00010966776006953D5567439E5E39F86A0D273BEED61967F6", "16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM", nil},
}

func TestEncodeToString(t *testing.T) {
	for _, testCase := range encodeTestCases {
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
