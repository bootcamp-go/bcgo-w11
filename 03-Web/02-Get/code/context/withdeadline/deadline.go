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
	// deadline
	// - date with time.Date(year, month, day, hour, min, sec, nsec, location)
	// deadline := time.Date(2024, 1, 9, 16, 46, 40, 0, time.FixedZone("UTC-3", -3*60*60))
	// - date with time.Now().Add() for add a duration to the current time
	deadl := time.Now().Add(5 * time.Second)

	// context
	ctx := context.Background()
	// - context.WithDeadline() returns a copy of parent with a new Done channel
	ctx, cancelFunc := context.WithDeadline(ctx, deadl)
	defer cancelFunc()

	// Simulate a long running operation in a goroutine
	go SimulateOperation(ctx)

	// Wait for a key press to continue main
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println("End of the program")
}
