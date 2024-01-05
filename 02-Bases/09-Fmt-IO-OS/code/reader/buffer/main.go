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

	// read by 8 bytes per time / iteration
	buf := make([]byte, 8)
	for {
		// read bytes
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}

		fmt.Println("Data read:", string(buf[:n]))
	}
}