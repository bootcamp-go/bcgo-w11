package main

import "fmt"

// funcion1 call to funcion2
func funcion1() {
	funcion2()
	fmt.Println("Despues de llamado a funcion 2")
}

// funcion2 call to funcion3
func funcion2() {
	funcion3()
	fmt.Println("Despues de llamado a funcion 3")
}

// funcion3 panic
func funcion3() {
	// simulando un panic
	panic("Panic en funcion3")
}

func main() {
	// start main
	fmt.Println("Iniciando programa")

	funcion1()

	// end main
	fmt.Println("Finalizando programa")
}
