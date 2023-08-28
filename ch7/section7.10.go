package main

import (
	"io"
	"os"
	"fmt"
	"bytes"
)

func main() {
	var w io.Writer
	w = os.Stdout
	x := w.(*os.File)
	// b := w.(*bytes.Buffer)  //panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	// _ = b
	fmt.Printf("%T, %#[1]v\n", x)
	a := w.(io.Writer)
	fmt.Printf("%T, %#[1]v\n", a)
	// b := w.(error)   //panic: interface conversion: *os.File is not error: missing method Error
	// fmt.Printf("%T, %#[1]v\n", b)

	var w2 io.Writer
	w2 = os.Stdout
	rw2 := w2.(io.ReadWriter) // success: *os.File has both Read and Write
	w2.Write([]byte("hello "))
	rw2.Write([]byte("world !\n"))
	var buf []byte
	rw2.Read(buf)

	w2 = rw2
	w2 = rw2.(io.Writer)

	var w3 io.Writer
	w3 = os.Stdout
	f, ok := w3.(*os.File)
	fmt.Printf("%T, %#[1]v, %v\n", f, ok)
	b, ok := w3.(*bytes.Buffer)
	fmt.Printf("%T, %#[1]v, %v\n", b, ok)
	// fmt.Printf("%T, %#[1]v\n", b)
	if w3, ok2 := w3.(*os.File); ok2 {
		fmt.Printf("%T, %#[1]v, %v\n", w3, ok2)
	}
}