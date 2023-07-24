package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Println(commaWithBuffer("12345"))
	fmt.Println(commaWithBuffer("123456789"))

}

func commaWithBuffer(s string) string {
	var buf bytes.Buffer
	for i, n := len(s)-1, 0; i >= 0; i-- {
		n++
		buf.WriteByte(s[i])
		if 3 == n && 0 != i {
			buf.WriteByte(',')
			n = 0
		}
	}

	b := buf.Bytes()
	reverse(b)
	return string(b)
}

func reverse(s []byte) {
	for i,j := 0, len(s)-1; i < j; i,j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}