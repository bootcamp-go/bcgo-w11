/*
	Ejercicio 4 - Qué edad tiene...

	Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.


	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}


	Por otro lado también es necesario:

	Saber cuántos de sus empleados son mayores de 21 años.
	Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	Eliminar a Pedro del mapa.
*/

package main

func main() {
	// declare map
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// print Benjamin's age
	benjaminAge, ok := employees["Benjamin"]
	if !ok {
		println("Benjamin is not an employee")
		return
	}
	println("Benjamin's age:", benjaminAge)

	// count employees older than 21
	var olderThan21 int
	for _, age := range employees {
		if age > 21 {
			olderThan21++
		}
	}
	println("Employees older than 21:", olderThan21)

	// add new employee
	employees["Federico"] = 25

	// delete employee
	delete(employees, "Pedro")
}