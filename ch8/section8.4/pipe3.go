package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares  := make(chan int)

	ch  := make(chan<- int)
	go counter(ch)
	ch = naturals
	// naturals = ch  //cannot use ch (variable of type chan<- int) as chan int value in assignment

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x :=0; x < 100; x++ {
		out <- x
	}
	// y := <-out  //cannot receive from send-only channel out
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for n := range in {
		out <- n*n
	}
	close(out)
}

func printer(in <-chan int) {
	for s := range in {
		fmt.Println(s)
	}
}