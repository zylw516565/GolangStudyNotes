package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			go fmt.Print(0)
			fmt.Print(1)
		}
	}()

	time.Sleep(10 * time.Millisecond)
}
