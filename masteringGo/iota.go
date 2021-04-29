package main

import (
	"fmt"
	"strconv"
)

func main() {
	// go支持iota来辅助生成常量，在const作用域内
	// 同时iota是以顺序增1的方式生成常量，也支持表达式，来生成自己需要的值
	const (
		C1 = iota // 0 起始值为0
		C2 // 1
		C3 // 2
	)
	fmt.Println(C1, C2, C3)

	const (
		C4 = 1 << iota // 1
		C5 // 2
		_  // 4 这个被省略，4不会使用
		C6 // 8
	)
	fmt.Println(C4, C5, C6)

	// 如果作为枚举替代的话，可以给iota生成的值设置类型
	const (
		C7 ErrCode = 100000 + iota // 100000
		C8 // 100001
		_ // 100002
		_ // 100003
		C9 // 100004
	)
	fmt.Println(C7, C8, C9)
}

type ErrCode int
func (e ErrCode) String() string {
	return "{\"ErrCode\": " + strconv.Itoa(int(e)) + "}"
}
