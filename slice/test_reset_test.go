package main

import "testing"

type RestType struct {
	field0  int
	field1  int
	field2  int
	field3  int
	field4  int
	field5  int
	field6  int
	field7  int
	field8  int
	field9  int
	field10 int
	field11 int
	field12 int
	field13 int
	field14 int
	field15 int
	field16 int
	field17 int
	field18 int
	field19 int
	field20 int
	field21 int
	field22 int
	field23 int
	field24 int
	field25 int
	field26 int
	field27 int
	field28 int
	field29 int
	field30 int
	field31 int
	field32 int
	field33 int
	field34 int
	field35 int
	field36 int
	field37 int
	field38 int
	field39 int
	field40 int
	field41 int
	field42 int
	field43 int
	field44 int
	field45 int
	field46 int
	field47 int
	field48 int
	field49 int
	field50 int
	field51 int
	field52 int
	field53 int
	field54 int
	field55 int
	field56 int
	field57 int
	field58 int
	field59 int
	field60 int
	field61 int
	field62 int
	field63 int
	field64 int
	field65 int
	field66 int
	field67 int
	field68 int
	field69 int
	field70 int
	field71 int
	field72 int
	field73 int
	field74 int
	field75 int
	field76 int
	field77 int
	field78 int
	field79 int
	field80 int
	field81 int
	field82 int
	field83 int
	field84 int
	field85 int
	field86 int
	field87 int
	field88 int
	field89 int
	field90 int
	field91 int
	field92 int
	field93 int
	field94 int
	field95 int
	field96 int
	field97 int
	field98 int
	field99 int
}

func (r *RestType) Reset1() {
	r.field0 = 0
	r.field1 = 0
	r.field2 = 0
	r.field3 = 0
	r.field4 = 0
	r.field5 = 0
	r.field6 = 0
	r.field7 = 0
	r.field8 = 0
	r.field9 = 0
	r.field10 = 0
	r.field11 = 0
	r.field12 = 0
	r.field13 = 0
	r.field14 = 0
	r.field15 = 0
	r.field16 = 0
	r.field17 = 0
	r.field18 = 0
	r.field19 = 0
	r.field20 = 0
	r.field21 = 0
	r.field22 = 0
	r.field23 = 0
	r.field24 = 0
	r.field25 = 0
	r.field26 = 0
	r.field27 = 0
	r.field28 = 0
	r.field29 = 0
	r.field30 = 0
	r.field31 = 0
	r.field32 = 0
	r.field33 = 0
	r.field34 = 0
	r.field35 = 0
	r.field36 = 0
	r.field37 = 0
	r.field38 = 0
	r.field39 = 0
	r.field40 = 0
	r.field41 = 0
	r.field42 = 0
	r.field43 = 0
	r.field44 = 0
	r.field45 = 0
	r.field46 = 0
	r.field47 = 0
	r.field48 = 0
	r.field49 = 0
	r.field50 = 0
	r.field51 = 0
	r.field52 = 0
	r.field53 = 0
	r.field54 = 0
	r.field55 = 0
	r.field56 = 0
	r.field57 = 0
	r.field58 = 0
	r.field59 = 0
	r.field60 = 0
	r.field61 = 0
	r.field62 = 0
	r.field63 = 0
	r.field64 = 0
	r.field65 = 0
	r.field66 = 0
	r.field67 = 0
	r.field68 = 0
	r.field69 = 0
	r.field70 = 0
	r.field71 = 0
	r.field72 = 0
	r.field73 = 0
	r.field74 = 0
	r.field75 = 0
	r.field76 = 0
	r.field77 = 0
	r.field78 = 0
	r.field79 = 0
	r.field80 = 0
	r.field81 = 0
	r.field82 = 0
	r.field83 = 0
	r.field84 = 0
	r.field85 = 0
	r.field86 = 0
	r.field87 = 0
	r.field88 = 0
	r.field89 = 0
	r.field90 = 0
	r.field91 = 0
	r.field92 = 0
	r.field93 = 0
	r.field94 = 0
	r.field95 = 0
	r.field96 = 0
	r.field97 = 0
	r.field98 = 0
	r.field99 = 0
}

func (r *RestType) Reset2() {
	*r = RestType{}
}

var r = &RestType{}

func BenchmarkReset1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r.Reset1()
	}
}

func BenchmarkReset2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r.Reset2()
	}
}
