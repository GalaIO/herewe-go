package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	//var tmp string = "111"
	var tmp2 interface{} = nil
	s, _ := tmp2.(string)
	fmt.Println(s)
	fun1()
	i := int64(10)
	fmt.Println(i)
	atomic.AddInt64(&i, 1)
	fmt.Println()
}
func fun1() {
	i := 0
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d\n", i, j)
}
