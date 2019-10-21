package main

import "fmt"

type HelloService interface {
	Hi() string
	Nihao() string
}

type HelloServiceImpl struct {
}

func (h *HelloServiceImpl) Hi() string {
	return "hello~~"
}

func (h *HelloServiceImpl) Nihao() string {
	return "你好~~"
}

func main() {
	h := HelloServiceImpl{}

	fmt.Println(h.Hi())
	fmt.Println(h.Nihao())
}
