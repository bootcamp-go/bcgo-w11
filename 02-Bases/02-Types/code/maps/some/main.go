package main

func main() {
	// map is a dynamic type
	// - key: clave
	// - value: valor que almacena la clave
	// names := map[string]uint{}
	names := make(map[string]uint)

	// write operation
	// - create
	names["John"] = 18
	// - update
	names["John"] = 19

	// read operation
	age := names["John"]
	println(age)
}