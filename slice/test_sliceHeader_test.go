package main

import "testing"

var s = "hello~~~~~~~hello~~~~~~~hello~~~~~~~hello~~~~~~~"
var bs = []byte("hello~~~~~~~hello~~~~~~~hello~~~~~~~hello~~~~~~~")

func BenchmarkString2bytes1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string2bytes1(s)
	}
}

func BenchmarkString2bytes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}
func BenchmarkByte2string1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bytes2string(bs)
	}
}

func BenchmarkByte2string2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(bs)
	}
}
