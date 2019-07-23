package main

import "fmt"

// 22 括号生成
// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	//lc, rc := n, n

	result := make([]string, 0)

	if n < 1 {
		return result
	}

	temp := make([]byte, n*2)
	idx := 0

	generatePattern(&result, temp, idx, n, n)
	return result
}
func generatePattern(result *[]string, temp []byte, idx int, lc int, rc int) {
	if lc <1 && rc < 1 {
		*result = append(*result, string(temp))
		return
	}
	if lc == rc {
		temp[idx] = '('
		generatePattern(result, temp, idx+1, lc-1, rc)
	} else if lc < 1 {
		temp[idx] = ')'
		generatePattern(result, temp, idx+1, lc, rc-1)
	} else {
		temp[idx] = '('
		generatePattern(result, temp, idx+1, lc-1, rc)
		temp[idx] = ')'
		generatePattern(result, temp, idx+1, lc, rc-1)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
