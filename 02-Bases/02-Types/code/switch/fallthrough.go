package main

func main() {
	// declare a variable i of type int
	i := 30

	// _________________________________________________________________
	// multiple paths (if statement)
	// if i%15 == 0 {
	// 	println("multiple of 15")
	// }
	// if i%5 == 0 {
	// 	println("multiple of 5")
	// }

	// _________________________________________________________________
	// unique path
	// >>> (if else statement)
	// if i%15 == 0 {
	// 	println("multiple of 15")
	// } else if i%5 == 0 {
	// 	println("multiple of 5")
	// } else {
	// 	println("no match")
	// }

	// >>> switch statement with fallthrough
	switch {
	case i%15 == 0:
		println("multiple of 15")
		fallthrough
	case i%5 == 0:
		println("multiple of 5")
	default:
		println("no match")
	}

	// other example
	age := 20
	switch {
	case age >= 200:
		println("are you inmortal?")
	case age >= 100:
		println("You are too old")
	case age >= 21 && age <= 100:
		println("You can drink")
		fallthrough
	case age >= 18:
		println("You can vote")
		fallthrough
	case age >= 16:
		println("You can drive")
	default:
		println("You are too young")
	}
}