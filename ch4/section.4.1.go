package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	var a1 [3]int // array of 3 integers
	fmt.Println(a1[0]) // print the first element
	fmt.Println(a1[len(a1)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a1 {
	fmt.Printf("%d %d\n", i, v)
	}
 
	// Print the elements only.
	for _, v := range a1 {
	fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2], q) // "0"

	s := [...]int{1, 2, 3}
	fmt.Printf("%T\n", s) // "[3]int"

	// t := [3]int{1, 2, 3}
	// t = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int

	type Currency int

	const (
		USD Currency = iota // 美元
		EUR // 欧元
		GBP // 英镑
		RMB // 人民币
		)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB], symbol) // "3 ￥"

	u := [...]int{99: -1}
	fmt.Println(u)

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	// d := [3]int{1, 2}
	// fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	v := [32]byte{31:2}
	fmt.Println(v)
	zero(&v)
	fmt.Println(v)

	w := [32]byte{31:3}
	fmt.Println(w)
	zeroV2(&w)
	fmt.Println(w)

}

func zero(ptr *[32]byte) {
	for i := range ptr {
	ptr[i] = 0
	}
}

func zeroV2(ptr *[32]byte) {
	*ptr = [32]byte{}
}