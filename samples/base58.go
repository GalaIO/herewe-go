package main

import "fmt"

const encodeBase58Std = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func main() {
	str := "hello worlduibuh1u021392038/.,;'[]"
	fmt.Println(base58Encode(append(make([]byte, 10), []byte(str)...)))
	fmt.Println(base58Decode([]byte(base58Encode(append(make([]byte, 10), []byte(str)...)))))
	//fmt.Println(base64.StdEncoding.EncodeToString([]byte(str)))
}
func base58Encode(src []byte) string {
	// 记录0值数量
	zeros := 0
	for src[zeros] == 0x00 {
		zeros++
	}
	// 推断目标存储大小，最多不超过2倍
	dstLen := len(src)*2 - zeros
	dst := make([]byte, dstLen)
	i, j := zeros, dstLen - 1
	for i < len(src) {
		dst[j] = encodeBase58Std[divmode(src, i, 256, 58)]
		j--
		if src[i] == 0 {
			i++
		}
	}
	// 忽略多余的0
	for dst[j+1] == 0x00 {
		j ++
	}
	for zeros > 0 {
		dst[j] = encodeBase58Std[0]
		zeros--
		j--
	}
	return string(dst[j+1:])
}

// 在原址计算，会破坏原slice
func divmode(src []byte, i int, base uint, mod uint) byte {
	var remainer uint = 0
	for ; i < len(src); i++ {
		temp := remainer*base + uint(src[i])
		src[i] = byte(temp / mod)
		remainer = temp % mod
	}
	return byte(remainer)
}

func base58Decode(src []byte) string {
	// 记录0值数量
	zeros := 0
	for src[zeros] == '1' {
		zeros++
	}
	// 推断目标存储大小，最多不超过1倍
	dstLen := len(src)
	dst := make([]byte, dstLen)
	i, j := zeros, dstLen - 1

	for k := 0; k < len(src); k++ {
		src[k] = reserveBase58(src[k])
	}
	for i < len(src) {
		dst[j] = divmode(src, i, 58, 256)
		j--
		if src[i] == 0 {
			i++
		}
	}
	for dst[j+1] == 0x00 {
		j ++
	}
	return string(dst[j-zeros+1:])
}
func reserveBase58(ch byte) byte {
	if ch >= '1' && ch <= '9'{
		return ch - '1'
	}else if ch >= 'A' && ch <= 'H' {
		return ch - 'A' + 9
	}else if ch >= 'J' && ch <= 'N' {
		return ch - 'J' + 17
	}else if ch >= 'P' && ch <= 'Z' {
		return ch - 'P' + 22
	}else if ch >= 'a' && ch <= 'k' {
		return ch - 'a' + 33
	}else if ch >= 'm' && ch <= 'z' {
		return ch - 'm' + 44
	}
	return 0
}
