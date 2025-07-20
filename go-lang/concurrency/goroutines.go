package concurrency

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Goroutines
// Goroutines are lightweight threads managed by the Go runtime
// Goroutines are created using the go keyword followed by a function call
// Goroutines are used for concurrent programming in Go (functions that run concurrently with other functions)
// Goroutines are not OS threads, they are scheduled and managed by the Go runtime
// Goroutines are multiplexed onto OS threads
// Goroutines are not preemptive (meaning: they don't yield control to other goroutines)
// Goroutines are cooperative (meaning: they yield control to other goroutines when they are blocked)

func DemonstrateGoroutines() {
	fmt.Println("\n=== Goroutines ===")
	
	// Basic goroutine
	// enclosing function won't wait for goroutine to finish, 
	// so we need to use wait group or channel
	// to wait for goroutine to finish
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d\n", id)
		}(i)
	}
	
	// Goroutine with wait group
	// WaitGroup: used to wait for a collection of goroutines to finish
	// Add(1): increments the WaitGroup counter
	// Done(): decrements the WaitGroup counter
	// Wait(): blocks until the WaitGroup counter is zero
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Wait group goroutine %d\n", id)
		}(i)
	}
	wg.Wait()
	
	// Goroutine with mutex
	// Mutex: used to provide mutual exclusion
	// Lock(): locks the mutex
	// Unlock(): unlocks the mutex
	// Mutex is used to protect shared data from concurrent access
	var mu sync.Mutex
	counter := 0
	for i := 0; i < 3; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	
	// Goroutine with RWMutex
	// RWMutex: used to provide read-write mutual exclusion
	// RLock(): locks the mutex for reading
	// RUnlock(): unlocks the mutex for reading
	// RWMutex is used to allow multiple readers or a single writer
	// RWMutex is more efficient than Mutex when there are many readers and few writers
	var rwmu sync.RWMutex
	for i := 0; i < 3; i++ {
		go func() {
			rwmu.RLock()
			fmt.Printf("Read counter: %d\n", counter)
			rwmu.RUnlock()
		}()
	}
	
	// Goroutine with once
	// Once: used to ensure that a function is only called once
	// Do(): calls the function only once
	// Once is used to initialize shared data
	// even if multiple goroutines call it simultaneously it'll 
	var once sync.Once
	for i := 0; i < 3; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("Once executed")
			})
		}()
	}
	
	// Goroutine with cond
	// Cond: used to synchronize goroutines
	// L: the mutex that protects the condition variable
	// Wait(): waits for the condition to be met
	// Signal(): wakes up one goroutine waiting on the condition
	// Broadcast(): wakes up all goroutines waiting on the condition
	// Cond is used to signal between goroutines
	// Cond is used to implement condition variables
	// Cond is used to implement producer-consumer pattern
	// Cond is used to implement notify-wait pattern
	var cond sync.Cond
	cond.L = &sync.Mutex{}
	
	go func() {
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()
		fmt.Println("Condition met")
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		cond.L.Lock()
		cond.Signal()
		cond.L.Unlock()
	}()
	
	// Goroutine with pool
	pool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}
	
	for i := 0; i < 3; i++ {
		go func() {
			buf := pool.Get().([]byte)
			defer pool.Put(buf)
			fmt.Printf("Pool buffer size: %d\n", len(buf))
		}()
	}
	
	// Goroutine with atomic operations
	var atomicCounter int64
	for i := 0; i < 3; i++ {
		go func() {
			atomic.AddInt64(&atomicCounter, 1)
		}()
	}
	
	// Goroutine with context
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled")
		case <-time.After(100 * time.Millisecond):
			fmt.Println("Context completed")
		}
	}()
}

/*

GOROUTINE METHODS
===============
1. go func() {} - Create goroutine
   go func() {
       // code
   }()

2. runtime.Gosched() - Yield processor
   runtime.Gosched()

3. runtime.Goexit() - Exit goroutine
   runtime.Goexit()

4. runtime.NumGoroutine() - Get number of goroutines
   n := runtime.NumGoroutine()

5. runtime.GOMAXPROCS() - Set number of processors
   runtime.GOMAXPROCS(4)

6. runtime.GC() - Force garbage collection
   runtime.GC()

7. runtime.SetFinalizer() - Set finalizer
   runtime.SetFinalizer(obj, func(obj interface{}) {})

8. runtime.KeepAlive() - Keep object alive
   runtime.KeepAlive(obj)

9. runtime.Stack() - Get goroutine stack
   runtime.Stack(buf, false)

10. runtime.Caller() - Get caller information
    pc, file, line, ok := runtime.Caller(0)

11. runtime.Callers() - Get call stack
    pcs := make([]uintptr, 10)
    n := runtime.Callers(0, pcs)

12. runtime.FuncForPC() - Get function for PC
    fn := runtime.FuncForPC(pc)

13. runtime.GOROOT() - Get GOROOT
    root := runtime.GOROOT()

14. runtime.GOOS() - Get GOOS
    os := runtime.GOOS()

15. runtime.GOARCH() - Get GOARCH
    arch := runtime.GOARCH()

16. runtime.Compiler() - Get compiler
    compiler := runtime.Compiler()

17. runtime.Version() - Get Go version
    version := runtime.Version()

18. runtime.MemProfile() - Get memory profile
    var p runtime.MemProfile
    runtime.MemProfile(&p)

19. runtime.BlockProfile() - Get block profile
    var p runtime.BlockProfile
    runtime.BlockProfile(&p)

20. runtime.MutexProfile() - Get mutex profile
    var p runtime.MutexProfile
    runtime.MutexProfile(&p)
*/
