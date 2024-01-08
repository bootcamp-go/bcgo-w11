/*
	Ejercicio 3 - Declaración de variables


	Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.

	Necesita ayuda para:

	Detectar cuáles de estas variables que declaró el alumno son correctas.
	Corregir las incorrectas.
		var 1firstName string

		var lastName string

		var int age

		1lastName := 6

		var driver_license = true

		var person height int

		childsNumber := 2

*/

package main

func main() {
	// declare variables
	var firstName string
	var lastName string
	var age int
	var driverLicense bool
	var personHeight int
	var childsNumber int

	// initialize variables
	firstName = "Juan"
	lastName = "Perez"
	age = 30
	driverLicense = true
	personHeight = 180
	childsNumber = 2

	// print variables
	println("FirstName:", firstName)
	println("LastName:", lastName)
	println("Age:", age)
	println("Has driver license:", driverLicense)
	println("Height:", personHeight)
	println("Childs number:", childsNumber)
}