package main

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"github.com/btcsuite/btcutil/base58"
	"fmt"
	"encoding/hex"
)

func scriptAddr(rawScript []byte) string {
	shaone := sha256.Sum256(rawScript)
	hash160 := ripemd160.New()
	hash160.Write(shaone[:])
	rawAddr := hash160.Sum(nil)
	final := make([]byte, 0, 25)
	final = append(final, 0x05)
	final = append(final, rawAddr...)
	checksumonce := sha256.Sum256(final)
	checksumtwice := sha256.Sum256(checksumonce[:])
	checksum := checksumtwice[:4]
	final = append(final, checksum...)
	return base58.Encode(final)
}

func main() {

	// test 1
	// OP_DUP 0x76
	// OP_HASH160 0xa9
	// OP_EQUALVERIFY 0x88
	// OP_CHECKSIG	 0xac
	// dup hash160 [89abcdefabbaabbaabbaabbaabbaabbaabbaabba] equalverify checksig
	// 0x76 0xa9 0x14 89abcdefabbaabbaabbaabbaabbaabbaabbaabba 0x88 0xac
	//scriptDecode := []byte{0x76, 0xa9, 0x14}
	//bs, _ := hex.DecodeString("89abcdefabbaabbaabbaabbaabbaabbaabbaabba")
	//scriptDecode = append(scriptDecode, bs...)
	//scriptDecode = append(scriptDecode, 0x88)
	//scriptDecode = append(scriptDecode, 0xac)
	//fmt.Println(scriptAddr(scriptDecode))

	bs, _ := hex.DecodeString("5221020ae29f86f404e4b302cfa17ff15d93149af6a54c80a4172d47e41f55f6a78d732103664d528eb80096671ef9011c533ceb5df133238e3690d88f2960c786398b86b121029a449ea4a2155ea10002d704604bb3e8606631d35af20889a74b82b2dab572f6210321602d78046d63256b1730b119b1aca3428039f18fdb73ccf45ad3e148dd9b1754ae")
	fmt.Println(scriptAddr(bs))
	fmt.Printf("%x, %x", 111, 196)
}
