package main

import (
	"fmt"
	"os"
)

// runProgram simulates a run of a program that opens a file
func runProgram() (err error) {
	// defer function to recover from panic
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
			panic("panic en defer func()")
		}
	}()

	// open file
	_, err = os.Open("sin-archivo.txt")
	if err != nil {
		panic(err)
	}

	return
}

// runProgram2 simulates a run of a program that opens a file
func runProgram2() {
	// defer function to recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado en main", r)
		}
	}()

	// call to runProgram
	err := runProgram()
	if err != nil {
		fmt.Println(err)
	}

	// this line is executed
	fmt.Println("Se ejecuta de igual manera")
}

func main() {
	// start main
	fmt.Println("Iniciando")

	// process
	runProgram2()

	// end main
	fmt.Println("Finalizando")
}
