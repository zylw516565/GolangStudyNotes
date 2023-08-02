//练习 4.6： 编写一个函数， 原地将一个UTF-8编码的[]byte类型的slice中相邻的空格(参考unicode.IsSpace) 替换成一个空格返回

package main

import (
	"os"
	"fmt"
	"bufio"
	// "unicode"
	// "unicode/utf8"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters
	invalid := 0 // count of invalid UTF-8 characters

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		counts[word]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nword\tcount\n")

	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}