package main

import "time"

func SumOperation(n1, n2 int) {
	time.Sleep(500 * time.Millisecond)
	result := n1 + n2
	println(result)
}

func main() {
	start := time.Now()
	go SumOperation(1, 2)
	go SumOperation(3, 4)
	go SumOperation(5, 6)
	go SumOperation(7, 8)

	// wait for goroutines to finish
	time.Sleep(3 * time.Second)

	end := time.Since(start)
	println("time:", end.Milliseconds(), "ms")
}