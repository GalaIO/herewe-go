package main

import "fmt"

type Item struct {
	key string
	val interface{}
}

func main() {
	itemsA := []Item{
		Item{"1", 1},
		Item{"2", 2},
		Item{"3", 3},
	}
	tmp := Item{"4", 4}

	itemsA = append(itemsA, tmp)
	fmt.Printf("itemsA[3] addr:%p, tmp addr:%p\r\n", &itemsA[3], &tmp) // copy

	itemsB := []*Item{
		&Item{"1", 1},
		&Item{"2", 2},
		&Item{"3", 3},
	}
	tmpB := &Item{"4", 4}

	itemsB = append(itemsB, tmpB)
	fmt.Printf("itemsB[3] addr:%p, tmpB addr:%p\r\n", itemsB[3], tmpB) // copy
}
