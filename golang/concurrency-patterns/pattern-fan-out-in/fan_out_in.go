package main

import (
	"fmt"
	"sync"
	"time"
)

// FanOutFanIn demonstrates a pattern where work is distributed to multiple
// goroutines (fan-out) and then results are collected back (fan-in).
func main() {
	// Create input channel
	jobs := make(chan int)

	// Start multiple workers (fan-out)
	numWorkers := 3
	results := make(chan int)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	go func() {
		defer close(jobs)
		for i := 1; i <= 10; i++ {
			jobs <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Wait for all workers to finish in a separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results (fan-in)
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// worker processes jobs and sends results
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate work
		time.Sleep(200 * time.Millisecond)
		results <- job * 2
		fmt.Printf("Worker %d processed job %d\n", id, job)
	}
} 