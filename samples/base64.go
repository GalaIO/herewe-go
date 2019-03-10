package main

import (
	"fmt"
)

// 64编码是把每3个byte变成4个6bit，
// 其实是对64进制的取巧，因为3个byte可以表示为原数的256^3进制，然后该进制每位转换为64^4进制
// 由于256 64都是2的幂次所以转换非常简单，不需要取余计算，直接按位切分即可
// 想象一下吧一个普通的bytes转换为64进制，可以把bytes比作256进制的数，比如0x12 0x23 0x2a
// (0x12 * 256 + 0x23) * 256 + 0x2a 变成64进制很简单，依次取余即可，或者利用底层2进制特性
// 每3位256进制取成4位64进制，更通用的解法是因为64进制比256要小，所以最多一位256进制可以拆为两位64进制
// 由于除不尽，余数*256合并下一位继续取余，直至除尽，同样该方法通用可以用于base58，但没有之前算法性能优异
// (a * b) % p = (a % p * b % p) % p = ((a/p+a%p)*b)%p

const (
	encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	encodePad = '='
)

func base64Encode(src []byte ) string {
	n := len(src) / 3
	if len(src) % 3 != 0 {
		n ++
	}
	dst := make([]byte, 4 * n)
	i, j := 0, 0
	for i < len(src) / 3 * 3 {
		val := uint(src[i])<<16 | uint(src[i+1])<<8 | uint(src[i+2])
		dst[j] = encodeStd[uint8(val>>18 & 0x3f)]
		dst[j+1] = encodeStd[uint8(val>>12 & 0x3f)]
		dst[j+2] = encodeStd[uint8(val>>6 & 0x3f)]
		dst[j+3] = encodeStd[uint8(val & 0x3f)]
		i += 3
		j += 4
	}

	// 处理填充字符
	switch len(src)%3 {
	case 1:
		val := uint(src[i])<<16
		dst[j] = encodeStd[uint8(val>>18 & 0x3f)]
		dst[j+1] = encodeStd[uint8(val>>12 & 0x3f)]
		dst[j+2] = encodePad
		dst[j+3] = encodePad
	case 2:
		val := uint(src[i])<<16 | uint(src[i+1])<<8
		dst[j] = encodeStd[uint8(val>>18 & 0x3f)]
		dst[j+1] = encodeStd[uint8(val>>12 & 0x3f)]
		dst[j+2] = encodeStd[uint8(val>>6 & 0x3f)]
		dst[j+3] = encodePad
	}
	return string(dst)
}

func base64Decode(src []byte) string {
	srcLen := len(src)
	if srcLen < 4 || srcLen% 4 != 0 {
		panic("输入的无效base64编码")
	}
	// 计算需要申请的内存
	dstLen := srcLen / 4 * 3
	n := 0
	if src[srcLen-1] == encodePad && src[srcLen-2] == encodePad {
		n = 2
	}else if src[srcLen-1] == encodePad {
		n = 1
	}
	dstLen -= n
	dst := make([]byte, dstLen)
	i, j := 0, 0
	for i<srcLen-4 {
		val := uint(reserve(src[i])) << 18 | uint(reserve(src[i+1])) <<12| uint(reserve(src[i+2])) << 6 | uint(reserve(src[i+3]))
		dst[j] = uint8(val>>16 & 0xFF)
		dst[j+1] = uint8(val>>8 & 0xFF)
		dst[j+2] = uint8(val & 0xFF)
		i += 4
		j += 3
	}
	switch n {
	case 1:
		val := uint(reserve(src[i])) << 18 | uint(reserve(src[i+1])) <<12 | uint(reserve(src[i+2])) << 6
		dst[j] = uint8(val>>16 & 0xFF)
		dst[j+1] = uint8(val>>8 & 0xFF)
	case 2:
		val := uint(reserve(src[i])) << 18 | uint(reserve(src[i+1])) <<12
		dst[j] = uint8(val>>16 & 0xFF)
	}
	return string(dst)
}

// 从base64 推到索引值
func reserve(ch byte) byte {
	if ch == '+' {
		return 62
	}else if ch == '/' {
		return 63
	}else if ch >= '0' && ch <= '9' {
		return ch - '0' + 52
	}else if ch >= 'a' && ch <= 'z' {
		return ch - 'a' + 26
	}
	return ch - 'A'
}

func main() {

	str := "hello worlduibuh1u021392038/.,;'[]"
	fmt.Println(base64Encode([]byte(str)))
	fmt.Println(base64Decode([]byte(base64Encode([]byte(str)))))
	//fmt.Println(base64.StdEncoding.EncodeToString([]byte(str)))
}
