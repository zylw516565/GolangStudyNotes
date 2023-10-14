package main

func main() {
	ch := make(chan int)

	go func(ch chan int) int {
		return <-ch
	}(ch)
}
