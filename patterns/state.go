package main

import "fmt"

type Greeter interface {
	Greet()
}

type Frog struct {
}

func (f *Frog) Greet() {
	fmt.Println("guagua~~")
}

type Prince struct {
}

func (p *Prince) Greet() {
	fmt.Println("hello~~")
}

type FrogPrinceState struct {
	cstate Greeter
}

func NewFrogPrinceState() *FrogPrinceState {
	return &FrogPrinceState{
		cstate: &Frog{},
	}
}

func (f *FrogPrinceState) Kiss() {
	f.cstate = &Prince{}
}

func (f *FrogPrinceState) Greet() {
	f.cstate.Greet()
}

func main() {
	frogPrinceState := NewFrogPrinceState()
	frogPrinceState.Greet()
	frogPrinceState.Kiss()
	frogPrinceState.Greet()
}

