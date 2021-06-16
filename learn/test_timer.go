package main

import (
	"fmt"
	"time"
)

func main() {
	timer_0 := time.NewTimer(-1)
	timer_1000 := time.NewTimer(1000 * time.Millisecond)

	fmt.Println("now", time.Now())
	for {
		select {
		case t := <-timer_0.C:
			fmt.Println("fired timer_0,", t, "now", time.Now())
		case t := <-timer_1000.C:
			fmt.Println("fired timer_1000,", t, "now", time.Now())
		}
		time.Sleep(1000)
	}
}
