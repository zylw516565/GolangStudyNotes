package main

import (
	"fmt"
	"errors"
	"syscall"
)

func main() {
	fmt.Println(errors.New("EOF") == errors.New("EOF"))
	
	var err error = syscall.Errno(2)
	fmt.Println(err.Error()) // "no such file or directory"
	fmt.Println(err) // "no such file or directory"
}