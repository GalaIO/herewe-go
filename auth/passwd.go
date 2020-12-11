package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"strings"
)

func main() {
	passwdHashTest()

}

func passwdHashTest() {
	src := "hello"
	passwdHash := scryptHash(src, 16)
	fmt.Println("passwd hash", passwdHash)
	fmt.Println(strings.EqualFold(passwdHash, scryptHash(src, 16)))
	fmt.Println(strings.EqualFold(passwdHash, scryptHash("234", 16)))
	appKey := scryptHash(src, 8)
	fmt.Println("appkey:", appKey)
	fmt.Println("appsecret:", scryptHash(appKey, 16))
}

var passwdSalt = []byte("W@!*YYSAqw44")
func scryptHash(passwd string, keyLen int) string {
	bytes, err := scrypt.Key([]byte(passwd), passwdSalt, 16384, 8, 1, keyLen)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
