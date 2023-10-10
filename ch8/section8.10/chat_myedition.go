package main

import (
	"io"
	"net"
	"log"
	"bufio"
)

var clients []net.Conn
var messages chan string

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		clients = append(clients, conn)
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)
	for input.Scan() {
		go clientWriter(c, input.Text())
	}
}

func broadcaster() {
	select {
		case message, ok := <- messages:
			if ok {
				for _, client := range clients {
					_, err := io.WriteString(client, message)
					if err != nil {
						return // e.g., client disconnected
					}
				}
			}
	}
}

func clientWriter(c net.Conn, text string) {
	// messages <- text
	for _, client := range clients {
		_, err := io.WriteString(client, text)
		if err != nil {
			return // e.g., client disconnected
		}
	}
}