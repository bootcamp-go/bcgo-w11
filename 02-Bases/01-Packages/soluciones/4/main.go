/*

	Ejercicio 4 - Tipos de datos

	Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go, pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:

	Verificar su código y realizar las correcciones necesarias.

		var lastName string = "Smith"

		var age int = "35"

		boolean := "false"

		var salary string = 45857.90

		var firstName string = "Mary"

*/

package main

func main() {
	// Declarar las variables
	var lastName string = "Smith"
	var age int = 35
	var boolean bool = false
	var salary float32 = 45857.90
	var firstName string = "Mary"

	// Imprimir en consola el valor de cada una de las variables
	println("lastName:", lastName)
	println("age:", age)
	println("boolean:", boolean)
	println("salary:", salary)
	println("firstName:", firstName)
}