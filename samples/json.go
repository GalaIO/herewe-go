package main

import (
	"fmt"
	"encoding/json"
)

type Message struct {
	Content string
	Tips string
}

func main() {

	msg := Message{"hello", "world"}
	data, _ := json.Marshal(&msg)
	fmt.Println(string(data))
	fmt.Println(json.Marshal([]byte{1, 2, 3}))
	fmt.Println(json.Marshal(123456))

}
