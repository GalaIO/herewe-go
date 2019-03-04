package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// 加密
func AESEncrypt(source []byte, key []byte) []byte {
	// 校验秘钥
	block, _ := aes.NewCipher(key)
	// 对明文补码
	source = PKCS7Padding(source, block.BlockSize())
	// 选择加密方式
	cbcMode := cipher.NewCBCEncrypter(block, key)

	// 申请缓冲区, 使用cbc模式加密
	encrypted := make([]byte, len(source))
	cbcMode.CryptBlocks(encrypted, source)
	return encrypted
}

// 解密
func AESDecrypt(source []byte, key []byte) []byte {
	// 校验秘钥
	block, _ := aes.NewCipher(key)
	// 选择加密方式
	cbcMode := cipher.NewCBCDecrypter(block, key)

	// 申请缓冲区, 使用cbc模式解密
	decrypted := make([]byte, len(source))
	cbcMode.CryptBlocks(decrypted, source)

	// 对解密后去码
	return PKCS7UnPadding(decrypted)
}

// 对加密后密文去码，通过字节最后byte来判断
func PKCS7UnPadding(source []byte) []byte {
	size := len(source)
	pad := int(source[size-1])
	return source[:size-pad]
}

// 对byte进行补码，满足对size长度分组均分
// PKCS5Padding 算法指对8位固定长度补码
func PKCS7Padding(source []byte, size int) []byte {
	// 计算补码位数
	sourceSize := len(source)
	padCount := size - sourceSize%size
	// 申请新内存，并copy
	dst := make([]byte, padCount)

	padByte := byte(padCount)
	for index:=0; index<padCount; index++ {
		dst[index] = padByte
	}
	return append(source, dst...)
}

func main() {
	//fmt.Println(PKCS7Padding([]byte("1111"), 10))
	//fmt.Println(base64.StdEncoding.EncodeToString(PKCS7Padding([]byte("1111"), 10)))
	key := "aeF34H-qwsDEsaDe"
	encrypt := AESEncrypt([]byte("hello world"), []byte(key))
	fmt.Println(encrypt)
	fmt.Println(string(AESDecrypt(encrypt, []byte(key))))
}
