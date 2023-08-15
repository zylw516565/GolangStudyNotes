package main

import (
	"os"
	"fmt"
)

func f(...int) {}
func g([]int) {}

func main () {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1,2,3,4))
	fmt.Println(sum([]int{1,2,3,4}...))


	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

	linenum, name := 5, "mmmm"
	errorf(linenum, "undef: %s", name)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args)
	fmt.Fprintln(os.Stderr)
}