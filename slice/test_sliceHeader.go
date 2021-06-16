package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "脑子进煎鱼了"
	// will re alloc for []byte
	b := []byte(s)
	fmt.Printf("s %v \n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	fmt.Printf("b %v \n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	news := bytes2string(b)
	newb := string2bytes1(news)
	fmt.Printf("s %v \n", (*reflect.StringHeader)(unsafe.Pointer(&news)).Data)
	fmt.Printf("b %v \n", (*reflect.SliceHeader)(unsafe.Pointer(&newb)).Data)
	fmt.Println(news, newb)
}

func string2bytes1(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	//make a typed var, and point to, will make data safe, not be gc
	var b []byte
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pbytes.Data = stringHeader.Data
	pbytes.Len = stringHeader.Len
	pbytes.Cap = stringHeader.Len

	//if string is back of static string, it will panic
	//unexpected fault address 0x4bcc93
	//fatal error: fault
	//[signal SIGSEGV: segmentation violation code=0x2 addr=0x4bcc93 pc=0x4977d3]

	// if string back of heap, it's ok
	//b[0] = 1

	return b
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	var s string
	tmp := (*reflect.StringHeader)(unsafe.Pointer(&s))
	tmp.Data = sliceHeader.Data
	tmp.Len = sliceHeader.Len
	return s
}
