package main

import (
	"io"
	"os"
	"fmt"
	"bytes"
	"time"
)

const debug = false

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

func main() {
	var w io.Writer
	fmt.Printf("w type is %T\n", w)
	// w.Write([]byte("hello"))  //panic: runtime error: invalid memory address or nil pointer dereference

	w = os.Stdout
	fmt.Printf("w type is %T\n", w)
	n, err := w.Write([]byte("hello\n"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Write ", n, " bytes")


	os.Stdout.Write([]byte("Stdout: hello\n"))

	var buf bytes.Buffer
	w = &buf
	fmt.Printf("w type is %T\n", w)
	w.Write([]byte("ni hao buffer\n"))
	fmt.Println(buf.String())

	io.Writer(&buf).Write([]byte("hello world\n"))
	fmt.Println(buf.String())

	new(bytes.Buffer).Write([]byte("hello\n"))

	w = nil
	fmt.Printf("w type is %T\n", w)
	// w.Write([]byte("hello"))  //panic: runtime error: invalid memory address or nil pointer dereference

	var x interface{} = time.Now()
	fmt.Printf("w type is %T\n", x)

	// var iY interface{} = []int{1,2,3}
	// fmt.Println(iY == iY)  //comparing uncomparable type []int

	var iZ interface{}
	fmt.Println(iZ)

	var buf2 io.Writer
	if debug {
		buf2 = new(bytes.Buffer) // enable collection of output
	}
	f(buf2) // NOTE: subtly incorrect!
	if debug {
	// ...use buf...
	}
}