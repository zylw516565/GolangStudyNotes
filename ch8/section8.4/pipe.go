package main

import (
	"fmt"
)

// func main() {
// 	naturals := make(chan int)
// 	squares  := make(chan int)

// 	ch  := make(chan struct{})

// 	go func () {
// 		for x :=0; x < 100; x++ {
// 			naturals <- x
// 		}

// 		close(naturals)
// 	}()

// 	go func () {
// 		for {
// 			y := <- naturals
// 			squares <- y*y
// 		}
// 	}()

// 	go func () {
// 		for {
// 			fmt.Println(<- squares)
// 		}
// 		ch <- struct{}{}
// 	}()

// 	<- ch
// }

func main() {
	naturals := make(chan int)
	squares  := make(chan int)

	ch  := make(chan struct{})

	go func () {
		for x :=0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func () {
		for {
			y, ok := <- naturals
			if !ok {
				break
			}
			squares <- y*y
		}
		close(squares)
	}()

	go func () {
		for {
			x, ok := <- squares
			if !ok {
				break
			}

			fmt.Println(x)
		}
		ch <- struct{}{}
	}()

	<- ch
}