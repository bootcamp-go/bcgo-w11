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

import "fmt"

const (
	Minimum = "Minimum"
	Average = "Average"
	Maximum = "Maximum"
)

func main() {
	// fetch min function
	minFunc, err := operation(Minimum)
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
	averageFunc, err := operation(Average)
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
	maxFunc, err := operation(Maximum)
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


// ________________________________________________________________________________
type Operation func(...int) (int, string)

// operation returns the operation function and an error message
func operation(operator string) (op Operation, err string) {
	switch operator {
	case Minimum:
		op = minimum
	case Average:
		op = average
	case Maximum:
		op = maximum
	default:
		err = "Invalid operation"
	}
	return
}

// minimum returns the minimum value of the given values
func minimum(values ...int) (result int, err string) {
	var min int
	for ix, value := range values {
		if ix == 0 {
			min = value
			continue
		}
		if value < min {
			min = value
		}
	}
	result = min
	return
}

// average returns the average value of the given values
func average(values ...int) (result int, err string) {
	// check if values is empty
	if len(values) == 0 {
		err = "Empty values"
		return
	}

	var sum int
	for _, value := range values {
		sum += value
	}
	result = sum / len(values)
	return
}

// maximum returns the maximum value of the given values
func maximum(values ...int) (result int, err string) {
	var max int
	for ix, value := range values {
		if ix == 0 {
			max = value
			continue
		}
		if value > max {
			max = value
		}
	}
	result = max
	return
}