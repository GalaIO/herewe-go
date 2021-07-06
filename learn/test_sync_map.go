package main

import "sync"

func main() {
	m := sync.Map{}
	for i := 0; i < 10; i++ {
		m.Store(i, i)
	}
	for i := 0; i < 10; i++ {
		m.LoadAndDelete(i)
	}
}
