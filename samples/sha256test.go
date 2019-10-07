package main

import (
	"encoding/hex"
	"crypto/sha256"
	"fmt"
)

func main() {
	raw := "0100000000000000000000000000000000000000000000000000000000000000000000003ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a29ab5f49ffff001d1dac2b7c"

	bs, _ := hex.DecodeString(raw)
	onece := sha256.Sum256(bs)
	twice := sha256.Sum256(onece[:])
	twiceBs := twice[:]
	fmt.Println(hex.EncodeToString(twiceBs))
}

func reserveBytes(bs []byte) {
	for i, j := 0, len(bs)-1; i < j; {
		bs[i], bs[j] = bs[j], bs[i]
		i ++
		j --
	}
}
