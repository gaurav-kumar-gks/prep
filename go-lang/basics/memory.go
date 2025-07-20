package basics

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"unsafe"
)

/*
Memory Management in Go

Basic Usage:
- Memory allocation
- Garbage collection
- Memory statistics
- Memory profiling
- Memory leaks

Advanced Usage:
- Memory pools
- Object pooling
- Memory alignment
- Memory barriers
- Memory ordering

Internals:
- Memory allocator
- Garbage collector
- Memory model
- Memory barriers
- Memory ordering

Best Practices:
- Memory allocation
- Memory reuse
- Memory pooling
- Memory alignment
- Memory barriers

Common Pitfalls:
- Memory leaks
- Memory fragmentation
- Memory pressure
- Memory ordering
- Memory barriers

Performance Considerations:
- Memory allocation
- Memory reuse
- Memory pooling
- Memory alignment
- Memory barriers

Interview Tips:
- Memory allocation
- Garbage collection
- Memory model
- Memory barriers
- Memory ordering
*/

func DemonstrateMemoryLayout() {
	fmt.Println("\n=== Memory Layout ===")
	
	// 1. Struct alignment
	type AlignedStruct struct {
		a int32    // 4 bytes
		b int64    // 8 bytes
		c int32    // 4 bytes
		d int64    // 8 bytes
	}
	
	type UnalignedStruct struct {
		a int32    // 4 bytes
		c int32    // 4 bytes
		b int64    // 8 bytes
		d int64    // 8 bytes
	}
	
	aligned := AlignedStruct{}
	unaligned := UnalignedStruct{}
	
	fmt.Printf("Aligned struct size: %v bytes\n", unsafe.Sizeof(aligned)) // 32 bytes
	fmt.Printf("Unaligned struct size: %v bytes\n", unsafe.Sizeof(unaligned)) // 24 bytes
}

// DemonstrateMemoryAllocation shows memory allocation patterns
func DemonstrateMemoryAllocation() {
	fmt.Println("\n=== Memory Allocation ===")
	
	// Stack vs Heap allocation
	// Small values are allocated on the stack
	x := 42
	fmt.Printf("Stack allocation: x=%v, size=%v bytes\n", 
		x, unsafe.Sizeof(x))
	
	// Large values are allocated on the heap
	largeSlice := make([]int, 1000000)
	fmt.Printf("Heap allocation: largeSlice size=%v bytes\n", 
		unsafe.Sizeof(largeSlice))

	// Memory pooling
	pool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}
	
	// Get from pool
	buf := pool.Get().([]byte)
	fmt.Printf("Pool allocation: buf size=%v bytes\n", len(buf))
	
	// Put back to pool
	pool.Put(buf)
}

// DemonstrateGarbageCollection shows garbage collection patterns
func DemonstrateGarbageCollection() {
	fmt.Println("\n=== Garbage Collection ===")
	
	// 1. GC statistics
	var stats debug.GCStats
	debug.ReadGCStats(&stats)
	fmt.Printf("GC stats: %+v\n", stats)
	
	// 2. GC control
	// Set GC percent (default is 100)
	debug.SetGCPercent(100)
	
	// 3. GC trigger
	debug.FreeOSMemory()
	
	// 4. Memory statistics
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Memory stats: %+v\n", memStats)
	
	// 5. GC cycle
	fmt.Println("Starting GC cycle...")
	runtime.GC()
	fmt.Println("GC cycle completed")
}

// DemonstrateMemoryModel shows memory model patterns
func DemonstrateMemoryModel() {
	fmt.Println("\n=== Memory Model ===")
	
	// 1. Memory ordering
	var x, y int
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	// Goroutine 1
	go func() {
		defer wg.Done()
		x = 1
	}()
	
	// Goroutine 2
	go func() {
		defer wg.Done()
		y = 1
	}()
	
	wg.Wait()
	fmt.Printf("Memory ordering: x=%v, y=%v\n", x, y)
	
	// 2. Memory barriers
	var a, b int
	var mu sync.Mutex
	
	mu.Lock()
	a = 1
	b = 2
	mu.Unlock()
	
	fmt.Printf("Memory barriers: a=%v, b=%v\n", a, b)
	
	// 3. Memory model guarantees
	// - No data races
	// - No out-of-thin-air values
	// - No reordering of operations
	fmt.Println("Memory model guarantees:")
	fmt.Println("- No data races")
	fmt.Println("- No out-of-thin-air values")
	fmt.Println("- No reordering of operations")
}

// DemonstrateMemoryPools shows memory pool patterns
func DemonstrateMemoryPools() {
	fmt.Println("\n=== Memory Pools ===")
	
	// 1. Object pooling
	type Object struct {
		Data []byte
	}
	
	pool := sync.Pool{
		New: func() interface{} {
			return &Object{
				Data: make([]byte, 1024),
			}
		},
	}
	
	// Get from pool
	obj := pool.Get().(*Object)
	fmt.Printf("Pool object: size=%v bytes\n", len(obj.Data))
	
	// Put back to pool
	pool.Put(obj)
	
	// 2. Memory reuse
	// Reuse slice capacity
	slice := make([]int, 0, 100)
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}
	fmt.Printf("Reused slice: len=%v, cap=%v\n", 
		len(slice), cap(slice))
	
	// 3. Memory pressure
	// Allocate memory to trigger GC
	pressure := make([]byte, 1024*1024*10)
	fmt.Printf("Memory pressure: size=%v bytes\n", len(pressure))
}

// DemonstrateMemoryLeaks shows memory leak patterns
func DemonstrateMemoryLeaks() {
	fmt.Println("\n=== Memory Leaks ===")
	
	// 1. Goroutine leaks
	// Bad: Goroutine leak
	badChan := make(chan int)
	go func() {
		<-badChan // Block forever
	}()
	
	// Good: Proper cleanup
	goodChan := make(chan int)
	done := make(chan struct{})
	go func() {
		select {
		case <-goodChan:
			fmt.Println("Received value")
		case <-done:
			fmt.Println("Cleanup")
		}
	}()
	
	// Cleanup
	close(done)
	
	// 2. Resource leaks
	// Bad: Resource leak
	// badFile, _ := os.Open("file.txt")
	// defer badFile.Close() // Missing cleanup
	
	// Good: Proper cleanup
	goodFile, _ := os.Open("file.txt")
	defer goodFile.Close()
	
	// 3. Memory leak detection
	fmt.Println("Memory leak detection:")
	fmt.Println("- Use pprof")
	fmt.Println("- Monitor memory usage")
	fmt.Println("- Check goroutine count")
	fmt.Println("- Profile heap allocations")
}