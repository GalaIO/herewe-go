package main

import (
	"net"
	"fmt"
)

func main() {
	tcpServer, err := net.Listen("tcp", ":9980")
	if err != nil {
		panic(err)
	}

	shutdown := make(chan bool, 1)
	
	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			fmt.Println("listen err, ", err)
		}

		go func(conn net.Conn) {
			output := make([]byte, 0, 10)
			for {
				buf := make([]byte, 10)
				n, e := conn.Read(buf)
				if n <= 0 || e != nil {
					fmt.Println("n, e", n, e)
					break
				}
				fmt.Println("n, buf", n, string(buf))
				output = append(output, buf...)
			}
			fmt.Println("get=", string(output))
		}(conn)

		select {
		case <-shutdown:
			fmt.Println("shutdown...")
			return
		default:
		}
	}
}
