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

type CoffeeDecorator struct {
	base CoffeeDrink
	addMilk float64
}

func (c *CoffeeDecorator) CoffeeContent() float64 {
	return c.base.CoffeeContent()
}
func (c *CoffeeDecorator) MilkContent() float64 {
	return c.base.MilkContent()+c.addMilk
}

func getContent(c CoffeeDrink) {
	fmt.Printf("coffee content: coffee=%fml, milk=%fml\r\n", c.CoffeeContent(), c.MilkContent())
}

func getContentDecorator(f func (CoffeeDrink), addMilk float64) func (CoffeeDrink) {
	// 声明新的接口
	return func (c CoffeeDrink) {
		// 增加实现
		f(c)
		fmt.Printf("aded milk=%fml\r\n", addMilk)
	}
}

func main() {
	coffee := &Coffee{}
	getContent(coffee)
	mkilkCoffee := &CoffeeDecorator{
		base:    coffee,
		addMilk: 100,
	}
	getContent(mkilkCoffee)
	mCoffee := &CoffeeDecorator{
		base:    coffee,
		addMilk: 80,
	}
	getContent(mCoffee)
	getContentDecorator(getContent, 100)(mCoffee)
}
