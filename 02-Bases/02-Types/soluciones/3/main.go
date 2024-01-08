/*
	Ejercicio 3 - A qué mes corresponde

	Realizar una aplicación que reciba  una variable con el número del mes.

	Según el número, imprimir el mes que corresponda en texto.
	¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
	Ej: 7, Julio.
*/

package main

func main() {
	// declare variables
	var month int = 7

	// print month
	switch month {
	case 1:
		println("January")
	case 2:
		println("February")
	case 3:
		println("March")
	case 4:
		println("April")
	case 5:
		println("May")
	case 6:
		println("June")
	case 7:
		println("July")
	case 8:
		println("August")
	case 9:
		println("September")
	case 10:
		println("October")
	case 11:
		println("November")
	case 12:
		println("December")
	default:
		println("Invalid month")
	}
}