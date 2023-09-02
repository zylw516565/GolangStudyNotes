package main

// import (
// 	"os"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	abort := make(chan struct{})
// 	go func() {
// 		os.Stdin.Read(make([]byte, 1)) // read a single byte
// 		abort <- struct{}{}
// 	}()

// 	fmt.Println("Commencing countdown. Press return to abort.")
// 	tick := time.Tick(1 * time.Second)
// 	for countdown := 10; countdown > 0; countdown-- {
// 		fmt.Println(countdown)

// 		select {
// 			case	<-abort:
// 				fmt.Println("Launch aborted!")
// 				return
// 			case	<-tick:
// 		}
// 	}

// 	launch()
// }

// func launch() {
// 	fmt.Println("Launch... !!!")
// }

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	fmt.Println("Before ticker...")
	<- ticker.C
	fmt.Println("After ticker...")
	ticker.Stop()
}