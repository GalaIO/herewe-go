package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9980")
	if err != nil {
		panic(err)
	}

	n, err := conn.Write([]byte("hello world,hello world,hello world,hello world,hello world"))

	fmt.Println("n, err, ", n, err)

	conn.Close()
}
