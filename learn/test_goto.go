package main

import "fmt"

func main() {

	// break label, continue label 都是和for循环配合使用，分别表示跳出多层循环，和继续多层循环
	// 默认 break continue只能作用于本次循环，label可以作用于多层循环，语法必须写在控制的for循环上面，不能写在后面
	// break 直接跳出对应循环语句，继续执行
	// continue 继续执行 i++(循环末更新语句) 和 条件判断，不会执行初始化语句，继续执行该循环
loop100_1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("loop100_1", i, j)
			if j == 5 {
				break loop100_1
			}
		}
	}

loop100_2:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("loop100_2", i, j)
			if j == 5 {
				continue loop100_2
			}
		}
	}

	// goto 直接跳到某一个语句，非常自由
	// 在这里重复执行了for的初始化语句，相当于重新执行循环
	k := 0
loop100_3:
	for i := 0; i < 10; i++ {
		if k++; k == 20 {
			break loop100_3
		}
		for j := 0; j < 10; j++ {
			fmt.Println("loop100_3", i, j, k)
			if j == 5 {
				goto loop100_3
			}
		}
	}
}
