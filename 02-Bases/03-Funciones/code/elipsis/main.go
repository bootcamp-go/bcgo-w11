package main

func main() {
	a := [3]float64{7.0, 8.5, 9.1} // array
	b := a[:]                      // slice

	total := addition(b...)
	println(total)

	// c := [...]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // array
}

func addition(values ...float64) float64 {
	var total float64
	for _, value := range values {
		total += value
	}
	return total
}