package main

import "fmt"

type FrogPrince struct {
	isFrog bool
}

func NewFrogPrince() *FrogPrince {
	return &FrogPrince{
		isFrog: true,
	}
}

func (f *FrogPrince) Kiss() {
	f.isFrog = false
}

func (f *FrogPrince) Greet() {
	if f.isFrog {
		fmt.Println("guagua~~")
	} else {
		fmt.Println("hello~~")
	}
}

func main() {
	frogPrince := NewFrogPrince()
	frogPrince.Greet()
	frogPrince.Kiss()
	frogPrince.Greet()
}
