package main

import (
	"io"
	"os"
	"fmt"
	"bytes"
)

import ("time")
type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP3", "WAV"
}

type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP4", "WMV"
	Resolution() (x, y int)
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

type IntSet struct { /* ... */ }
func (s *IntSet) String() string {
	fmt.Println("(s *IntSet) String()")
	return ""
}

func main() {
	var iW io.Writer
	fmt.Printf("%T\n", iW)
	iW = os.Stdout
	fmt.Printf("%T\n", iW)
	iW = new(bytes.Buffer)
	fmt.Println(iW)
	fmt.Printf("%T\n", iW)
	// iW = time.Second    //time.Duration does not implement io.Writer (missing method Write)

	fmt.Println("************************")
	var iRwc io.ReadWriteCloser
	fmt.Printf("%T\n", iRwc)
	iRwc = os.Stdout
	fmt.Printf("%T\n", iRwc)
	// iRwc = new(bytes.Buffer)  //*bytes.Buffer does not implement io.ReadWriteCloser (missing method Close)
	// fmt.Printf("%T\n", iRwc)
	fmt.Println("************************")
	iW = iRwc
	fmt.Printf("%T\n", iW)
	fmt.Printf("%T\n", iRwc)
	// iRwc = iW  //io.Writer does not implement io.ReadWriteCloser (missing method Close)
	fmt.Println("************************")
	// var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
	var s IntSet
	s.String()

	var _ fmt.Stringer = &s
	// var _ fmt.Stringer = s  //IntSet does not implement fmt.Stringer (method String has pointer receiver)

	fmt.Println("************************")
	os.Stdout.Write([]byte("hello\n"))
	// os.Stdout.Close()

	var iW2 io.Writer
	iW2 = os.Stdout
	iW2.Write([]byte("world\n"))
	// iW2.Close()  //iW2.Close undefined (type io.Writer has no field or method Close)

	fmt.Println("************************")
	var iEmpty interface{}
	iEmpty = nil
	iEmpty = os.Stdout
	_ = iEmpty

	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	_   = any
	fmt.Println(any)

	fmt.Println("************************")
	var iW3 io.Writer = new(bytes.Buffer)
	_ = iW3
	var _ io.Writer = (*bytes.Buffer)(nil)
}