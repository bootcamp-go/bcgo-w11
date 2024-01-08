/*
	Ejercicio 5 - Calcular cantidad de alimento

	Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
	Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.

	Tienen los siguientes datos:
	- Perro 10 kg de alimento.
	- Gato 5 kg de alimento.
	- Hamster 250 g de alimento.
	- Tarántula 150 g de alimento.

	Se solicita:
	Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje
	(en caso que no exista el animal).
	Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

	Ejemplo:
	const (
	dog    = "dog"
	cat    = "cat"
	)

	animalDog, msg := animal(dog)
	animalCat, msg := animal(cat)

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
*/

package main

import (
	"app/unit-testing/fifthy/pets"
	"fmt"
)

func main() {
	// fetch dog function
	dogFunc, err := pets.Orchestrator(pets.Dog)
	if err != "" {
		fmt.Println(err)
		return
	}
	result := dogFunc(10)
	fmt.Println("Dog:", result)

	// fetch cat function
	catFunc, err := pets.Orchestrator(pets.Cat)
	if err != "" {
		fmt.Println(err)
		return
	}
	result = catFunc(10)
	fmt.Println("Cat:", result)

	// fetch hamster function
	hamsterFunc, err := pets.Orchestrator(pets.Hamster)
	if err != "" {
		fmt.Println(err)
		return
	}
	result = hamsterFunc(10)
	fmt.Println("Hamster:", result)

	// fetch tarantula function
	tarantulaFunc, err := pets.Orchestrator(pets.Tarantula)
	if err != "" {
		fmt.Println(err)
		return
	}
	result = tarantulaFunc(10)
	fmt.Println("Tarantula:", result)
}