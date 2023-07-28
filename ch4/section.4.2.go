package main

import (
	"fmt"
)

var months = [...]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December"}

func main() {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("slice:%v, len:%v, cap:%v\n", x[:], len(x[:]), cap(x[:]))
	fmt.Printf("slice:%v, len:%v, cap:%v\n", x[2:5], len(x[2:5]), cap(x[2:5]))
	fmt.Printf("slice:%v, len:%v, cap:%v\n", x[2:5:7], len(x[2:5:7]), cap(x[2:5:7]))
	fmt.Printf("slice:%v, len:%v, cap:%v\n", x[:4:6], len(x[:4:6]), cap(x[:4:6]))

	fmt.Println(months)

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2) // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
			fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// fmt.Println(summer[:20]) // panic: out of range
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer) // "[June July August September October]"

	name := "what is your name !!!"
	fmt.Printf("%T,%[1]v\n", name)
	n1 := name[:]
	fmt.Printf("%T,%[1]v\n", n1)
	n2 := name[5:7]
	fmt.Printf("%T,%[1]v\n", n2)

	bname := []byte("what is your name !!!")
	fmt.Printf("%T,%[1]v\n", bname)
	n3 := bname[:]
	fmt.Printf("%T, %[1]v, %[1]s\n", n3)
	n4 := bname[5:7]
	fmt.Printf("%T, %[1]v, %[1]s\n", n4)

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("slice:%v, len:%v, cap:%v\n", s, len(s), cap(s))
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	s2 := []int{60:3}
	fmt.Printf("slice:%v, len:%v, cap:%v\n", s2, len(s2), cap(s2))
	// s3 := []int{10:3}
	// fmt.Println(s2 == s3)  //invalid operation: (slice can only be compared to nil)

	s4 := []int{}
	var s5 []int
	fmt.Println(s4 == nil, s5 == nil)

	var s6 []int // len(s) == 0, s == nil
	fmt.Printf("slice:%v, %[1]T, len:%v, cap:%v\n", s6, len(s6), cap(s6))
	s6 = nil // len(s) == 0, s == nil
	fmt.Printf("slice:%v, %[1]T, len:%v, cap:%v\n", s6, len(s6), cap(s6))
	s6 = []int(nil) // len(s) == 0, s == nil
	fmt.Printf("slice:%v, %[1]T, len:%v, cap:%v\n", s6, len(s6), cap(s6))
	s6 = []int{} // len(s) == 0, s != nil
	fmt.Printf("slice:%v, %[1]T, len:%v, cap:%v\n", s6, len(s6), cap(s6))
	s6 = make([]int, 3)[3:]
	fmt.Printf("slice:%v, %[1]T, len:%v, cap:%v\n", s6, len(s6), cap(s6))

	s7 := make([]int, 5)
	fmt.Printf("slice:%v, %[1]T, len:%v, cap:%v\n", s7, len(s7), cap(s7))
	s8 := make([]int, 5, 7)
	fmt.Printf("slice:%v, %[1]T, len:%v, cap:%v\n", s8, len(s8), cap(s8))

	var runes []rune
	for _, r := range "Hello, 世界" {
	runes = append(runes, r)
	} 
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"

	var x1, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x1, i)
		// fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		fmt.Printf("%d cap=%d\n", i, cap(y))
		x1 = y
	}
	fmt.Println("******************************")

	var x2, y2 []int
	for i := 0; i < 10; i++ {
		y2 = append(x2, i)
		// fmt.Printf("%d cap=%d\t%v\n", i, cap(y2), y2)
		fmt.Printf("%d cap=%d\n", i, cap(y2))
		x2 = y2
	}

	var x3 []int
	x3 = append(x3, 1)
	x3 = append(x3, 2, 3)
	x3 = append(x3, 4, 5, 6)
	x3 = append(x3, x3...) // append the slice x
	fmt.Println(x3) // "[1 2 3 4 5 6 1 2 3 4 5 6]"

	str1 := []string{"nihao", "", "hello", "", "world"}
	fmt.Println(nonempty(str1))

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data) // `["one" "three" "three"]`

	data = nonempty(data)
	fmt.Printf("%q\n", data) // `["one" "three" "three"]`

	str2 := []string{"nihao", "", "hello", "", "world"}
	fmt.Println(nonemptyWithAppend(str2))

	str3 := []string{"nihao", "", "hello", "", "world"}
	fmt.Println(nonemptyMyEdition(str3))
	str4 := []string{"nihao", "", "hello", "", "world", ""}
	fmt.Println(nonemptyMyEdition(str4))

	//slice <--> stack
	var stack []int; var v1 int
	stack = append(stack, v1)    //push
	top := stack[len(stack) - 1] //top of stack
	_ = top
	stack = stack[:len(stack) - 1] //pop
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	s[i], s[j] = s[j], s[i]
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
	// There is room to grow. Extend the slice.
	z = x[:zlen]
	} else {
	// There is insufficient space. Allocate a new array.
	// Grow by doubling, for amortized linear complexity.
	zcap := zlen
	if zcap < 2*len(x) {
	zcap = 2 * len(x)
	} 
	z = make([]int, zlen, zcap)
	copy(z, x) // a built-in function; see text
	} 
	z[len(x)] = y
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _,str := range strings {
		if str != "" {
			strings[i] = str
			i++
		}
	}

	return strings[:i]
}

func nonemptyWithAppend(strings []string) []string {
	i := 0
	for _,str := range strings {
		if str != "" {
			strings = append(strings[:i], str)
			i++
		}
	}

	return strings[:i]
}

func nonemptyMyEdition(strings []string) []string {
	j := 0
	for ; j + 1 < len(strings); {
		if len(strings[j]) == 0 {
			copy(strings[j:], strings[j+1:])
			strings = strings[:len(strings)-1]
		} else {
			j++
		}
	}

	if strings[len(strings)-1] == "" {
		strings = strings[:len(strings)-1]
	}

	return strings
}