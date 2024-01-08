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

import "fmt"

func main() {
	// fetch dog function
	dogFunc, err := Animal(Dog)
	if err != "" {
		fmt.Println(err)
		return
	}
	result := dogFunc(10)
	fmt.Println("Dog:", result)

	// fetch cat function
	catFunc, err := Animal(Cat)
	if err != "" {
		fmt.Println(err)
		return
	}
	result = catFunc(10)
	fmt.Println("Cat:", result)

	// fetch hamster function
	hamsterFunc, err := Animal(Hamster)
	if err != "" {
		fmt.Println(err)
		return
	}
	result = hamsterFunc(10)
	fmt.Println("Hamster:", result)

	// fetch tarantula function
	tarantulaFunc, err := Animal(Tarantula)
	if err != "" {
		fmt.Println(err)
		return
	}
	result = tarantulaFunc(10)
	fmt.Println("Tarantula:", result)
}

var (
	Dog       = "dog"
	Cat       = "cat"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)

// Animal is a function that receives a string and returns a function and a string
func Animal(animal string) (grams AnimalGramsIntake, err string) {
	switch animal {
	case Dog:
		grams = DogGramsIntake
	case Cat:
		grams = CatGramsIntake
	case Hamster:
		grams = HamsterGramsIntake
	case Tarantula:
		grams = TarantulaGramsIntake
	default:
		err = "Animal not found"
	}
	return
}

// AnimalGramsIntake is a function that receives an int and returns a float64
type AnimalGramsIntake func(int) float64

// DogGramsIntake is a function that receives an int and returns a float64
func DogGramsIntake(quantity int) float64 {
	return float64(quantity) * 10000
}

// CatGramsIntake is a function that receives an int and returns a float64
func CatGramsIntake(quantity int) float64 {
	return float64(quantity) * 5000
}

// HamsterGramsIntake is a function that receives an int and returns a float64
func HamsterGramsIntake(quantity int) float64 {
	return float64(quantity) * 250
}

// TarantulaGramsIntake is a function that receives an int and returns a float64
func TarantulaGramsIntake(quantity int) float64 {
	return float64(quantity) * 150
}