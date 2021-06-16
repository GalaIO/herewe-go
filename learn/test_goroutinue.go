package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("go routinue in....")
		time.Sleep(time.Hour)
		fmt.Println("go routinue out....")
	}()

	time.Sleep(time.Second)
	fmt.Println("just shutdown...")
}
