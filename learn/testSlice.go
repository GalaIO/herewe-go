package main

import "fmt"

func main() {
	var modifySlice = func(s []int) {
		s[0] = 1
	}

	s := []int{0, 0}
	modifySlice(s)
	fmt.Println(s)
}