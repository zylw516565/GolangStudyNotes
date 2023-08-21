package main

import (
	"os"
	"fmt"
	"runtime"
)

func main () {
	defer printStack()
	f(3)
	// test1()
}

func test1() {
	fmt.Println("Will be panic")
	defer fmt.Println("defer 1")
	panic(1)
	defer fmt.Println("defer 2")
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)  // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x -1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}