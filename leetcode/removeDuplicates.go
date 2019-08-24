package main

import "fmt"

func removeDuplicates(nums []int) int {
	var i, count = 0, 1
	for i<len(nums)-count {
		if nums[i] != nums[i+1] {
			i++; continue
		}
		for j:=i; j<len(nums)-1; j++ {
			nums[j] = nums[j+1]
		}
		count ++
	}
	return i+1
}

func main() {
	case1 := []int{1, 1, 2}
	fmt.Println(removeDuplicates(case1), case1)
	fmt.Println(removeDuplicates([]int{1,1,1,1,1,2}))
	fmt.Println(removeDuplicates([]int{1,1,2,2,3,4}))
}
