package tips

/*
Debugging and Troubleshooting

1. DEBUGGING CONCURRENT CODE
===========================
func debugConcurrent() {
    debugCh := make(chan string, 100)
    go func() {
        for msg := range debugCh {
            log.Printf("[DEBUG] %s", msg)
        }
    }()

    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            debugCh <- fmt.Sprintf("Goroutine %d started", id)
            // Do work
            time.Sleep(100 * time.Millisecond)
            debugCh <- fmt.Sprintf("Goroutine %d completed", id)
        }(i)
    }

    wg.Wait()
    close(debugCh)
}

2. RACE DETECTOR
===============
Using the race detector:

// To run with race detector:
go run -race main.go

3. PANIC AND RECOVER
===================
Using panic and recover for debugging:

func processWithRecover() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
            // Log the stack trace
            debug.PrintStack()
        }
    }()
    panic("intentional panic for debugging")
}

3. DEBUGGER USAGE
================
Using a debugger (e.g., Delve):

1) go install github.com/go-delve/delve/cmd/dlv@latest
2) create a launch config for debugger
mkdir -p .vscode
cd .vscode
cat > launch.json <<EOF
{
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}"
        },
        {
            "name": "Launch File",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${file}"
        },
        {
            "name": "Launch test/t.go",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/test/t.go"
        }
    ]
}
EOF
3) set breakpoints & then just run the debugger


4. RUNTIME DEBUGGING
===================
Using runtime package for debugging:

func debugRuntime() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
    fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
    fmt.Printf("\tNumGC = %v\n", m.NumGC)

    fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())

    buf := make([]byte, 1<<16)
    runtime.Stack(buf, true)
    fmt.Printf("%s", buf)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

5. PPROF PROFILING
=================
Using pprof for profiling:

import (
    "net/http"
    _ "net/http/pprof"
)

func startProfiling() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // Your application code here
}

To view profiles:
go tool pprof http://localhost:6060/debug/pprof/heap
go tool pprof http://localhost:6060/debug/pprof/profile
go tool pprof http://localhost:6060/debug/pprof/goroutine

*/