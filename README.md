# base58

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/base58?status.svg)](https://godoc.org/github.com/m0t0k1ch1/base58) [![wercker status](https://app.wercker.com/status/43ee805196ba2483d58fee224adfa649/s/master "wercker status")](https://app.wercker.com/project/byKey/43ee805196ba2483d58fee224adfa649) [![codecov](https://codecov.io/gh/m0t0k1ch1/base58/branch/master/graph/badge.svg)](https://codecov.io/gh/m0t0k1ch1/base58)

a BASE58 encoder/decoder for Go

``` sh
$ go get github.com/m0t0k1ch1/base58
```

## Example

``` go
package main

import (
	"encoding/hex"
	"log"

	"github.com/m0t0k1ch1/base58"
)

func main() {
	b58 := base58.NewBitcoinBase58()

	validPkh := "00010966776006953d5567439e5e39f86a0d273beed61967f6"
	validAddress := "16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM"

	validPkhBytes, err := hex.DecodeString(validPkh)
	if err != nil {
		log.Fatal(err)
	}

	address, err := b58.EncodeToString(validPkhBytes)
	if err != nil {
		log.Fatal(err)
	}
	if address == validAddress {
		log.Println("valid encoding")
	}

	pkhBytes, err := b58.DecodeString(address)
	if err != nil {
		log.Fatal(err)
	}
	if hex.EncodeToString(pkhBytes) == validPkh {
		log.Println("valid decoding")
	}
}
```
