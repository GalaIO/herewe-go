package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)
func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.StackInuse:", mem.StackInuse)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}
func main() {
	f, err := os.Create("traceFile.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()
	var mem runtime.MemStats
	printStats(mem)

	tmp := make([]byte, 1000000000)
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(time.Millisecond)
	}

	// 这时候tmp变量，指针的内存没有回收
	printStats(mem)
	fmt.Println(len(tmp), tmp[0])
	runtime.GC()
	// 这时候依然使用tmp变量，但是指针指向的内存已经被回收了
	printStats(mem)
	fmt.Println(len(tmp))
}