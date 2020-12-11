package main

import (
	"fmt"
)

type worker interface {
	work()
}

type person struct {
	name string
	worker
}

func main() {
	var w worker = person{name: "hello"}
	fmt.Println(w)
}
