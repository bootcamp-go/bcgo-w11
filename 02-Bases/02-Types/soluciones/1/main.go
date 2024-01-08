/*
	Ejercicio 1 - Letras de una palabra
	La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:

	Crear una aplicación que reciba una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
	Luego, imprimir cada una de las letras.
*/

package main

func main() {
	// declare word
	word := "hello"

	// print word length
	println("Word length:", len(word))

	// print each letter
	for _, letter := range word {
		println(string(letter))
	}
}