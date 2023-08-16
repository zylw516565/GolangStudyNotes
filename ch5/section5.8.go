package main

import (
	"os"
	"log"
	"fmt"
	"sync"
	"time"
	"io/ioutil"
)

func main () {
	filename := "./test.data"

	data, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%q\n", data)

	// bigSlowOperation()
	testDeferInLoop()
	testDeferScope()
	// testPreCompute()

	_ = double(4)
	fmt.Println(triple(4))
}

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

var mu sync.Mutex
var m = make(map[string]int)
func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func bigSlowOperation() {
	// defer trace("bigSlowOperation") // don't forget the
	defer trace("bigSlowOperation")()

	time.Sleep(3 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)

	return func () {
		log.Printf("exit %s (%s)\n", msg, time.Since(start))
	}
}

func testDeferInLoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func testDeferScope() {
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}

	fmt.Println("main ends")
}

func testPreCompute() {
	startedAt := time.Now()
	defer fmt.Println(time.Since(startedAt))    //incorrect


	//*** correct code ****
	// defer func (start time.Time) {
	// 	fmt.Println(time.Since(start))
	// }(startedAt)

	time.Sleep(2*time.Second)
}

func double(x int) (result int) {
	defer func() {fmt.Printf("double(%d) = %d\n", x,result)}()
	return x + x
}

func triple(x int) (result int) {
	defer func() {result += x}()
	return double(4)
}