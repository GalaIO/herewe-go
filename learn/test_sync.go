package main

import (
	"fmt"
	"sync"
)

func main() {
	test_rwlock()
}

func test_rwlock() {
	lock := sync.RWMutex{}

	lock.Lock()
	fmt.Println("lock wlock")
	defer func() {
		lock.Unlock()
		fmt.Println("release wlock")
	}()

	lock.RLock()
	fmt.Println("lock rlock")
	defer func() {
		lock.RUnlock()
		fmt.Println("release rlock")
	}()
}
