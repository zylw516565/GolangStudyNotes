package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := 123
	y := fmt.Sprintf("%d", s)
	fmt.Println(y)
	fmt.Println(strconv.Itoa(s))
	fmt.Println(strconv.FormatInt(int64(s), 2))

	m := fmt.Sprintf("x=%b", s)
	fmt.Println(m)

	a, _ := strconv.Atoi("123")
	b, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(a, b)
	fmt.Printf("a type = %T, b type = %T\n", a, b)
}
