package main

import (
	"bufio"
	"os"
)

func main() {
	// create reader from stdin
	rd := bufio.NewReader(os.Stdin)
	
	// read with delimiter '\n'
	for {
		// read string until '\n'
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}

		// print line
		println(line)
	}
}