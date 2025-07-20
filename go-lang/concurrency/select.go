package concurrency

import (
	"context"
	"fmt"
	"time"
)

/*
Select in Go

Basic Usage:
- Select statement
- Default case
- Multiple cases
- Channel operations
- Timeout handling

Advanced Usage:
- Non-blocking operations
- Channel multiplexing
- Channel cancellation
- Channel timeouts
- Channel priorities
*/

// DemonstrateSelect shows select statement usage
func DemonstrateSelect() {
	fmt.Println("\n=== Select Statement ===")
	
	// Create channels
	ch1 := make(chan int)
	ch2 := make(chan string)
	
	// Start goroutines to send values
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- 42
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "hello"
	}()
	
	// Select from multiple channels
	select {
	case x := <-ch1:
		fmt.Printf("Received from ch1: %d\n", x)
	case s := <-ch2:
		fmt.Printf("Received from ch2: %s\n", s)
	case <-time.After(150 * time.Millisecond):
		fmt.Println("Timeout")
	default:
		fmt.Println("No value available")	
	}
}

// DemonstrateSelectCancellation shows select with cancellation
func DemonstrateSelectCancellation() {
	fmt.Println("\n=== Select with Cancellation ===")
	
	// Create channels
	ch := make(chan int)
	// make(chan struct{}) is used for signalling 
	// and is more efficient than using a channel of bool
	done := make(chan struct{})
	
	// Start a goroutine to send values
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			select {
			case <-done:
				fmt.Println("Cancelled")
				return
			case ch <- i:
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()
	
	// Receive values until cancellation
	for i := 0; i < 3; i++ {
		select {
		case x := <-ch:
			fmt.Printf("Received: %d\n", x)
		case <-time.After(200 * time.Millisecond):
			fmt.Println("Timeout")
			close(done)
			return
		}
	}
	
	// Cancel the goroutine
	close(done)
}

// DemonstrateSelectMultiplexing shows select with channel multiplexing
func DemonstrateSelectMultiplexing() {
	fmt.Println("\n=== Select with Channel Multiplexing ===")
	
	// Create channels
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	
	// Start goroutines to send values
	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			ch1 <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(ch2)
		for i := 0; i < 3; i++ {
			ch2 <- i * 2
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(ch3)
		for i := 0; i < 3; i++ {
			ch3 <- i * 3
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	// Multiplex channels
	for {
		select {
		case x, ok := <-ch1:
			if !ok {
				// channel is made nil
				// so that it won't be selected again
				// and will not block the select
				ch1 = nil
				continue
			}
			fmt.Printf("Received from ch1: %d\n", x)
		case x, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Printf("Received from ch2: %d\n", x)
		case x, ok := <-ch3:
			if !ok {
				ch3 = nil
				continue
			}
			fmt.Printf("Received from ch3: %d\n", x)
		default:
			fmt.Println("No value available, doing some work")
			time.Sleep(10 * time.Millisecond)
		}
		
		// Exit when all channels are closed
		if ch1 == nil && ch2 == nil && ch3 == nil {
			break
		}
	}
}


// DemonstrateSelectWithContext shows select with context
func DemonstrateSelectWithContext() {
	fmt.Println("\n=== Select with Context ===")
	
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	// Create a channel
	ch := make(chan int)
	
	// Start a goroutine to send a value
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch <- 42
	}()
	
	// Select with context
	select {
	case x := <-ch:
		fmt.Printf("Received: %d\n", x)
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	}
}
