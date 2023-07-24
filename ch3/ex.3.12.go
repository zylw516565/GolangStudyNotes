package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isDisordered("12345", "123456789"))
	fmt.Println(isDisordered("123456789", "123456789"))
	fmt.Println(isDisordered("912345678", "123456789"))
	fmt.Println(isDisordered("912345678,", "123456789"))
}

func isDisordered(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, r := range s1 {
		if strings.Count(s1, string(r)) != strings.Count(s2, string(r)) {
			return false
		}
	}

	if s1 == s2 {
		return false
	}

	return true
}