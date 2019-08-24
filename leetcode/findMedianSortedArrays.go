package main

import "fmt"

// 4. 寻找两个有序数组的中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

// 模拟归并 找到中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j, k = 0, 0, 0
	tcount := len(nums1) + len(nums2) + 1
	var t1, t2 = tcount/2-1, tcount/2+tcount%2-1
	var val1 = 0

	for ; i<len(nums1) && j<len(nums2); k++ {
		var cVal int
		if nums1[i] < nums2[j] {
			cVal = nums1[i]
			i++
		}else {
			cVal = nums2[j]
			j++
		}
		if k == t1 {
			val1 = cVal
		}

		if k == t2 {
			return float64(val1 + cVal) / 2
		}
	}
	for ; i<len(nums1); k++ {
		var cVal = nums1[i]
		if k == t1 {
			val1 = cVal
		}

		if k == t2 {
			return float64(val1 + cVal) / 2
		}
		i++
	}
	for ; j<len(nums2); k++ {
		var cVal = nums2[j]
		if k == t1 {
			val1 = cVal
		}

		if k == t2 {
			return float64(val1 + cVal) / 2
		}
		j++
	}
	return 0.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2,4}))
}
