package main

import (
	"fmt"
	"encoding/base64"
	crypto2 "github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"math/big"
	"crypto/ecdsa"
)

// 使用以太坊的椭圆加密系数和生成元（起始点）
func main() {

	// 生成私钥
	privKey, _ := crypto2.GenerateKey()
	// 私钥
	fmt.Println(base64.StdEncoding.EncodeToString(privKey.D.Bytes()))
	// 公钥
	pubAddress := crypto2.PubkeyToAddress(privKey.PublicKey)
	fmt.Printf("pub x: %s, y:%s\r\n", base64.StdEncoding.EncodeToString(privKey.X.Bytes()), base64.StdEncoding.EncodeToString(privKey.Y.Bytes()))
	// 生成短公钥 base58(version + RIPEMD160(SHA256(pub)) + SHA256(SHA256(version+RIPEMD160(SHA256(pub))))[:4])
	fmt.Println("短公钥：" + pubAddress.String())

	// 私钥签名 并验证
	sha3 := sha3.NewLegacyKeccak256()
	sha3.Reset()
	sha3.Write([]byte("hello world~~~"))
	msgHash := sha3.Sum(nil)
	fmt.Printf("hash: %s\n", base64.StdEncoding.EncodeToString(msgHash))

	// 签名后有r s v，r s是签名后结果，v是标识公钥在x轴上方或者下方，用来恢复公钥
	sigs, _ := crypto2.Sign(msgHash, privKey)
	fmt.Printf("sig, R:%s, S:%s, V:%d\n", base64.StdEncoding.EncodeToString(sigs[:32]), base64.StdEncoding.EncodeToString(sigs[32:64]), sigs[64])
	// 利用签名恢复公钥
	recPubs, _ := crypto2.Ecrecover(msgHash, sigs)
	fmt.Printf("recpub compressflag:%d, x: %s, y:%s\r\n", recPubs[0], base64.StdEncoding.EncodeToString(recPubs[1:33]), base64.StdEncoding.EncodeToString(recPubs[33:]))

	// 恢复的公钥打上非压缩标识，摘要剔除最后的恢复标识
	// 非压缩公钥 用第一个额外的字节0x04表示，随后是32字节的公钥x轴，32字节的y轴
	// 压缩公钥 用字节0x02表示y轴为偶数 0x03表示y轴为奇数
	fmt.Println(crypto2.VerifySignature(recPubs, msgHash, sigs[:64]))

	//利用压缩公钥验证
	recPubLeys, _ := crypto2.SigToPub(msgHash, sigs)
	fmt.Println(crypto2.VerifySignature(crypto2.CompressPubkey(recPubLeys), msgHash, sigs[:64]))

	// 使用ecdsa工具验证摘要
	var r, s big.Int
	r.SetBytes(sigs[:32])
	s.SetBytes(sigs[32:64])
	fmt.Println(ecdsa.Verify(recPubLeys, msgHash, &r, &s))
}
