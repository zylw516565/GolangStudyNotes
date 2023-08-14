package main

import (
	"log"
	"fmt"
	"strings"
)

func square(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }

func add1(r rune) rune {return r + 1}

func main () {
	f := square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(3)) // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"
	// f = product  //cannot use product (value of type func(m int, n int) int) as func(n int) int value in assignment

	// var f1 func (n int) int
	// f1(1)  //panic: runtime error: invalid memory address or nil pointer dereference

	var f2 func (n int) int
	if f2 != nil {
		f2(3)
	}

	log.SetPrefix("wait: ")
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Printf("Site is down:")

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS")) // "WNT"
	fmt.Println(strings.Map(add1, "Admix")) // "Benjy"
}