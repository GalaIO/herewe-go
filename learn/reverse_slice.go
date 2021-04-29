package main

import "fmt"

func main() {
	res := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(res)
	for i, j := 0, len(res)-1; i < j; i++ {
		res[i], res[j] = res[j], res[i]
		j--
	}
	fmt.Println(res)
}
