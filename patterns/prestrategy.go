package main

import "fmt"

func compareInt(a int, b int) bool {
	if a < b {
		return false
	}
	return true
}

func FindMaxInIntArr(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}
	tmp := arr[0]
	for i := 1; i < len(arr); i++ {
		if !compareInt(tmp, arr[i]) {
			tmp = arr[i]
		}
	}
	return tmp
}

func compareFloat(a float64, b float64) bool {
	if a < b {
		return false
	}
	return true
}

func FindMaxInFloatArr(arr []float64) float64 {
	if len(arr) <= 0 {
		return 0
	}
	tmp := arr[0]
	for i := 1; i < len(arr); i++ {
		if !compareFloat(tmp, arr[i]) {
			tmp = arr[i]
		}
	}
	return tmp
}

func main() {
	fmt.Println(FindMaxInIntArr([]int{1, 4, 3, 1}))
	fmt.Println(FindMaxInFloatArr([]float64{1, 4.2, 3, 1}))
}
