package main

import (
	"net"
	"log"
	"fmt"
	"time"
	"bufio"
	"strings"
)

const Interval = 10 * time.Second

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
	timeout := time.NewTicker(10 * time.Second)
	defer timeout.Stop()

	fmt.Println("handleConn 1")
	go handleEcho(c, timeout)

	select {
		case <- timeout.C:
			fmt.Println("Timeout: shutdown conn ")
			return
	}
}

func handleEcho(c net.Conn, timeout *time.Ticker) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 3 * time.Second, timeout)
	}
}

func echo(c net.Conn, shout string, delay time.Duration, timeout *time.Ticker) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	timeout.Reset(Interval)
}