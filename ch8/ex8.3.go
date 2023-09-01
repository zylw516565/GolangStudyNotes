package main

import (
	"os"
	"io"
	"net"
	"log"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	mustCopy(conn, os.Stdin)
	log.Println("main goroutine wait")

	switch conn2 := conn.(type) {
		case *net.TCPConn:
			fmt.Println("CloseWrite")
			conn2.CloseWrite()  //?  也会导致后台goroutine的io.Copy(os.Stdout, conn)返回, Why?
		default:
			fmt.Println("not surport")
	}

	// conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}