/* Go quine */
package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	fmt.Println(bytes.Equal(nil, nil))

	testClosure2()
	fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

func testClosure1() {
	var err error = fmt.Errorf("test1 %s", "outer")
	go func() {
		ex, err := errInner()
		fmt.Println(ex, err)
	}()
	time.Sleep(2 * time.Millisecond)
	fmt.Println(err)
	ex, err := errInner()
	fmt.Println(ex, err)
}

func testClosure2() {
	var err error = fmt.Errorf("test1 %s", "outer")
	{
		ex, err := errInner()
		fmt.Println(ex, err)
	}
	fmt.Println(err)
	ex, err := errInner()
	fmt.Println(ex, err)
}

func errInner() (string, error) {
	return "inner", fmt.Errorf("test1 %s", "inner")
}

var q = `/* Go quine */
package main

import "fmt"

func main() {
    fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

var q = `
