package main

import (
	"fmt"
	"time"
)

type Timekeeper struct {
	myTime *time.Time
}

func main() {

	t := make([]Timekeeper, 100)
	var now time.Time
	for _, v := range t {
		now = time.Now()
		v.myTime = &now
		time.Sleep(1 * time.Second)
	}

	for i, times := range t {
		fmt.Printf("Timekeeper %s time: %s.", i, times.myTime.Format(time.RFC3339Nano))
	}

}
