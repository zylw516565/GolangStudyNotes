package main

import (
	"fmt"
	"time"
	"sync"
	"unsafe"
	"strconv"
	"math"
	"runtime/debug"
)

type hmap struct {
	count int 
	flags uint8
	B     uint8  
	hash0 uint32
	buckets    unsafe.Pointer 
	oldbuckets unsafe.Pointer 
}

func main() {

	// concurrentReadWriteMap()
	// readWriteMapWithRWLock()
	testMapCap()
	select{}
}

func testMapCap() {
	m := make(map[string]string)
	c, b := getInfo(m)
	fmt.Println("count: ", c, "b: ", b)
	for i := 0; i < 10000; i++ {
			m[strconv.Itoa(i)] = strconv.Itoa(i)
			if i%200 == 0 {
					c, b := getInfo(m)
					cap := math.Pow(float64(2), float64(b))
					fmt.Printf("count: %d, b: %d, load: %f\n", c, b, float64(c)/cap)
			}
	}
	println("开始删除------")
	for i := 0; i < 10000; i++ {
			delete(m, strconv.Itoa(i))
			if i%200 == 0 {
					c, b := getInfo(m)
					cap := math.Pow(float64(2), float64(b))
					fmt.Println("count: ", c, "b:", b, "load: ", float64(c)/cap)
			}
	}


	debug.FreeOSMemory()
	c, b = getInfo(m)
	fmt.Println("释放后: ", "count: ", c, "b:", b)
}

func getInfo(m map[string]string) (int, int) {
	point := (**hmap)(unsafe.Pointer(&m))
	value := *point
	return value.count, int(value.B)
}

func concurrentReadWriteMap() {
	m := make(map[string]int)

	go func() {
		for {
			m["a"]++
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			_ = m["b"]
			time.Sleep(time.Microsecond)
		}
	}()
}

func readWriteMapWithRWLock() {
	var rw sync.RWMutex
	m := make(map[string]int)

	go func() {
		for {
			rw.Lock()
			m["a"]++
			rw.Unlock()
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			rw.RLock()
			_ = m["b"]
			rw.RUnlock()
			time.Sleep(time.Microsecond)
		}
	}()
}