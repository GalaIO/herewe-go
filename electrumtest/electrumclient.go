package main

import (
	"net"
	"fmt"
	bytes2 "bytes"
)

func main() {
	conn, e := net.Dial("tcp", "127.0.0.1:50002")
	if e != nil {
		panic(e)
	}

	// 换行符 必须有
	jsonRpcMsg := "{\"jsonrpc\": \"2.0\", \"method\": \"%s\", \"params\": %s, \"id\": %d}\r\n"
	counter := 0
	msg := fmt.Sprintf(jsonRpcMsg, "server.version", "[ \"1.9.5\", \"0.4\" ]", counter)
	conn.Write([]byte(msg))

	bytes := make([]byte, 0, 1024)
	buffer := make([]byte, 128)
	for _, err := conn.Read(buffer); err==nil; {
		bytes = append(bytes, buffer...)
		index := bytes2.Index(bytes, []byte("\n"))
		if index >= 0 {
			bytes = bytes[:index+1]
			conn.Close()
			break
		}
	}
	fmt.Println(string(bytes))

	conn.Close()
}
