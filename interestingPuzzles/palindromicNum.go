package main

import "fmt"

func main() {
	num := 11
	fmt.Printf("find num: %d", findLatestPalindNum(num))
}

func findLatestPalindNum(start int) int {
	for {
		if isPalindromicNumber(start, 10) && isPalindromicNumber(start, 2) && isPalindromicNumber(start, 8) {
			return start
		}
		start+=2
	}
}


func isPalindromicNumber(num, scale int) bool {
	tmp := make([]int, 0, 10000)
	for num > 0 {
		tmp = append(tmp , num % scale)
		num = num / scale
	}

	for i, j := 0, len(tmp) - 1; i < j; {
		if tmp[i] != tmp[j] {
			return false
		}
		i++; j--
	}

	return true
}
