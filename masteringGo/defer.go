package main
import (
"fmt"
)

// 直接传参 i，保存在栈中，利用LIFO特性后进先出，返回1 2 3
func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

// 匿名函数，直接输出i，这时候由于闭包特性，i会被共享，返回0 0 0
func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

// 匿名函数，把i传参给函数，输出1 2 3
// 使用defer推荐这种方式 匿名函数+入参，防止非预期行为
func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}
func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
