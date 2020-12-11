package main

import "fmt"

type tFunc func()

func main() {
	var h tFunc
	fmt.Println(h)
}
