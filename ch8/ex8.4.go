package main

import (
	"net"
	"log"
	"fmt"
	"time"
	"bufio"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}


	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 3 * time.Second)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}