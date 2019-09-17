package main

import (
	"fmt"
	"reflect"
)

type Drawer interface {
	Draw() error
}

type Circle struct {
	x string
}

func (c Circle) Draw() error {
	fmt.Printf("c addr: %p\n", &c)
	fmt.Println("draw circle...")
	return nil
}

type Circle2 struct {
	x int
}

func (c2 *Circle2) Draw() error {
	fmt.Printf("c2 addr: %p\n", c2)
	fmt.Println("draw circle...")
	return nil
}

func DrawCircle(d Drawer) {
	fmt.Printf("c addr: %p\n", d)
	fmt.Printf("c addr: %p\n", &d)
	c, ok := d.(Circle)
	fmt.Println("d type", reflect.ValueOf(d).Kind())
	fmt.Println("DrawCircle", c, ok)
}

func DrawCircle2(d Drawer) {
	fmt.Printf("c2 addr: %p\n", d)
	fmt.Printf("c2 addr: %p\n", &d)
	c, ok := d.(*Circle2)
	fmt.Println("d type", reflect.ValueOf(d).Kind())
	fmt.Println("DrawCircle2", c, ok)
}

func DrawSomething(d interface{}) {
	fmt.Printf("c2 addr: %p\n", d)
	fmt.Printf("c2 addr: %p\n", &d)
	c, ok := d.(*Circle2)
	fmt.Println("d type", reflect.ValueOf(d).Kind())
	fmt.Println("DrawCircle2", c, ok)
}

func main() {
	c := Circle{}
	//c2 := Circle2{}
	DrawCircle(c)
	DrawCircle(c)
	DrawSomething(c)
	//DrawCircle2(c)
	var i Drawer = c
	fmt.Println("d type", reflect.ValueOf(i).Kind())
}
