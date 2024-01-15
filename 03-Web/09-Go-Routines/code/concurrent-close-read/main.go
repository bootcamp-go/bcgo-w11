package main

func main() {
	ch := make(chan int)
	close(ch)

	// send to a closed channel (panic)
	// ch <- 100

	// receive from channel (zero value)
	println(<-ch)
	println(<-ch)
	println(<-ch)
}