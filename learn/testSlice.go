package main

import "fmt"

func main() {
	var modifySlice = func(s []int) {
		s[0] = 1
	}

	s := []int{0, 0}
	modifySlice(s)
	fmt.Println(s)

	s2 := []int{0, 1, 2, 3}
	fmt.Println(len(s2), cap(s2))
	s3 := s2[:3]
	fmt.Println(len(s3), cap(s3))
	s3 = s3[:cap(s3)]
	fmt.Println(len(s3), cap(s3))
}