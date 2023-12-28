package main

func main() {
	// counter (infinite loop)
	// var counter int
	// for true {
	// 	counter++
	// 	println(counter)
	// 	time.Sleep(time.Second)
	// }

	// counter (condition - break)
	var counter int
	for {
		// break condition
		if counter == 10 {
			break
		}
		counter++
		println(counter)
	}

	// _____________________________________________________________
	// counter (condition)
	// var counter int
	// for counter <= 10 {
	// 	counter++
	// 	println(counter)
	// }

	// counter (condition - skip first iteration)
	// 	var counter int
	// 	var i int
	// 	for counter <= 10 {
	// 		// check if its first iteration
	// 		i++
	// 		if i == 1 {
	// 			print("first iteration")
	// 			continue
	// 		}

	// 		counter++
	// 		println(counter)
	// 	}
}
