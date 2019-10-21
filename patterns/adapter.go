package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type Circle struct {
	X int
	Y int
	R int
}

type CircleAdapter struct {
	base Circle
}

func (c *CircleAdapter) GetPoint() *Point {
	return &Point{
		X:c.base.X,
		Y:c.base.Y,
	}
}

func (c *CircleAdapter) GetR() int {
	return c.base.R
}

type CircleV2 interface {
	GetPoint() *Point
	GetR() int
}

func main() {
	c := CircleAdapter{
		base: Circle{
			X: 2,
			Y: 3,
			R: 10,
		},
	}
	fmt.Printf("circle's point: %s, radius: %d.\r\n", c.GetPoint(), c.GetR())
}


