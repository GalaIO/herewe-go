package main
import (
"fmt"
)

// 如愿从b的panic恢复，但是a为完全执行完，但是不会导致程序直接退出，main函数顺利执行
func a1() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!, error:", c)
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

// 因为defer是LIFO执行，所以，如果recover不是第一个defer位置的话，可能存在再次panic问题，
// 这时候虽然recover捕获到了b的panic但是无法捕获自己的
func a3() {
	fmt.Println("Inside a()")
	defer func() {
		panic("panic in a()!")
	}()
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!, error:", c)
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

// recover写在第一个defer，这时候b会触发panic，a也会通过defer触发panic，结果是recover只捕获到a的panic，b的被覆盖了
// 如果在panic之前再加一个recover，可以处理b的，所以如果defer存在panic可以在最开始和最后都加recover来保证panic不会被忽略
// 但是最好拆分模块和函数，分开处理
func a2() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			_, ok := c.(string)
			fmt.Println("Recover inside a()!, error:", c, ok)
		}
	}()
	defer func() {
		panic("panic in a()!")
	}()

	// 这个recover可以捕获到b的panic
	//defer func() {
	//	if c := recover(); c != nil {
	//		_, ok := c.(string)
	//		fmt.Println("Recover inside a()!, error:", c, ok)
	//	}
	//}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

func main() {
	a1()
	fmt.Print("\n\n------------\n\n")
	a2()
	fmt.Print("\n\n------------\n\n")
	a3()
	fmt.Print("\n\n------------\n\n")
	fmt.Println("main() ended!")
}
