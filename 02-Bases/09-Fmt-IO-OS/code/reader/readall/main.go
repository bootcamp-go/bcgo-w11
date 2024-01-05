package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// generate a reader from a string
	text := "lorem ipsum dolor sit amet"
	reader := strings.NewReader(text)

	// read 30 bytes from the reader
	buf := make([]byte, 30)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read:", string(buf[:n]))

	// restart the cursor to the beginning of the reader	
	reader.Seek(0, io.SeekStart)

	buf = make([]byte, 30)
	n, err = reader.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data read:", string(buf[:n]))
}