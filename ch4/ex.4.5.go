//练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。

package main

import (
	"fmt"
)

var str = []string{"abcccccdef", "abcdddddd", "aaaabbbb"}

func main() {
	fmt.Println(str)
	fmt.Println(dedup(str))
}

func dedup(all []string) []string {

	for k := range all {
		s := all[k]
		
		if len(s) < 2 {
			return all
		}

		i, j := 0, 1
		for i < len(s) - 1 && j < len(s) {
			if s[i] != s[j] {
				i++
				j++
			} else {
				s = s[:i] + s[j:]
			}
		}

		all[k] = s[:i+1]
	}

	return all
}