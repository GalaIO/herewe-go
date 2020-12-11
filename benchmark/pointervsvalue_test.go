package benchmark

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"
)

type pvStruct struct {
	val int
	a, b, c int
	s1, s2, s3 string
}

func newPtrPvStruct() *pvStruct {
	return &pvStruct{
		a:  100000,
		b:  100001,
		c:  100002,
		s1: "foo",
		s2: "bar",
		s3: "car",
	}
}

func newValPvStruct() pvStruct {
	return pvStruct{
		a:  100000,
		b:  100001,
		c:  100002,
		s1: "foo",
		s2: "bar",
		s3: "car",
	}
}



func (p pvStruct) valIncr() int {
	return p.val + 1
}

func (p pvStruct) valAdd(c pvStruct) int {
	return p.val + c.val
}

func (p pvStruct) valReturn() pvStruct {
	return p
}

func (p pvStruct) ptrReturn() *pvStruct {
	return &p
}

func (p *pvStruct) ptrIncr() int {
	return p.val + 1
}

func (p *pvStruct) ptrAdd(c *pvStruct) int {
	return p.val + c.val
}

func Benchmark_valIncr(t *testing.B) {
	var p pvStruct
	for i := 0; i < t.N; i++ {
		p.valIncr()
	}
}

func Benchmark_ptrIncr(t *testing.B) {
	var p pvStruct
	for i := 0; i < t.N; i++ {
		p.ptrIncr()
	}
}

func Benchmark_valAdd(t *testing.B) {
	var p pvStruct
	for i := 0; i < t.N; i++ {
		p2 := pvStruct{val: i}
		p.valAdd(p2)
	}
}

func Benchmark_ptrAdd(t *testing.B) {
	var p pvStruct
	for i := 0; i < t.N; i++ {
		p2 := pvStruct{val: i}
		p.ptrAdd(&p2)
	}
}

func Benchmark_valReturn(t *testing.B) {
	//var p pvStruct
	var l pvStruct
	for i := 0; i < t.N; i++ {
		l = newValPvStruct()
		l.valIncr()
		//p.valAdd(l)
	}
}

func Benchmark_ptrretun(t *testing.B) {
	//var p pvStruct
	var l *pvStruct
	for i := 0; i < t.N; i++ {
		l = newPtrPvStruct()
		l.ptrIncr()
		//p.ptrAdd(l)
	}
}

type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func BenchmarkMemoryStack(b *testing.B) {
	var s S

	f, err := os.Create("stack.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		s = byCopy()
	}

	trace.Stop()

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}

func BenchmarkMemoryHeap(b *testing.B) {
	var s *S

	f, err := os.Create("heap.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		s = byPointer()
	}

	trace.Stop()

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}
