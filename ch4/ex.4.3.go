//练习 4.3： 重写reverse函数， 使用数组指针代替slice。

package main

import (
	"fmt"
)


var a = [6]int{0, 1, 2, 3, 4, 5}

func main() {
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
}

func reverse(s *[6]int) {
	for i, j := 0, len(*s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}