//练习 4.6： 编写一个函数， 原地将一个UTF-8编码的[]byte类型的slice中相邻的空格(参考unicode.IsSpace) 替换成一个空格返回

package main

import (
	"os"
	"io"
	"fmt"
	"bufio"
	"unicode"
	"unicode/utf8"
)

func main() {
	charCount2()
}

func charCount2() {
	counts := make(map[rune]int) // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0 // count of invalid UTF-8 characters
	var letters, digits int
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		if unicode.IsLetter(r) {
			letters++
		}

		if unicode.IsNumber(r) {
			digits++
		}
	}
	fmt.Printf("letters\tcount:%d\n", letters)
	fmt.Printf("digits\tcount:%d\n", digits)

	fmt.Printf("rune\tcount\n")

	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")

	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}