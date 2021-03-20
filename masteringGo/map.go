package main

import "fmt"

func main() {
	// map就是哈希映射，可以用任何类型做键值对，不过对于键来说只有可比较类型才行，必须满足 ==
	// 如果hash散列均匀，理想情况下 查询时间O(1) 存储时间 O(1)
	// 显而易见使用bool做键非常不灵活，因为只会存在两个hash桶查询效率很低
	// 另外，由于不同机器和操作系统的浮点数精度定义不同，使用浮点数作为键值可能会出现异常。

	// 这里测试下 map键比较是只比较当前值，还是会比较内存
	// 切片和map不能相互比较，字符串可以，基本类型可以、自定义结构体(不包含切片、映射可以)，注意切片指针、映射指针是可以比较的
	type EqualTest struct {
		s *[]int
		//s []int // 无法比较，同时无法作为键
	}
	s1 := []int{1, 2}
	e1 := EqualTest{s: &s1}
	e2 := EqualTest{s: &s1}
	fmt.Println(e1==e2)

	// map的零值也是nil 这时候无法存储，但是可以查询，返回都是零值
	var m1 map[string]int
	fmt.Println(m1, m1==nil)
	//m1["1"] = 1 // panic: assignment to entry in nil map
	fmt.Println(m1["1"])

	// map初始化
	// 字面量初始化
	m1 = map[string]int {
		"1": 1,
		"2": 2,
	}
	// 手动申请内存初始化
	m2 := make(map[string]int, 10)
	m2["1"] = 1
	m2["2"] = 2
	fmt.Println(m1, m2)

	// 增加或修改映射，直接用[]方式设置即可，有两种情况，一种是键对应存在旧元素，一种是键对应没有旧元素，存在会覆盖，没有会新增
	m1["3"] = 3
	fmt.Println(m1)

	// 删除映射关系
	delete(m1, "3")
	fmt.Println(m1)

	// 查询也很简单，也有两种情况，一种是键对应存在元素，直接返回对应值，一种是键对应没有旧元素，会返回零值
	fmt.Println(m1["2"], m1["3"])
	// 这种可以主动获取是否存在，因为值可能也会有零值
	if r, ok := m1["3"]; !ok {
		fmt.Println("m1[\"3\"] not exist!, get", r)
	}

	// 遍历所有元素 只能使用range，获取kv
	for k, v := range m1 {
		fmt.Println("k:", k, "v:", v)
	}

	// 遍历所有k
	for k := range m1 {
		fmt.Println("k:", k)
	}
}
