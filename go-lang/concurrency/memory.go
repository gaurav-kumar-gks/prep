package concurrency

/*
Go Memory Management and Garbage Collection

This file demonstrates Go's memory management and garbage collection concepts
with detailed examples. These examples are for educational purposes and may
contain code that would trigger linter errors if actually run.

1. MEMORY ALLOCATION
===================
Go manages memory allocation automatically, but understanding how it works is important.

Example:
```go
func memoryAllocationExample() {
    // Stack vs. Heap allocation
    // Small values are typically allocated on the stack
    x := 42 // Stack allocation

    // Larger values or values that escape the function are allocated on the heap
    y := make([]int, 1000) // Heap allocation

    // Pointer types are typically allocated on the heap
    z := &struct{ value int }{value: 42} // Heap allocation

    // Escape analysis determines if a value escapes to the heap
    // The following would escape to the heap
    escapeExample := func() *int {
        v := 42
        return &v // v escapes to the heap
    }

    // The following would not escape to the heap
    noEscapeExample := func() int {
        v := 42
        return v // v stays on the stack
    }
}
```

2. GARBAGE COLLECTION
====================
Go uses a concurrent, tri-color mark-and-sweep garbage collector.

Example:
```go
func garbageCollectionExample() {
    // Force garbage collection
    runtime.GC()

    // Get garbage collection statistics
    var stats runtime.MemStats
    runtime.ReadMemStats(&stats)

    fmt.Printf("Alloc: %v MiB\n", stats.Alloc / 1024 / 1024)
    fmt.Printf("TotalAlloc: %v MiB\n", stats.TotalAlloc / 1024 / 1024)
    fmt.Printf("Sys: %v MiB\n", stats.Sys / 1024 / 1024)
    fmt.Printf("NumGC: %v\n", stats.NumGC)

    // Create garbage to trigger GC
    for i := 0; i < 1000000; i++ {
        _ = make([]byte, 1024) // Allocate 1KB
    }

    // Force garbage collection again
    runtime.GC()

    // Read stats again
    runtime.ReadMemStats(&stats)
    fmt.Printf("After GC - Alloc: %v MiB\n", stats.Alloc / 1024 / 1024)
    fmt.Printf("After GC - NumGC: %v\n", stats.NumGC)
}
```

3. MEMORY LEAKS
==============
Memory leaks can occur in Go, especially with goroutines and resources.

Example:
```go
func memoryLeakExample() {
    // Goroutine leak
    leakyGoroutine := func() {
        ch := make(chan int)
        go func() {
            // This goroutine is blocked forever
            <-ch
        }()
        // ch is never closed, goroutine leaks
    }

    // Resource leak
    resourceLeak := func() {
        file, err := os.Open("file.txt")
        if err != nil {
            return
        }
        // file is never closed, resource leak
    }

    // Proper resource cleanup
    properCleanup := func() {
        file, err := os.Open("file.txt")
        if err != nil {
            return
        }
        defer file.Close() // Ensure file is closed
    }
}
```

4. MEMORY POOLING
================
Memory pooling can reduce garbage collection pressure.

Example:
```go
func memoryPoolingExample() {
    // Create a pool
    pool := sync.Pool{
        New: func() interface{} {
            return make([]byte, 1024) // 1KB buffer
        },
    }

    // Get a buffer from the pool
    buf := pool.Get().([]byte)

    // Use the buffer
    copy(buf, []byte("Hello, World!"))

    // Return the buffer to the pool
    pool.Put(buf)

    // Get another buffer (might be the same one)
    buf2 := pool.Get().([]byte)

    // Use the buffer
    fmt.Println(string(buf2))

    // Return the buffer to the pool
    pool.Put(buf2)
}
```

5. MEMORY PROFILING
==================
Go provides tools for memory profiling.

Example:
```go
func memoryProfilingExample() {
    // Create a memory profile
    f, err := os.Create("mem.prof")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    // Write memory profile
    pprof.WriteHeapProfile(f)

    // Create a memory profile with runtime/pprof
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    // Do some work
    for i := 0; i < 1000000; i++ {
        _ = make([]byte, 1024)
    }
}
```

6. MEMORY LAYOUT
===============
Understanding memory layout can help optimize performance.

Example:
```go
func memoryLayoutExample() {
    // Struct field ordering affects memory layout
    type BadStruct struct {
        a int8    // 1 byte
        b int64   // 8 bytes
        c int8    // 1 byte
        d int32   // 4 bytes
    }

    type GoodStruct struct {
        b int64   // 8 bytes
        d int32   // 4 bytes
        a int8    // 1 byte
        c int8    // 1 byte
    }

    // BadStruct has padding: 1 + 7(padding) + 8 + 1 + 3(padding) + 4 = 24 bytes
    // GoodStruct has padding: 8 + 4 + 1 + 1 + 2(padding) = 16 bytes

    fmt.Printf("BadStruct size: %d\n", unsafe.Sizeof(BadStruct{}))
    fmt.Printf("GoodStruct size: %d\n", unsafe.Sizeof(GoodStruct{}))
}
```

7. MEMORY BARRIERS
=================
Memory barriers ensure proper ordering of memory operations.

Example:
```go
func memoryBarrierExample() {
    // Atomic operations provide memory barriers
    var x int64
    var y int64

    // This ensures proper ordering
    atomic.StoreInt64(&x, 1)
    atomic.StoreInt64(&y, 2)

    // Without atomic, the compiler might reorder these
    x = 1
    y = 2 // Might be reordered before x = 1
}
```

8. MEMORY MAPPING
================
Memory mapping allows direct access to files.

Example:
```go
func memoryMappingExample() {
    // Open a file
    file, err := os.OpenFile("file.txt", os.O_RDWR, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Get file size
    stat, err := file.Stat()
    if err != nil {
        log.Fatal(err)
    }
    size := stat.Size()

    // Memory map the file
    data, err := syscall.Mmap(int(file.Fd()), 0, int(size), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
    if err != nil {
        log.Fatal(err)
    }
    defer syscall.Munmap(data)

    // Modify the memory-mapped file
    copy(data, []byte("Hello, Memory Mapping!"))

    // Sync changes to disk
    syscall.Msync(data, syscall.MS_SYNC)
}
```

9. MEMORY PRESSURE
=================
Memory pressure can affect application performance.

Example:
```go
func memoryPressureExample() {
    // Create memory pressure
    pressure := func() {
        // Allocate a large amount of memory
        data := make([][]byte, 1000)
        for i := range data {
            data[i] = make([]byte, 1024*1024) // 1MB
        }

        // Hold the memory for a while
        time.Sleep(1 * time.Second)

        // Release the memory
        data = nil
    }

    // Run under memory pressure
    go pressure()

    // Do some work while under memory pressure
    for i := 0; i < 100; i++ {
        _ = make([]byte, 1024)
    }
}
```

10. MEMORY FRAGMENTATION
======================
Memory fragmentation can lead to inefficient memory usage.

Example:
```go
func memoryFragmentationExample() {
    // Create memory fragmentation
    fragments := make([][]byte, 1000)

    // Allocate memory in small chunks
    for i := range fragments {
        fragments[i] = make([]byte, 1024) // 1KB
    }

    // Free some chunks to create fragmentation
    for i := 0; i < 1000; i += 2 {
        fragments[i] = nil
    }

    // Try to allocate a large chunk
    large := make([]byte, 1024*1024) // 1MB

    // The garbage collector might need to compact memory
    runtime.GC()
}
*/