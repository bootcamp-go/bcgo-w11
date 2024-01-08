/*
	Ejercicio 3 - Calcular salario

	Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

	Categoría C, su salario es de $1.000 por hora.
	Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
	Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
	Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.
*/

package main

import "app/unit-testing/third/salary"

func main() {
	// declare variables
	var hours float64 = 160.0
	category := salary.CategoryA

	// fetch salary
	salary, err := salary.PerHourAndCategory(hours, category)
	if err != "" {
		println(err)
		return
	}

	// print salary
	println("Salary:", salary)
}