package main

import "fmt"

func isValid(s string) bool {
	if s == "" {
		return true
	}

	sbytes := []byte(s)
	stack := make([]byte, len(sbytes))
	stackIndex := -1
	for _, ch := range sbytes {
		if ch == '(' || ch =='{' || ch=='[' {
			stackIndex++
			stack[stackIndex] = ch
			continue
		}
		if stackIndex<0 {
			return false
		}
		switch ch {
		case ')':
			if stack[stackIndex] != '('{
				return false
			}
		case '}':
			if stack[stackIndex] != '{'{
				return false
			}
		case ']':
			if stack[stackIndex] != '['{
				return false
			}

		}
		stackIndex --
	}
	if stackIndex >= 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}