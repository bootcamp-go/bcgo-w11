/*
	Ejercicio 2 - Clima


	Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares.

	1. Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
	2. Imprimí los valores de las variables en consola.
	3. ¿Qué tipo de dato le asignarías a las variables?
*/

package main

func main() {
	// declare variables
	var temperature float32
	var humidity float32
	var preassure float32

	// initialize variables
	temperature = 25.5
	humidity = 0.5
	preassure = 1.2

	// print variables
	println("Temperature:", temperature)
	println("Humidity:", humidity)
	println("Preassure:", preassure)
}