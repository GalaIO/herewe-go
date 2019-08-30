package main

import "fmt"

func main() {
	//var tmp string = "111"
	var tmp2 interface{} = nil
	s, _ := tmp2.(string)
	fmt.Println(s)
	fun1()
}
func fun1() {
	i := 0
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d\n", i, j)
}
