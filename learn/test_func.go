package main

import (
	"encoding/hex"
	"fmt"
)

type TestStruct struct {
	name string
}

func (t *TestStruct) Clean1() {
}

func (t *TestStruct) Clean2() {
	t.name = ""
}

type Hash struct {
	data []byte
}

func (h *Hash) ToHex() string {
	if h == nil {
		return ""
	}

	return hex.EncodeToString(h.data)
}

func (m *Hash) Equal(h *Hash) bool {
	if m == h {
		return true
	}
	if h == nil {
		return false
	}

	if m.ToHex() == h.ToHex() {
		return true
	}

	return false
}

func main() {
	var t *TestStruct = nil
	// nil pointer can call func
	t.Clean1()
	// nil can not index filed
	//t.Clean2()

	var h *Hash
	compared := []byte("hello")
	fmt.Println(h.Equal(&Hash{data: compared}))

}
