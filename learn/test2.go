package main

import "fmt"

func main() {
	//var tmp string = "111"
	var tmp2 interface{} = nil
	s, _ := tmp2.(string)
	fmt.Println(s)
}
