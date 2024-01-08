/*
	Ejercicio 4 - Calcular estadísticas


	Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso.
	Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
	Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
	y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y
	devuelva el cálculo que se indicó en la función anterior.
	Ejemplo:
	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/

package main

import (
	"app/unit-testing/fourth/operations"
	"fmt"
)

func main() {
	// fetch min function
	minFunc, err := operations.Orchestrator(operations.MinimumOperator)
	if err != "" {
		fmt.Println(err)
		return
	}
	result, err := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println("Minimum:", result)

	// fetch average function
	averageFunc, err := operations.Orchestrator(operations.AverageOperator)
	if err != "" {
		fmt.Println(err)
		return
	}
	result, err = averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println("Average:", result)

	// fetch max function
	maxFunc, err := operations.Orchestrator(operations.MaximumOperator)
	if err != "" {
		fmt.Println(err)
		return
	}
	result, err = maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println("Maximum:", result)
}