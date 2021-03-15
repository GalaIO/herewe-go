package main

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Notify() error
}

// 接口的嵌入其实就是接口的组合，需要支持多个接口和组合
type Receiver interface {
	Notifier
	Receive() error
}

// 嵌入实际上是一种类型复用，复用字段、方法等等，结构体有一个匿名嵌入字段
// 在结构体调用被嵌入方法时(不管结构体还是接口)都会被提升到结构体上，可以直接调用
// 也可以使用u.Notifier.Notify() 执行匿名嵌入进行调用
// 当结构体重新实现相同方法或者接口时，会优先调用重新实现的方法，不会提升匿名嵌入字段
// 同时如果嵌入结构体 默认是结构体零值，如果嵌入接口默认为nil
type User struct {
	Notifier
	Name string
}

type User2 struct {
	Receiver
	Name string
}

func (u *User2) Notify() error {
	return errors.New("mock error")
}

func main() {
	//u := User{}
	// 报nil指针问题
	//fmt.Println(u.Notify())

	u := User2{}
	fmt.Println(u.Receiver)
	fmt.Println(u.Notify())
}
