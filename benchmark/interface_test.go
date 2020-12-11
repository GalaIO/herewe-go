package benchmark

import (
	"fmt"
	"testing"
)

type pString string

func (p *pString) String() string {
	return string(*(p))
}

type cString string
func (c cString) String() string {
	return string(c)
}

const (
	debug byte = iota
	prod
)

var logLevel = prod

func logIface(level byte, msg interface{}) {
	if level >= logLevel {
		fmt.Println(msg)
	}
}

func logString(level byte, msg string) {
	if level >= logLevel {
		fmt.Println(msg)
	}
}

func testStringer(level byte, s fmt.Stringer) {
	if level >= logLevel {
		fmt.Println(s)
	}
}

func testPString(level byte, s *pString) {
	if level >= logLevel {
		fmt.Println(s)
	}
}

func BenchmarkStringer(b *testing.B) {
	p := pString("test string")
	for i := 0; i < b.N; i++ {
		testPString(debug, &p)
	}
}

func BenchmarkStringer2(b *testing.B) {
	c := cString("test string")
	for i := 0; i < b.N; i++ {
		testStringer(debug, c)
	}
}

func BenchmarkFStringer(b *testing.B) {
	p := pString("test string")
	for i := 0; i < b.N; i++ {
		testStringer(debug, &p)
	}
}

func BenchmarkIface(b *testing.B) {
	p := pString("test string")
	for i := 0; i < b.N; i++ {
		logIface(debug, &p)
	}
}
func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logString(debug, "test string")
	}
}
