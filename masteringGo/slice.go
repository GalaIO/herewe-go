package main

import (
	"fmt"
	"sort"
)

func main() {
	// 切片的定义 很简单，需要和数组区别的是，数组是长度确定的类型，切片是长度不确定的
	var s []int // 切片 未初始化时，默认值为nil go支持切片为nil计算len和cap
	var a [4]int // 数组，未初始化，默认值为[0 0 0 0]
	fmt.Println(s, len(s), s==nil, a)

	// 判断切片为空，有两个可能一个是 切片为零值nil，或者切片长度为0，可以使用如下判断兼容两个场景：
	if len(s) == 0 {
		fmt.Println("s is empty")
	}

	// 切片和数组的初始化，是一样的，都是通过{1, 2, 3}设置初始值，否则默认是该类型的零值
	s = []int{1, 2, 3, 4}
	s2 := make([]int, 5)
	a = [4]int{1, 2}
	a2 := [...]int{1, 3, 4} // 数组初始化可以不指定具体数量，让编译器推断，区别于切片 定义为[...]int{1, 2}
	fmt.Println(s, s2, a, a2)

	// 切片的底层实现就是数组，一块连续的内存区域，同时具有两个属性，一个当前切片长度，一个容量
	s3 := make([]int, 5, 10) // 创建一个长度5 但是容量10的切片，在append元素超过5时，可以继续使用降低扩容频次
	fmt.Println(s3, "len =", len(s3), "cap =", cap(s3))
	// 数组可以非常容易变为切片, 通过[:]即可，或者[start:end]取start和end中间不包含end元素，
	// 如果省略start或者end，默认start=0 end=len(a)
	// 切片也可以通过[start:end] 生成另一个切片不过他们底层内存是共享的
	//s3 := a[:]
	s4 := s3[3:]
	fmt.Println(s4, "len =", len(s4), "cap =", cap(s4))
	// 切片重建的时候 可以以容量大小重建
	fmt.Println(s3[:10])

	s4[0] = 1; s4[1] = 2
	fmt.Println(s3, s4) // [0 0 0 1 2] [1 2] 共享底层存储，只是复制了栈中的切片内容而已

	// go提供了copy方法，来拷贝值，这样新的切片修改和原来的毫无影响，注意要重新申请切片内存才行，不能使用[:]创造切片
	// copy在复制时会从索引0开始复制，复制的长度为两个切片的最短切片
	s5 := make([]int, 5)
	copy(s5, s3)
	s5[0] = 1; s5[1] = 2
	fmt.Println(s3, s5) // [0 0 0 1 2] [1 2 0 1 2]

	// 多维切片，和多维数组类似，直接看初始化 这个是字面量初始化
	s6 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}
	a3 := [2][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
	}
	fmt.Println(s6, a3)

	// 使用make动态初始化，只能初始化一级，二维是零值nil 因为切片零值就是nil
	s7 := make([][]int, 3)
	fmt.Println(s7, s7[0] == nil)
	// 一次初始化二维切片
	for i:=0; i<len(s7); i++ {
		s7[i] = make([]int, 3)
	}
	fmt.Println(s7, s7[0] == nil, s7[0])

	// 切片扩容，切片其实就是动态数组，支持动态扩容，就是使用append，但是append返回时新的切片，所以一般覆盖旧切片值
	// 如果当前容量足够 是不会申请新的内存，因为触发扩容会默认申请2倍的内存
	// 所以切片最好是申请足够容量  不扩容，还有类似pool的用法 复用切片内存
	s8 := make([]int, 3, 10)
	_ = append(s8, 4, 5, 6)
	fmt.Println(s8, cap(s8))
	s8 = append(s8, 5)
	fmt.Println(s8, cap(s8), s8[:10]) // [0 0 0 5] 10 [0 0 0 5 5 6 0 0 0 0]

	// 切片排序，排序是刷题或者业务常见的操作，由于go不支持泛型，所以内置几种常见类型排序
	s9 := []int{9, 8, 6 ,2}
	sort.Ints(s9)
	fmt.Println(s9)

	s10 := []string{"z", "zs", "ba", "bz", "a"}
	sort.Strings(s10)
	fmt.Println(s10)

	// 提供一个通用方案，通过反射实现的，需要反射处理多种类型，操作内存。。。很恶心，要考虑情况太多了
	// 第一个入参时切片，第二个是传递一个比较元素方法，切片还会以闭包形式在匿名函数
	s9 = []int{9, 8, 6 ,2}
	sort.Slice(s9, func(i, j int) bool {
		if s9[i] < s9[j] {
			return true
		}
		return false
	})
	fmt.Println(s9)

}
