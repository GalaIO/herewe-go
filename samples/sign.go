package main

import (
	"crypto/rsa"
	"crypto/rand"
	"fmt"
	"crypto/md5"
	"encoding/base64"
)

// 编程实现公约私钥加解密过程
func crypto() {
	// 创建私钥
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	fmt.Println("系统输出的私钥", key)

	// 得到公钥
	pub := key.PublicKey
	fmt.Println("系统输出的公钥", pub)

	// 源文件
	source := []byte("hello world")
	cipherTest, _ := rsa.EncryptOAEP(md5.New(), rand.Reader, &pub, source, nil)
	fmt.Println(base64.StdEncoding.EncodeToString(cipherTest))

	plainText, _ := rsa.DecryptOAEP(md5.New(), rand.Reader, key, cipherTest, nil)
	fmt.Println(string(plainText))

}

func main() {
	crypto()
}
