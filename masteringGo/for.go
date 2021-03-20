package main

import "fmt"

func main() {
	// go语言的循环非常简单，只有一个关键字for
	// 但是支持4种语法

	arr := [...]int{1, 2, 3, 4}
	sum := 0

	// 第一种语法，类似c语言的for使用，for 初始化语句; 循环条件; 后置语句 {}
	// 是否继续循环的条件放在中间，由于go不支持逗号表达式，所以没办法想c一样 for(i=0, j=0; i<100; i++, j++) {}
	// 类似的用法可以是，for i, j:=0, 0; i<100; {i++;j++} 初始化和操作多个变量
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("sum =", sum)

	sum = 0
	i := 0
	// 第二种用法是 for 循环条件 {}
	// 这种非常类似 while语句，可以作为替换
	for i < len(arr) {
		sum += arr[i]
		i++
	}
	fmt.Println("sum =", sum)

	sum = 0
	i = 0
	// 第三种其实是死循环写法，只能靠break语句跳出，常见很多looper常见，比如等待连接。。。
	for {
		if i >= len(arr) {
			break
		}
		sum += arr[i]
		i++
	}
	fmt.Println("sum =", sum)

	sum = 0
	// 第四种是 for range语法，其实是一个语法糖，适合字符串、切片、数组、映射场景，可以同时获取索引+值，或者k+v
	// 在很多场景很方便，避免使用arr[i]的语法
	for i, v := range arr {
		fmt.Println("range index:", i)
		sum += v
	}
	fmt.Println("sum =", sum)

}
