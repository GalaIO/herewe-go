package main

import "fmt"

func main() {
	fmt.Println(longestMountain([]int{1,2,0,2,0,2}))
}

func longestMountain(A []int) int {
	max := 0
	for l:= 0; l + 2 < len(A); {
		r := l + 1
		if A[l] >= A[l+1] {
			l = r
			continue
		}

		for r + 1 < len(A) && A[r] < A[r+1]  {
			r++
		}

		if r + 1 >= len(A) || A[r] <= A[r+1]{
			l = r
			continue
		}

		for r + 1 < len(A) && A[r] > A[r+1]  {
			r++
		}

		if r - l + 1 > max {
			//fmt.Println(l ,r)
			max = r - l + 1
		}
		l = r
	}

	if max >= 3 {
		return max
	}

	return 0
}
