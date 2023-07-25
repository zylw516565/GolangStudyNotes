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
}
