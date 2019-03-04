package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1w 以内的素数表 用于非对称加密使用，素数之间互质
var primeList = []int{2, 3, 5, 7, 11, 13}

// 生产rsa的n e d
func GenerationNED() (int, int, int) {
	// 随机选两个质数p q n=p*q
	cRand := rand.New(rand.NewSource(time.Now().Unix()))
	p := cRand.Intn(len(primeList))
	q := cRand.Intn(len(primeList))
	fmt.Printf("rand p: %d, q: %d\r\n", p, q)
	for q == p {
		q = rand.Intn(len(primeList))
	}
	n := p * q

	fmt.Printf("find n: %d\r\n", n)
	// 求l l为p-1 q-1的最小公倍数
	l1 := gcd(p-1, q-1)
	l := (p - 1) * (q - 1) / l1
	fmt.Printf("find l: %d\r\n", l)

	// 求e，e与l互质，尽量求最大的e
	e := l - 1
	for gcd(l, e) != 1 {
		e = e - 1
	}
	fmt.Printf("find e: %d\r\n", e)

	// 求d，e * d mode l = 1
	var index int
	for index = 1; (index*l+1)%e != 0; index++ {

	}
	d := (index*l + 1) / e
	fmt.Printf("find d: %d\r\n", d)
	return n, e, d
}

func gcd(p int, q int) int {
	r := p % q
	for r != 0 {
		p = q
		q = r
		r = p % q
	}
	return q
}

func main() {
	n, e, d := GenerationNED()
	fmt.Println(n, e, d)
	// 加密 对字符串每一个字符加密
	length := 10
	sBytes := make([]int, length)
	for i:=0;i < length;i++ {
		sBytes[i] = i
	}
	fmt.Println(sBytes)

	for i, ch := range sBytes {
		sBytes[i] = pow(ch, e) % n
	}
	fmt.Println(sBytes)

	for i, ch := range sBytes {
		sBytes[i] = pow(ch, d) % n
	}
	fmt.Println(sBytes)
}
func pow(b int, i int) int {
	res := 1
	for ; i>0; i--{
		res *= b
	}
	return res
}
