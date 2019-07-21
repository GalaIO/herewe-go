package main

import "fmt"

var numsMap = [][]string{
	{},
	{},
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "g", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}
func letterCombinations(digits string) []string {
	comRes := make([]string, 0)

	if digits == "" {
		return comRes
	}
	dBytes := []byte(digits)


	for _, ch := range dBytes {
		comRes = letterCombinateWapper(comRes, numsMap[ch-'0'])
	}

	return comRes
}
func letterCombinateWapper(source []string, combinated []string) []string {

	if len(source) < 1 {
		return combinated
	}

	arrCount := len(source) * len(combinated)

	result := make([]string, 0, arrCount)
	for _, s1 := range source {
		for _, s2 := range combinated {
			result = append(result, s1+s2)
		}
	}

	return result
}

func main() {
	case1:= "222"
	fmt.Println(letterCombinations(case1))
}