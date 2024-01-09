package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

// SimulateOperation simulates a long running operation
func SimulateOperation(ctx context.Context) {
	fmt.Println("Simulating operation...")

	// Simulate a long running operation
	for i := 0; i < 5; i++ {
		select {
		// - time.After() returns a channel that blocks for the specified duration
		case <-time.After(1 * time.Second):
			fmt.Println("Doing something...")
			// - ctx.Done() returns a channel that's closed when work done on behalf of this context should be canceled
		case <-ctx.Done():
			fmt.Println("Operation canceled")
			return
		}
	}
	fmt.Println("Operation finished")
}

func main() {
	// context
	// - context.WithTimeout() returns a copy of parent with a new Done channel
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)

	// Simulate a long running operation in a goroutine
	go SimulateOperation(ctx)

	// Wait for a key press to continue main
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	// - cancelFunc() closes the Done channel, indicating to any child goroutines that the context has been canceled
	cancelFunc()
	fmt.Println("End of the program")
}
