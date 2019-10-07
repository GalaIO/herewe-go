package main

import (
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

type BaseHandler interface {
	Handler(*Person)
}

type printHandler struct {
}

func (p *printHandler) Handler(person *Person) {

	fmt.Println(person.Name)
}


func main() {

	person := &Person{
		Name: "xiaoguo",
	}
	handler := printHandler{}
	handler.Handler(person)
}
