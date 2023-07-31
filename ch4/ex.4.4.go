//练习 4.4： 编写一个rotate函数， 通过一次循环完成旋转。

package main

import (
	"fmt"
)

var a = []int{0, 1, 2, 3, 4, 5}

func main() {
	str := "ni hao ya"
	sliceStr := str[:]
	fmt.Println(sliceStr)
	fmt.Printf("sliceStr: %T\n", sliceStr)
	sliceStr = "hello"
	fmt.Println(str)

	fmt.Println(rotate(a, 0))
	fmt.Println(rotate(a, 1))
	fmt.Println(rotate(a, 3))

	fmt.Println(rotate(a, -1))
	fmt.Println(rotate(a, -3))
}

// count为+,向右旋转;  count为-,向左旋转
func rotate(array []int, count int) []int {
	dest := make([]int, len(array))
	steps := count % len(array)
	for i, a := range array {
		if i + steps >= 0 && i + steps < len(array) {
			dest[i + steps] = a
		} else {
			if i + steps >= len(array) {
				dest[i + steps - len(array)] = a
			}

			if i + steps < 0 {
				dest[i + steps + len(array)] = a
			}
		}
	}

	return dest
}