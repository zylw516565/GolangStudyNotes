package main

import (
	"fmt"
)

// func main() {
// 	ch := make(chan int)
	
// 	go test(ch)
	
// 	x := 2
// 	ch <- x
// 	fmt.Println("exit")
// }



func main() {
	ch := make(chan int)
	y := <- ch
	fmt.Println(y)
	fmt.Println("exit")
}


func test(ch chan int) {
	y := <- ch
	fmt.Println(y)

	close(ch)
	y = <- ch
	fmt.Printf("%T, %#[1]v\n", y)

	// ch <- 3   //panic: send on closed channel
}
