package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to server.
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// listen for reply.
	go func() {
		for {
			// read bytes from connection.
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}

			data := string(buf[:n])
			fmt.Println("Server reply:", data)
		}
	}()

	// Run loop forever (or until ctrl-c).
	for {
		// Read input from stdin.
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		// Send to server.
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}