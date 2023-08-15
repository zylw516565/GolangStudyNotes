package main

import (
	"fmt"
)

func main () {
	// arr := []int{1, 2, 3}
	// for i, _ := range arr {
	// 	println(i)
	// }

	test1()
	test2()
	fmt.Println()
	test3()

	clearSlice()
	clearSlice2()

	fmt.Println("**************")
	str := "nihao,世界,hello"
	for _, val := range str {
		fmt.Printf("%q\n", rune(val))
	}
	fmt.Println("**************")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%q\n", str[i])
	}

	var r rune = '世'
	fmt.Println(r)
}

func test1() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
}

func test2() {
	arr := []int{1,2,3}
	newArr := []*int{}

	for _, val := range arr {
		newArr = append(newArr, &val)
	}

	for _, v := range newArr {
		fmt.Println(*v)
	}
}

func test3() {
	arr := []int{1,2,3}
	newArr := []*int{}

	for i, _ := range arr {
		newArr = append(newArr, &arr[i])
	}

	for _, v := range newArr {
		fmt.Println(*v)
	}
}

func clearSlice() {
	arr := []int{1, 2, 3}
	for _, val := range arr {
		val = 0
		_ = val
	}

	fmt.Println(arr)
}

func clearSlice2() {
	arr := []int{1, 2, 3}
	for i, _ := range arr {
		arr[i] = 0
	}

	fmt.Println(arr)
}