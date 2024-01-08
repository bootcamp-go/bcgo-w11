/*
	Ejercicio 2 - Préstamo


	Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
	Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
	Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
	Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.

	Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.

	Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

package main

func main() {
	// declare variables
	var age int = 23
	var employed bool = true
	var yearsWorking int = 2
	var salary int = 100000

	// check if client can get a loan
	if age > 22 && employed && yearsWorking > 1 {
		if salary > 100000 {
			println("You can get a loan without interest")
		} else {
			println("You can get a loan")
		}
	} else {
		println("You can't get a loan")
	}
}
