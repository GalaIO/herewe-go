package main

import (
	"fmt"
	"sync"
)

func main() {
	//test_rwlock()
	test_renlock()
}

func test_renlock() {
	lock := sync.RWMutex{}

	lock.Lock()
	fmt.Println("lock wlock1")
	defer func() {
		lock.Unlock()
		fmt.Println("release wlock1")
	}()

	// can not reentry lock?????
	lock.Lock()
	fmt.Println("lock wlock2")
	defer func() {
		lock.Unlock()
		fmt.Println("release wlock2")
	}()
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
