package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	ch <- "C"
	// ch <- "D"  //fatal error: all goroutines are asleep - deadlock!
	fmt.Println(len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Println(len(ch), cap(ch))
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func mirroedQuery() string {
	responses := make(chan string, 3)
	go func(){responses <- request("asia.gopl.io") }()
	go func(){responses <- request("europe.gopl.io") }()
	go func(){responses <- request("europe.gopl.io") }()
	return <- responses
}

func request(hostname string) (response string) { /* ... */ return "" }