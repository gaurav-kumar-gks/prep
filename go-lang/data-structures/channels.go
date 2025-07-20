package data_structures

import (
	"fmt"
	"time"
)

/*
Channels in Go

Basic Usage:
CHANNEL METHODS
==============
1. make(chan type) - Create unbuffered channel
   ch := make(chan int)

2. make(chan type, size) - Create buffered channel
   ch := make(chan int, 10)

3. ch <- value - Send value to channel
   ch <- 42

4. value := <-ch - Receive value from channel
   x := <-ch

5. value, ok := <-ch - Receive with ok check
   x, ok := <-ch

6. close(ch) - Close channel
   close(ch)

7. cap(ch) - Get channel capacity
   capacity := cap(ch)

8. len(ch) - Get number of elements in channel
   length := len(ch)

9. select - Select from multiple channels
   select {
   case x := <-ch1:
   case ch2 <- y:
   case <-time.After(time.Second):
   }

10. <-ch - Receive and discard value
    <-ch


*/

// DemonstrateChannels shows channel operations
func DemonstrateChannels() {
	fmt.Println("\n=== Channels ===")
	
	// 1. Unbuffered channel
	// note: to receive from channe
	// first send to it (usually in a goroutine)
	// otherwise, it will block
	ch1 := make(chan int)
	go func() {
		ch1 <- 42
	}()
	fmt.Printf("Unbuffered channel value: %d\n", <-ch1)
	
	// 2. Buffered channel
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Printf("Buffered channel values: %d, %d\n", <-ch2, <-ch2)
	
	// 3. Channel with ok check
	ch3 := make(chan string)
	go func() {
		ch3 <- "hello"
		close(ch3)
	}()
	
	if value, ok := <-ch3; ok {
		fmt.Printf("Channel value with ok: %s\n", value)
	}
	
	// 4. Channel capacity and length
	ch4 := make(chan int, 3)
	ch4 <- 1
	ch4 <- 2
	fmt.Printf("Channel capacity: %d, length: %d\n", cap(ch4), len(ch4))
	
	// 5. Channel direction
	ch5 := make(chan int)
	// ch := make(chan<- int) // send-only channel
	// ch := make(<-chan int) // receive-only channel
	go func(ch chan<- int) {
		ch <- 42
	}(ch5)
	fmt.Printf("Send-only channel value: %d\n", <-ch5)
	
	// 6. Channel selection
	ch6 := make(chan int)
	ch7 := make(chan string)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch6 <- 42
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch7 <- "hello"
	}()
	
	select {
	case x := <-ch6:
		fmt.Printf("Received from ch6: %d\n", x)
	case s := <-ch7:
		fmt.Printf("Received from ch7: %s\n", s)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Timeout")
	}
	
	// 7. Channel can have
	//  int, slice, bool, map, interface, pointer 
	// struct, error, time, function, nil, empty values

	// close channel
	// after closing, no more values can be sent
	// but can still receive
	ch8 := make(chan int)
	go func() {
		ch8 <- 42
	}()
	close(ch8)
	if v, ok := <-ch8; !ok {
		fmt.Println("Channel closed")
	} else {
		fmt.Println("Channel not closed, value received: ", v)
	}
}

// DemonstrateChannelInternals shows channel internals
func DemonstrateChannelInternals() {
	fmt.Println("\n=== Channel Internals ===")
	// 4. Channel blocking
	ch3 := make(chan int)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch3 <- 42
	}()
	
	select {
	case x := <-ch3:
		fmt.Printf("Received value: %d\n", x)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Channel blocked")
	}
}

// DemonstrateChannelPatterns shows common channel patterns
func DemonstrateChannelPatterns() {
	fmt.Println("\n=== Channel Patterns ===")
	
	// 1. Channel as signal
	done := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		close(done)
	}()
	
	select {
	case <-done:
		fmt.Println("Signal received")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout")
	}
	
	// 2. Channel as semaphore
	// a semaphore is a signaling mechanism 
	sem := make(chan struct{}, 2)
	
	go func() {
		sem <- struct{}{}
		fmt.Println("Worker 1 started")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Worker 1 finished")
		<-sem
	}()
	
	go func() {
		sem <- struct{}{}
		fmt.Println("Worker 2 started")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Worker 2 finished")
		<-sem
	}()
	
	time.Sleep(200 * time.Millisecond)
	
	// 3. Channel as pipeline
	numbers := make(chan int)
	squares := make(chan int)
	
	go func() {
		defer close(numbers)
		for i := 0; i < 5; i++ {
			numbers <- i
		}
	}()
	
	go func() {
		defer close(squares)
		for n := range numbers {
			squares <- n * n
		}
	}()
	
	for square := range squares {
		fmt.Printf("Square: %d\n", square)
	}
	
	// 4. Channel as fan-out
	input := make(chan int)
	output1 := make(chan int)
	output2 := make(chan int)
	
	go func() {
		defer close(input)
		for i := 0; i < 5; i++ {
			input <- i
		}
	}()
	
	go func() {
		defer close(output1)
		for n := range input {
			output1 <- n * 2
		}
	}()
	
	go func() {
		defer close(output2)
		for n := range input {
			output2 <- n * 3
		}
	}()
	
	for i := 0; i < 5; i++ {
		select {
		case x := <-output1:
			fmt.Printf("Output 1: %d\n", x)
		case x := <-output2:
			fmt.Printf("Output 2: %d\n", x)
		}
	}
	
	// 5. Channel as fan-in
	ch1 := make(chan int)
	ch2 := make(chan int)
	merged := make(chan int)
	
	go func() {
		defer close(ch1)
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
	}()
	
	go func() {
		defer close(ch2)
		for i := 5; i < 10; i++ {
			ch2 <- i
		}
	}()
	
	go func() {
		defer close(merged)
		for {
			select {
			case x, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				merged <- x
			case x, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				merged <- x
			}
			if ch1 == nil && ch2 == nil {
				return
			}
		}
	}()
	
	for x := range merged {
		fmt.Printf("Merged: %d\n", x)
	}
	
	// 6. Channel as timeout
	timeout := make(chan bool)
	go func() {
		time.Sleep(100 * time.Millisecond)
		timeout <- true
	}()
	
	select {
	case <-timeout:
		fmt.Println("Timeout occurred")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Operation completed")
	}
	
	// 7. Channel as cancellation
	cancel := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		close(cancel)
	}()
	
	select {
	case <-cancel:
		fmt.Println("Operation cancelled")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Operation completed")
	}
	
	// 8. Channel as rate limiter
	limiter := time.Tick(100 * time.Millisecond)
	for i := 0; i < 5; i++ {
		<-limiter
		fmt.Printf("Rate limited operation %d\n", i+1)
	}
	
	// 9. Channel as worker pool
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	
	// Start workers
	for i := 0; i < 3; i++ {
		go func() {
			for j := range jobs {
				results <- j * 2
			}
		}()
	}
	
	// Send jobs
	go func() {
		defer close(jobs)
		for i := 0; i < 5; i++ {
			jobs <- i
		}
	}()
	
	// Collect results
	for i := 0; i < 5; i++ {
		fmt.Printf("Result: %d\n", <-results)
	}
	
	// 10. Channel as event bus
	type Event struct {
		Type    string
		Payload interface{}
	}
	
	bus := make(chan Event)
	
	go func() {
		bus <- Event{"user.created", "john"}
		bus <- Event{"user.updated", "jane"}
		bus <- Event{"user.deleted", "bob"}
		close(bus)
	}()
	
	for event := range bus {
		fmt.Printf("Event: %s, Payload: %v\n", event.Type, event.Payload)
	}
} 