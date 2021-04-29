package main

import (
	"errors"
	"fmt"
)

type UError struct {
	msg string
}

func (u UError) Error() string {
	return "UError: " + u.msg
}

func testUError(msg string) (string, *UError) {
	return msg, nil
}

func testError(msg string) (string, error) {
	return msg, nil
}

func demo1() {
	fmt.Println("-------------demo1---------------")
	err := errors.New("test err1")
	fmt.Println(err != nil, err)

	msg, err := testUError("test err2")
	fmt.Println(err != nil, err, msg)
}

func demo2() {
	fmt.Println("-------------demo2---------------")
	_, err := testError("test err1")
	fmt.Println(err != nil, err)

	msg, err := testUError("test err2")
	fmt.Println(err != nil, err, msg)
}

func demo3() {
	fmt.Println("-------------demo3---------------")
	msg, err := testUError("test err2")
	fmt.Println(err != nil, err, msg)
}

/*
output
-------------demo1---------------
true test err1
true <nil> test err2
-------------demo2---------------
false <nil>
true <nil> test err2
-------------demo3---------------
false <nil> test err2
*/

func main() {
	demo1()
	demo2()
	demo3()
}
