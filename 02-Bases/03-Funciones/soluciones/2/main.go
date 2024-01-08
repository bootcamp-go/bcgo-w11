/*
	Calcular promedio

	Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
	Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
	No se pueden introducir notas negativas.
*/

package main

func main() {
	// declare variables
	var grades = []uint{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	// print average
	println("Average:", average(grades...))
}

func average(grades ...uint) uint {
	// check if len is 0
	if len(grades) == 0 {
		return 0
	}

	// calculate average
	var sum uint
	for _, grade := range grades {
		sum += grade
	}
	return sum / uint(len(grades))
}