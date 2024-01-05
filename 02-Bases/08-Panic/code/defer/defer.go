package main

import (
	"fmt"
	"time"
)

// OpenFile simulates a function that opens a file
func OpenFile() {
	fmt.Println("open file")
	time.Sleep(time.Second)
}

// ProcessFile simulates a function that processes a file
func ProcessFile() {
	fmt.Println("process file")
	time.Sleep(time.Second)
}

// CloseFile simulates a function that closes a file
func CloseFile() {
	fmt.Println("close file")
	time.Sleep(time.Second)
}

// testDefer simulates a function that uses defer
func testDefer() {
	// open file
	OpenFile()
	// close file
	defer CloseFile()
	// process file
	ProcessFile()

	// multiple defers
	defer func() {
		fmt.Println("Esto se ejecutara segundo")
		time.Sleep(time.Second)
	}()

	defer func() {
		fmt.Println("Esto se ejecutara primero")
		time.Sleep(time.Second)
	}()
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("Esto se va a ejecutar a pesar de un panic")
	}()

	testDefer()
	time.Sleep(time.Second)
	panic("Panic en main")
}
