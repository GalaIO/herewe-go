package main

import "fmt"

type Msg struct {
}

func main() {
	var s *Msg
	var i interface{} = s
	// true false
	fmt.Println(s == nil, i == nil)
}
