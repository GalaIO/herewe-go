package main

import (
	"time"
	"fmt"
)

type ProxyHandler struct {
	baseHandler BaseHandler
}

func NewProxyHandler() BaseHandler {
	return &ProxyHandler{
		baseHandler: &printHandler{},
	}
}

func (p *ProxyHandler) Handler(person *Person) {
	stratTime := time.Now()
	p.baseHandler.Handler(person)
	fmt.Println("rt:", time.Now().Sub(stratTime))
}

func main() {

	person := &Person{
		Name: "xiaoguo",
	}
	handler := NewProxyHandler()
	handler.Handler(person)
}
