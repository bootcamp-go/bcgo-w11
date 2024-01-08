/*
	Calcular promedio

	Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
	Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
	No se pueden introducir notas negativas.
*/

package main

import "app/unit-testing/second/average"

func main() {
	// declare variables
	var grades = []uint{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	// print average
	result, err := average.Average(grades...)
	if err != "" {
		println(err)
		return
	}
	println(result)
}

