package main

import (
	"crypto/sha256"
	"fmt"
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
)

func main() {
	sec, _ := hex.DecodeString("801e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd")
	shaonce := sha256.Sum256(sec)
	shatwice := sha256.Sum256(shaonce[:])
	final := make([]byte, 0, len(sec)+33)
	//final = append(final, 0x80)
	final = append(final, sec...)
	final = append(final, shatwice[:4]...)
	fmt.Printf("wif raw: %x\nchecksum: %d\n", final, uint32(shatwice[0])<<24+uint32(shatwice[1])<<16+uint32(shatwice[2])<<8+uint32(shatwice[3]))
	fmt.Println(base58.Encode(final))
}
