package main

import (
	"math/rand"
	"fmt"
	"encoding/base64"
	"crypto/elliptic"
	"math/big"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	// 安全随机数，使用rand.Read()
	priKey := make([]byte, 32)
	if _, err := rand.Read(priKey); err!=nil {
		fmt.Println("rand error")
	}
	// 生成私钥
	fmt.Println(base64.StdEncoding.EncodeToString(priKey))

	// 生成公钥 指定椭圆曲线的生成点 大素数p
	p, _ := new(big.Int).SetString("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 0)
	n, _ := new(big.Int).SetString("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 0)
	b, _ := new(big.Int).SetString("0x0000000000000000000000000000000000000000000000000000000000000007", 0)
	gx, _ := new(big.Int).SetString("0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 0)
	gy, _ := new(big.Int).SetString("0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 0)
	curveParam := elliptic.CurveParams{
		p,
		n,
		b,
		gx,
		gy,
		256,
		"curve param",
	}
	px, py := curveParam.ScalarBaseMult(priKey)
	fmt.Printf("pub x: %s, y:%s\r\n", base64.StdEncoding.EncodeToString(px.Bytes()), base64.StdEncoding.EncodeToString(py.Bytes()))

	// 生成短公钥 base58(version + RIPEMD160(SHA256(pub)) + SHA256(SHA256(version+RIPEMD160(SHA256(pub))))[:4])
	pubs := append(px.Bytes(), py.Bytes()...)
	sha3 := sha256.New()
	sha3.Write(pubs)
	hashed := sha3.Sum(nil)
	ripemd160.New()
	pkgData := append([]byte("1"), hashed...)
	sha3.Reset()
	checksum := sha3.Sum(pkgData)[:4]
	//shortPubs := base58Encode(append(pkgData, checksum...))
	shortPubs := append(pkgData, checksum...)
	fmt.Println("私钥：" + base64.StdEncoding.EncodeToString(shortPubs))

}
