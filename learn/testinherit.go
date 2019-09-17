package main

import (
	"log"
	"strconv"
	"errors"
)

func main() {
	c := C{}
	c.tmp = 100
	log.Println(c.Say())
	log.Println(c.B.A.Say())
	log.Println(c.B.Say("hello"))
	log.Println(c.A.Say())
	log.Println(c.Say())
}

type A struct {
	tmp int
}

func (a A) innerSay() string {
	return "innerSay a " + strconv.Itoa(a.tmp)
}

func (a A) Say() string {
	return a.innerSay()
}

type B struct {
	A
}

func (b B) innerSay() string {
	return "innerSay b "
}

func (b B) Say(s string) string {
	return b.innerSay() + strconv.Itoa(b.tmp)
}

type C struct {
	B
}

func (c C) Say() string {
	return "c"
}

func testA(a A) error {
	return errors.New("test err")
}
