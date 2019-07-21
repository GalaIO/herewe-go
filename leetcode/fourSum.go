package main

import (
	"sort"
	"fmt"
)

func fourSum(nums []int, target int) [][]int {

	result := make([][]int, 0)

	if len(nums) < 1 {
		return result
	}

	sort.Ints(nums)
	for i, num := range nums {
		if i>0 && nums[i-1] == nums[i] {
			continue
		}
		threeSum := target - num
		for j:=i+1; j<len(nums); j++ {
			if j!=i+1 && j>0 && nums[j-1] == nums[j] {
				continue
			}
			twoSum := threeSum - nums[j]
			for l,r := j+1,len(nums)-1; l<r; {
				if l!=j+1 && l>0 && nums[l-1] == nums[l] {
					l++
					continue
				}
				cSum := nums[l] + nums[r]
				if cSum > twoSum {
					r--
				}else if cSum < twoSum {
					l++
				}else {
					// 找到
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				}
			}
		}
	}
	return result
}

func main() {
	case1 := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(case1, 0))
}
