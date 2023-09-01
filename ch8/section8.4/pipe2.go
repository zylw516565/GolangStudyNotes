package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares  := make(chan int)

	go func () {
		for x :=0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func () {
		for n := range naturals {
			squares <- n*n
		}
		close(squares)
	}()

	for s := range squares {
		fmt.Println(s)
	}
	// close(squares)  //panic: close of closed channel
}