package main

import (
	"fmt"
	"github.com/GalaIO/jsonrpc2"
)

type Result struct {
	Content string
}

// go 的jsonrpc是1.0 版本 不支持2.0，主要是参数列表不同 还有批量请求
// 重新实现2.0版本
func main() {
	client, err := jsonrpc2.Dial("tcp", "89.23.35.9:50002")
	if err != nil {
		panic(err)
	}

	result := make([]byte, 0, 1024)
	err = client.Call("server.version", []string{"1.9.5", "1.4"}, &result)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))

	err = client.Call("blockchain.block.header", []int{1, 0}, &result)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))

	client.Close()
}
