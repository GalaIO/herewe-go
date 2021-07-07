package main

import (
	"fmt"
	"html"
	"net/url"
)

func main() {
	str := "09.面试题：用队列\n实现栈&amp;用栈实现队列.mp4"
	str = html.UnescapeString(str)
	str = url.PathEscape(str)
	fmt.Println("https://d.shikey.com/jike/已完结的课程/33 算法面试通关40讲/" + str)
}
