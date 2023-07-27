package main

import (
	"fmt"
)

func main() {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("slice:%v, len:%v, cap:%v\n", x[:], len(x[:]), cap(x[:]))
	fmt.Printf("slice:%v, len:%v, cap:%v\n", x[2:5], len(x[2:5]), cap(x[2:5]))
	fmt.Printf("slice:%v, len:%v, cap:%v\n", x[2:5:7], len(x[2:5:7]), cap(x[2:5:7]))
	fmt.Printf("slice:%v, len:%v, cap:%v\n", x[:4:6], len(x[:4:6]), cap(x[:4:6]))
}
