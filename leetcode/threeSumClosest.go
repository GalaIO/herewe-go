package main

import (
	"sort"
	"fmt"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	minDiff := int(^uint(0)>>1)
	res := 0
	for i, num := range nums {
		wTarget := target - num

		for r, l := i+1, len(nums)-1; r < l; {
			sum := nums[r]+nums[l]
			if sum < wTarget {
				r++
			} else if sum > wTarget {
				l--
			} else {
				return target
			}

			diff := abs4Int(wTarget - sum)
			if diff < minDiff {
				minDiff = diff
				res = sum + num
			}
		}
	}
	return res
}

func abs4Int(n int) int {
	if n < 0{
		return -n
	}
	return n
}

func main() {
	case1 := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(case1, 1))
	case2 := []int{0, 0, 0}
	fmt.Println(threeSumClosest(case2, 1))
}
