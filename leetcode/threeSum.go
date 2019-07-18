package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	res := make([][]int, 0, 10)

	for i:=0; i<len(nums); i++ {
		if i-1>0 && nums[i]==nums[i-1] {
			continue
		}
		target := 0-nums[i]
		for l,r := i+1, len(nums)-1; l<r; {
			if nums[l] + nums[r] > target {
				r--
			} else if nums[l] + nums[r] < target {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l+1<len(nums) && nums[l] == nums[l+1] {
					l++
				}
				for r-1>0 && nums[r] == nums[r-1] {
					r--
				}
				l++; r--
			}
		}
	}
	return res
}

func main() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}
