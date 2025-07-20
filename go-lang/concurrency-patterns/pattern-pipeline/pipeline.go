package main

import (
	"fmt"
	"time"
)

// Pipeline demonstrates a simple pipeline pattern where data flows through
// multiple stages of processing, with each stage running in its own goroutine.
func main() {
	// Create input channel
	numbers := make(chan int)

	// Start pipeline stages
	squares := square(numbers)
	cubes := cube(squares)
	results := multiply(cubes)

	// Send numbers to pipeline
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Final result: %d\n", result)
	}
}

// square takes a number and returns its square
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// cube takes a number and returns its cube
func cube(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n * n
		}
	}()
	return out
}

// multiply takes a number and multiplies it by 2
func multiply(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
} 