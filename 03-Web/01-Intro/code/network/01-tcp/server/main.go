package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Listen on TCP port 2000 on all interfaces.
	ln, _ := net.Listen("tcp", ":2000")

	// Accept connection on port.
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Run loop forever (or until ctrl-c).
	for {
		// read bytes from connection.
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		// request
		data := string(buf[:n])
		data = strings.Trim(data, "\r\n")
		fmt.Println("Client request:", data)

		// response
		switch data {
		case "hello":
			conn.Write([]byte("world"))
		case "bye":
			conn.Write([]byte("bye"))
		default:
			conn.Write([]byte("..."))
		}
	}
}