package main

import (
	"fmt"
	"bytes"
	"strings"
)

func main() {
	fmt.Println(commaFloat("12345"))
	fmt.Println(commaFloat("123456789"))

	fmt.Println(commaFloat("+123456789.08960"))
	fmt.Println(commaFloat("+123456789.000000001"))

	fmt.Println(commaFloat("-123456789.08960"))
	fmt.Println(commaFloat("-123456789.000000001"))
}

func commaFloat(s string) string {
	var buf bytes.Buffer;  var intPart, decimal string
	if strings.Contains(s, ".") {
		intPart = s[:strings.Index(s, ".")]
		decimal = s[strings.Index(s, "."):]
	}else {
		intPart = s
	}

	for i, n := len(intPart)-1, 0; i >= 0; i-- {
		n++
		buf.WriteByte(intPart[i])
		if 3 == n && i >= 1 {
			if intPart[i-1] != '+' && intPart[i-1] != '-' {
				buf.WriteByte(',')
			}

			n = 0
		}
	}

	b := buf.Bytes()
	reverse(b)
	return string(b) + decimal
}

func reverse(s []byte) {
	for i,j := 0, len(s)-1; i < j; i,j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}