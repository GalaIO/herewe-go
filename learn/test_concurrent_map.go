package main

import (
	"fmt"
	"time"
)

func main() {
	var tmp = map[int]bool{}

	for i := 0; i < 100; i++ {
		tmp[i] = true
	}

	go func() {
		for {
			for i := 0; i < 100; i++ {
				fmt.Printf("g1, tmp[%v]: %v\r\n", i, tmp[i])
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			for i := 0; i < 100; i++ {
				fmt.Printf("g2, tmp[%v]: %v\r\n", i, tmp[i])
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			for i := 0; i < 100; i++ {
				//fatal error: concurrent map read and map write
				//tmp[i] = false
				fmt.Printf("g2, tmp[%v]: %v\r\n", i, tmp[i])
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {}
}
