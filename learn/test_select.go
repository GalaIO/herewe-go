package main

import (
	"fmt"
	"time"
)

type Tmp struct {
	ch chan int64
}

func NewTmp() Tmp {
	return Tmp{
		ch: make(chan int64, 10),
	}
}

func (c *Tmp) getCh() chan int64 {
	fmt.Println("select getCh....")
	return c.ch
}

func main() {

	ch1 := make(chan int64, 10)
	tmp := NewTmp()

	tmp.ch <- 1
	ch1 <- 1
	// main loop
	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println("case1")
			case <-tmp.getCh():
				fmt.Println("case2")
			}
		}
	}()

	fmt.Println("wait 5s...")
	time.Sleep(5 * time.Second)
	// send loop
	fmt.Println("replace tmp ch...")
	tmp = NewTmp()
	go func() {
		fmt.Println("trigger tmp.ch per 1s...")
		for {
			time.Sleep(1 * time.Second)
			tmp.ch <- 1
		}
	}()

	fmt.Println("wait 10s...")
	time.Sleep(10 * time.Second)
	fmt.Println("trigger ch1...")
	//if do not trigger ch1 again, the select will never print case2
	ch1 <- 1

	select {}
}
