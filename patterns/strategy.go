package main

import "fmt"

type Comparer interface {
	Compare(interface{}) bool
}

type CmpInt int

func (c CmpInt) Compare(b interface{}) bool {
	foo := int(c)
	bar := int(b.(CmpInt))
	if foo < bar {
		return false
	}
	return true
}

type CmpFloat float64

func (c CmpFloat) Compare(b interface{}) bool {
	foo := float64(c)
	bar := float64(b.(CmpFloat))
	if foo < bar {
		return false
	}
	return true
}

func FindMaxInArr(arr []Comparer) Comparer {
	tmp := arr[0]
	for i := 1; i < len(arr); i++ {
		if !tmp.Compare(arr[i]) {
			tmp = arr[i]
		}
	}
	return tmp
}

func main() {
	fmt.Println(FindMaxInArr([]Comparer{CmpInt(1), CmpInt(4), CmpInt(3)}))
	fmt.Println(FindMaxInArr([]Comparer{CmpFloat(1), CmpFloat(4.1), CmpFloat(3)}))
}
