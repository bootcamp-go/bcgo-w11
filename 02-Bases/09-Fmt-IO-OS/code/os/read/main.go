package main

import (
	"fmt"
	"os"
)

func main() {
	// open file
	f, err := os.Open("./code/os/read/file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	buf := make([]byte, 8)
	for {
		// read file
		n, err := f.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}