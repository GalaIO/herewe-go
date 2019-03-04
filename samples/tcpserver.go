package main

import (
	"net"
	"bytes"
	"fmt"
)

func main() {
	// 监听端口
	listener, _ := net.Listen("tcp", "127.0.0.1:12345")
	defer listener.Close()

	// 循环检查客户端连接
	for {
		fmt.Println("wait for connection....")
		// 阻塞等待客户端连接
		conn, _ := listener.Accept()
		fmt.Printf("conn from %s\r\n", conn.RemoteAddr().String())
		// 创建缓存接收数据
		buf := make([]byte, 10)
		buffer := bytes.Buffer{}
		count, _ := conn.Read(buf)
		buffer.Write(buf)
		for count >= len(buf) {
			count, _ = conn.Read(buf)
			buffer.Write(buf)
		}
		fmt.Println("reader: "+ buffer.String())
		// echo 并关闭
		conn.Write(buffer.Bytes())
		conn.Close()
	}
}
