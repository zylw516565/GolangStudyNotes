package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")

	select {
		case  <-time.After(10 * time.Second):
			//Do Nothing
		// case  <- chan int(nil):
		case	<-abort:
			fmt.Println("Launch aborted!")
			return
		// default:
		// 	fmt.Println("Waiting !!!")
	}

	launch()
}

func launch() {
	fmt.Println("Launch... !!!")
}

// import (
// 	"fmt"
// )

// func main() {
// 	ch := make(chan int, 1)
// 	// ch := make(chan int, 5)  //output random
// 	for i := 0; i < 10; i++ {
// 		select {
// 			case ch <- i:
// 				//Do Nothing!
// 			case x := <- ch:
// 				fmt.Println(x)
// 			default:
// 		}
// 	}
// }
