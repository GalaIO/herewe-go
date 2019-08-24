package main

import "fmt"

// 10. 正则表达式匹配
// https://leetcode-cn.com/problems/regular-expression-matching/

func isMatch(s string, p string) bool {
	return isMatchCurve(s, 0, p, 0)
}

func isMatchCurve(s string, i int, p string, j int) bool {
	if j >= len(p) {
		return i>=len(s)
	}
	firstMatch := i<len(s) && (s[i] == p[j] || p[j] == '.')

	res := false
	if j+1<len(p) && p[j+1] == '*' {
		res = res || isMatchCurve(s, i, p, j+2)
		if !res && firstMatch {
			res = res || isMatchCurve(s, i+1, p, j)
		}
	}else if !res && firstMatch {
		res = res || isMatchCurve(s, i+1, p, j+1)
	}
	return res
}

func main() {
	fmt.Println(isMatch("aba", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
