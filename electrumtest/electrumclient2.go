package main

import (
	"net/rpc/jsonrpc"
	"fmt"
)

type Result struct {
	Content string
}

// go 的jsonrpc是1.0 版本 不支持2.0，主要是参数列表不同 还有批量请求
// 重新实现2.0版本
func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:50002")
	if err != nil {
		panic(err)
	}

	result := []interface{}{}
	err = client.Call("server.version", struct{a string; b string}{"1.9.5", "0.6"}, &result)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	err = client.Call("blockchain.block.header", []interface{}{1, 0}, &result)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	client.Close()
}
