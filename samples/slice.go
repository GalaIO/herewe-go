package main

import "fmt"

func main() {
	tmp := make([]int, 2)
	for i:=2; i<10; i++{
		tmp = append(tmp, i)
		fmt.Println(tmp)
	}
}
