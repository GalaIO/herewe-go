//JSON序列化和反序列化

//可用在api序列化输出
//转成结构体,方便程序操作等

package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	//布尔型
	boolByte, _ := json.Marshal(true)
	fmt.Println(string(boolByte))

	//整数型
	intByte, _ := json.Marshal(100)
	fmt.Println(string(intByte))

	//浮点型
	floatByte, _ := json.Marshal(1.23456)
	fmt.Println(string(floatByte))

	//字符串
	stringByte, _ := json.Marshal("字符串啊啊啊")
	fmt.Println(string(stringByte))

	//切片
	sliceByte, _ := json.Marshal([]string{"apple", "orange", "banana"})
	fmt.Println(string(sliceByte))

	//字典
	mapByte, _ := json.Marshal(map[string]int{"apple": 5, "orange": 6, "banana": 7})
	fmt.Println(string(mapByte))

	//自定义类型1
	customsByte1, _ := json.Marshal(&Response1{Page: 1, Fruits: []string{"apple", "orange", "banana"}})
	fmt.Println(string(customsByte1))

	//自定义类型2,tag语法
	customsByte2, _ := json.Marshal(&Response2{Page: 2, Fruits: []string{"apple", "orange", "banana"}})
	fmt.Println(string(customsByte2))

	//反序列化到结构体
	json1 := `{"Page":1,"Fruits":["apple","orange","banana"]}`
	json2 := `{"page":2,"fruits":["apple","orange","banana"]}`
	response1 := Response1{}
	response2 := Response2{}
	json.Unmarshal([]byte(json1), &response1)
	fmt.Println(response1)
	json.Unmarshal([]byte(json2), &response2)
	fmt.Println(response2)
}