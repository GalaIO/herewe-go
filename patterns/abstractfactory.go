package main

import "fmt"

type Shape interface {
	Draw()
}

type Colorer interface {
	Color() string
}

type AbstractFactory interface {
	GetShape() Shape
	GetColor() Colorer
}

type Red struct {
}

func (r *Red) Color() string {
	return "red"
}

type Cube struct {
}

func (r *Cube) Draw() {
	fmt.Println("it's cube~~")
}

type RedCube struct {
}

func (r *RedCube) GetShape() Shape {
	return &Cube{}
}

func (r *RedCube) GetColor() Colorer {
	return &Red{}
}

func DrawPic(tool AbstractFactory) {
	fmt.Println("color is", tool.GetColor().Color())
	tool.GetShape().Draw()
}

func main() {
	cube := &RedCube{}
	DrawPic(cube)
}
