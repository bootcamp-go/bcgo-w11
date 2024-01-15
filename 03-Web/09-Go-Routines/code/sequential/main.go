package main

import "time"

func SumOperation(n1, n2 int) {
	time.Sleep(500 * time.Millisecond)
	result := n1 + n2
	println(result)
}

func main() {
	start := time.Now()
	SumOperation(1, 2)
	SumOperation(3, 4)
	SumOperation(5, 6)
	SumOperation(7, 8)
	end := time.Since(start)
	println("time:", end.Milliseconds(), "ms")	
}