package main

import (
	"math/rand"
)

// LoadRandomNumbers loads random numbers
func LoadRandomNumbers(ch chan<- int) {
	// send indeterminate number of random numbers
	for i := 0; i < rand.Intn(10); i++ {
		ch <- rand.Intn(100)
	}

	// close channel
	close(ch)
}

func main() {
	// channel
	ch := make(chan int)

	// initialize randomizer
	go LoadRandomNumbers(ch)

	for {
		v, ok := <-ch
		if !ok {
			println("channel closed")
			break
		}

		println(v)
	}

	// for v := range ch {
	// 	println(v)
	// }

	// // send to a closed channel (panic)
	// go func(ch chan<- int) {
	// 	ch <- 100
	// }(ch)

	// // receive from channel
	// println(<-ch)
}