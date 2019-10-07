package main

import (
	"strings"
	"fmt"
	"errors"
)

type Shape interface {
	Draw()
}

type Circle struct {
}

func (c *Circle) Draw()  {
	fmt.Println("it's circle~")
}

func ShapeFactory(name string) (Shape, error) {
	if strings.EqualFold(name, "circle") {
		return &Circle{}, nil
	}
	return nil, errors.New("cannot find target~")
}

func main() {
	shape, err := ShapeFactory("circle")
	if err != nil {
		panic(err)
	}
	shape.Draw()
}
