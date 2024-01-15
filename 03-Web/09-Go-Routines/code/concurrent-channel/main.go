package main

import (
	"time"
)

func SumOperation(n1, n2 int, ch chan<- int) {
	time.Sleep(500 * time.Millisecond)
	result := n1 + n2
	ch <- result
}

func main() {
	// channel
	ch := make(chan int)

	// i'm gonna read asynchronously from channel
	go func(ch <-chan int) {
		println("hello", <-ch)
	}(ch)

	// send channel
	ch <- 100  // not logical because nobody is receiving from channel (unbuffered channel) but if there is a receiver, it will receive this value and no deadlock
	// fmt.Println("hello", <-ch) // receive from channel

	start := time.Now()
	go SumOperation(1, 2, ch)
	go SumOperation(3, 4, ch)
	go SumOperation(5, 6, ch)
	go SumOperation(7, 8, ch)

	// time.Sleep(10 * time.Second)

	// receive from channel
	r1 := <-ch	// not necessarily 3
	r2 := <-ch	// not necessarily 7
	r3 := <-ch	// not necessarily 11
	r4 := <-ch	// not necessarily 15
	// r5 := <-ch  // deadlock
	println(r1, r2, r3, r4) //, r5)

	// wait for goroutines to finish
	time.Sleep(500 * time.Millisecond)

	end := time.Since(start)
	
	println("time:", end.Milliseconds(), "ms")
}