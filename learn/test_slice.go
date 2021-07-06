package main

import "fmt"

func main() {
	var tmp []int = nil

	for _, i := range tmp {
		fmt.Println(i)
	}

	tmp2 := append(tmp, 1)
	for _, i := range tmp2 {
		fmt.Println(i)
	}

	var tmap map[string]int = nil
	c, ok := tmap["hello"]
	fmt.Println(c, ok)

	for s, i := range tmap {
		fmt.Println(s, i)
	}

	var s []int
	for i := 0; i < 1000; i++ {
		if len(s) > 10 {
			s = s[10:]
		}
		s = append(s, i)
		fmt.Println(i, "s cap", cap(s))
	}
}
