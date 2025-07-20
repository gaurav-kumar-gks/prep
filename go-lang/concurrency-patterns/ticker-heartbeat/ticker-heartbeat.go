package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Ticker / Heartbeat ===")

	// Ticker pattern
	rateLimiter := time.NewTicker(500 * time.Millisecond)
	defer rateLimiter.Stop()
	// Timeout pattern
	timeout := time.After(2 * time.Second)
	// Done
	done := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()
		

	// Use a label to break out of the outer loop
heartbeatLoop:
	for {
		select {
		case <-done:
			fmt.Println("Operation cancelled")
			break heartbeatLoop // Break out of the labeled loop
		case <-rateLimiter.C:
			go func() {
				fmt.Printf("Rate limited operation")
			}()
		case <-timeout:
			fmt.Println("Timeout reached, stopping heartbeat.")
			break heartbeatLoop // Break out of the labeled loop
		}
	}

	// Wait for goroutines to complete
	time.Sleep(6 * time.Second)
}
