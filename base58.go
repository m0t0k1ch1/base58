package base58

import (
	"bytes"
	"errors"
	"math/big"
)

const (
	BitcoinBase58Chars = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	Base               = 58
)

var (
	div  = big.NewInt(Base)
	zero = big.NewInt(0)
)

var (
	ErrInvalidLengthBytes = errors.New("invalid length bytes")
)

type Base58 struct {
	chars [58]byte
}

func NewBase58(s string) (*Base58, error) {
	b58 := &Base58{}
	if err := b58.setChars(s); err != nil {
		return nil, err
	}

	return b58, nil
}

func NewBitcoinBase58() *Base58 {
	b58, _ := NewBase58(BitcoinBase58Chars)

	return b58
}

func (b58 *Base58) setChars(s string) error {
	if len(s) != 58 {
		return ErrInvalidLengthBytes
	}

	chars := []byte(s)
	copy(b58.chars[:], chars[:])

	return nil
}

func (b58 *Base58) EncodeToString(srcBytes []byte) (string, error) {
	n := &big.Int{}
	n.SetBytes(srcBytes)

	buf := &bytes.Buffer{}
	for _, srcByte := range srcBytes {
		if srcByte == 0x00 {
			if err := buf.WriteByte(b58.chars[0]); err != nil {
				return "", err
			}
		} else {
			break
		}
	}

	mod := &big.Int{}

	tmpBuf := &bytes.Buffer{}
	for {
		if n.Cmp(zero) == 0 {
			tmpBytes := tmpBuf.Bytes()
			tmpBytesLen := len(tmpBytes)
			for i := 1; i <= tmpBytesLen; i++ {
				buf.WriteByte(tmpBytes[tmpBytesLen-i])
			}
			return buf.String(), nil
		}

		n.DivMod(n, div, mod)
		if err := tmpBuf.WriteByte(b58.chars[mod.Int64()]); err != nil {
			return "", err
		}
	}
}
