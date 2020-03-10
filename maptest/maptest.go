package main

import (
	"errors"
	"fmt"
	errors2 "github.com/pkg/errors"
)

var e1 = errors.New("1")
var e2 = errors2.WithMessage(e1, "2")
var m = map[interface{}]interface{}{}

func init() {

	m[e1] = "1"
	m[e2] = "2"
}

func main() {
	fmt.Println(m[e1])
	fmt.Println(m[e2])
}
