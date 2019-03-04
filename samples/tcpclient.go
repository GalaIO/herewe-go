package main

import (
	"net"
	"fmt"
)

func main() {

	// 构建服务器连接
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:12345")
	for {
		conn, _ := net.DialTCP("tcp", nil, addr)

		conn.Write([]byte("hello nidayede , nizai ganma"))
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println("result: " + string(buf[:n]))
	}
}
