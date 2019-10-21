package main

import "fmt"

type CoffeeDrink interface {
	CoffeeContent() float64
	MilkContent() float64
}

type Coffee struct {
}

func (c *Coffee) CoffeeContent() float64 {
	return 100
}
func (c *Coffee) MilkContent() float64 {
	return 0
}

type MilkCoffee struct {
}

func (c *MilkCoffee) CoffeeContent() float64 {
	return 100
}
func (c *MilkCoffee) MilkContent() float64 {
	return 100
}

type MCoffee struct {
}

func (c *MCoffee) CoffeeContent() float64 {
	return 100
}
func (c *MCoffee) MilkContent() float64 {
	return 80
}

func getContent(c CoffeeDrink) {
	fmt.Printf("coffee content: coffee=%fml, milk=%fml\r\n", c.CoffeeContent(), c.MilkContent())
}

func main() {
	getContent(&Coffee{})
	getContent(&MilkCoffee{})
	getContent(&MCoffee{})
}
