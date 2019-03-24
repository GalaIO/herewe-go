package main

import (
	"fmt"
	"encoding/base64"
	"crypto/elliptic"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"crypto/ecdsa"
	"crypto/rand"
)

func main() {

	// 使用go提供的椭圆加密系数
	p256Curve := elliptic.P256()

	// 生成私钥
	privKey, _ := ecdsa.GenerateKey(p256Curve, rand.Reader)

	// 私钥
	fmt.Println(base64.StdEncoding.EncodeToString(privKey.D.Bytes()))
	// 公钥
	fmt.Printf("pub x: %s, y:%s\r\n", base64.StdEncoding.EncodeToString(privKey.X.Bytes()), base64.StdEncoding.EncodeToString(privKey.Y.Bytes()))

	// 生成短公钥 base58(version + RIPEMD160(SHA256(pub)) + SHA256(SHA256(version+RIPEMD160(SHA256(pub))))[:4])
	pubs := append(privKey.X.Bytes(), privKey.Y.Bytes()...)
	sha3 := sha256.New()
	sha3.Write(pubs)
	hashedPubs := sha3.Sum(nil)
	hash160 := ripemd160.New()
	hash160.Write(hashedPubs)
	midPart := hash160.Sum(nil)
	pkgData := append([]byte("1"), midPart...)
	// 对前两部分 两次hash
	sha3.Reset()
	sha3.Write(pkgData)
	temp := sha3.Sum(nil)
	sha3.Reset()
	sha3.Write(temp)
	checksum := sha3.Sum(nil)[:4]
	shortPubs := append(pkgData, checksum...)

	// 使用base64打印 而不是base58
	fmt.Println("短公钥：" + base64.StdEncoding.EncodeToString(shortPubs))

	// 私钥签名 并验证
	sha3.Reset()
	sha3.Write([]byte("hello world~~~"))
	msgHash := sha3.Sum(nil)
	// 说明 https://blog.csdn.net/chenmo187J3X1/article/details/80910447
	R, S, _ := ecdsa.Sign(rand.Reader, privKey, msgHash)
	fmt.Println(ecdsa.Verify(&privKey.PublicKey, msgHash, R, S))

}
