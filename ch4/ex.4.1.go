package main

import (
	"fmt"
)

const Size  = 32
var   aByte byte = 1
func main() {

	code1 := [32]byte{0x2d, 0x71, 0x16, 0x42, 0xb7, 0x26, 0xb0, 0x44, 0x01, 0x62, 0x7c, 0xa9, 0xfb, 0xac, 0x32, 0xf5, 0xc8, 0x53, 0x0f, 0xb1, 0x90, 0x3c, 0xc4, 0xdb, 0x02, 0x25, 0x87, 0x17, 0x92, 0x1a, 0x48, 0x81}
	var code2 [Size]byte
	fmt.Println(diffBitCount(code1, code2))

	code5 := [32]byte{0x01}
	fmt.Println(diffBitCount(code5, code2))


	var code3 [Size]byte
	var code4 [Size]byte
	fmt.Println(diffBitCount(code3, code4))
}

func diffBitCount(code1, code2 [Size]byte) int {
	count := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < 8; j++ {
			if code1[i] & byte(1 << j) != code2[j] & byte(1 << j) {
				count++
			}
		}
	}

	return count
}