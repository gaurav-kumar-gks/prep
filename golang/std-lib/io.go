package stdlib

/*
I/O Package

This file demonstrates the usage of Go's io package with detailed examples.
These examples are for educational purposes and may contain code that would trigger
linter errors if actually run.

1. BASIC I/O OPERATIONS
======================
Common I/O operations using the io package:

Example:
```go
import (
    "io"
    "os"
)

func main() {
    // Reading from a file
    file, err := os.Open("file.txt")
    if err != nil {
        return
    }
    defer file.Close()

    // Read all data
    data, err := io.ReadAll(file)
    if err != nil {
        return
    }

    // Write to a file
    err = os.WriteFile("output.txt", data, 0644)
    if err != nil {
        return
    }

    // Copy data
    src, _ := os.Open("source.txt")
    dst, _ := os.Create("destination.txt")
    defer src.Close()
    defer dst.Close()

    written, err := io.Copy(dst, src)
    if err != nil {
        return
    }
}
```

2. BUFFERED I/O
==============
Using buffered I/O for better performance:

Example:
```go
import (
    "bufio"
    "io"
    "os"
)

func main() {
    // Reading with buffer
    file, _ := os.Open("file.txt")
    defer file.Close()

    reader := bufio.NewReader(file)

    // Read line by line
    for {
        line, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }
        if err != nil {
            return
        }
        // Process line
    }

    // Writing with buffer
    out, _ := os.Create("output.txt")
    defer out.Close()

    writer := bufio.NewWriter(out)
    writer.WriteString("Hello, World!\n")
    writer.Flush()  // Don't forget to flush
}
```

3. PIPE OPERATIONS
================
Using io.Pipe for in-memory data transfer:

Example:
```go
func main() {
    // Create a pipe
    pr, pw := io.Pipe()

    // Write to pipe in a goroutine
    go func() {
        defer pw.Close()
        pw.Write([]byte("Hello, Pipe!"))
    }()

    // Read from pipe
    data, err := io.ReadAll(pr)
    if err != nil {
        return
    }
}
```

4. MULTI-READER/WRITER
=====================
Using MultiReader and MultiWriter:

Example:
```go
func main() {
    // MultiReader
    r1 := strings.NewReader("Hello")
    r2 := strings.NewReader(", ")
    r3 := strings.NewReader("World")

    reader := io.MultiReader(r1, r2, r3)
    data, _ := io.ReadAll(reader)  // "Hello, World"

    // MultiWriter
    var buf1, buf2 bytes.Buffer
    writer := io.MultiWriter(&buf1, &buf2)
    writer.Write([]byte("Hello, World"))
    // Both buffers contain "Hello, World"
}
```

5. TEMPORARY FILES
=================
Working with temporary files:

Example:
```go
func main() {
    // Create temporary file
    tmpFile, err := os.CreateTemp("", "prefix-*.txt")
    if err != nil {
        return
    }
    defer os.Remove(tmpFile.Name())  // Clean up

    // Write to temp file
    tmpFile.Write([]byte("Temporary data"))

    // Create temp directory
    tmpDir, err := os.MkdirTemp("", "prefix-*")
    if err != nil {
        return
    }
    defer os.RemoveAll(tmpDir)  // Clean up
}
```

6. FILE PERMISSIONS
==================
Working with file permissions:

Example:
```go
func main() {
    // Create file with permissions
    err := os.WriteFile("file.txt", []byte("data"), 0644)
    if err != nil {
        return
    }

    // Change permissions
    err = os.Chmod("file.txt", 0755)
    if err != nil {
        return
    }

    // Get file info
    info, err := os.Stat("file.txt")
    if err != nil {
        return
    }
    mode := info.Mode()
}
```

7. DIRECTORY OPERATIONS
======================
Working with directories:

Example:
```go
func main() {
    // Create directory
    err := os.Mkdir("newdir", 0755)
    if err != nil {
        return
    }

    // Create nested directories
    err = os.MkdirAll("parent/child/grandchild", 0755)
    if err != nil {
        return
    }

    // Read directory
    entries, err := os.ReadDir(".")
    if err != nil {
        return
    }

    for _, entry := range entries {
        if entry.IsDir() {
            // Handle directory
        } else {
            // Handle file
        }
    }
}
```

8. FILE SEEKING
==============
Using file seeking operations:

Example:
```go
func main() {
    file, _ := os.Open("file.txt")
    defer file.Close()

    // Seek to position
    offset, err := file.Seek(10, io.SeekStart)
    if err != nil {
        return
    }

    // Read from position
    data := make([]byte, 5)
    n, err := file.Read(data)
    if err != nil {
        return
    }

    // Seek relative to current position
    offset, err = file.Seek(-5, io.SeekCurrent)
    if err != nil {
        return
    }

    // Seek from end
    offset, err = file.Seek(-10, io.SeekEnd)
    if err != nil {
        return
    }
}
```

9. FILE LOCKING
==============
Using file locking:

Example:
```go
func main() {
    file, _ := os.OpenFile("file.txt", os.O_RDWR, 0644)
    defer file.Close()

    // Lock file
    err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
    if err != nil {
        return
    }
    defer syscall.Flock(int(file.Fd()), syscall.LOCK_UN)

    // Perform operations on locked file
    file.Write([]byte("data"))
}
```

10. BEST PRACTICES
=================
Guidelines for effective I/O operations:

1. Always check errors
2. Use defer to close files
3. Use buffered I/O for better performance
4. Clean up temporary files
5. Use appropriate file permissions
6. Handle EOF properly
7. Use io.Copy for large files
8. Consider using io.Pipe for in-memory operations
9. Use MultiReader/MultiWriter when appropriate
10. Handle file locking when needed

Example:
```go
// Good: Proper error handling and cleanup
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return err
    }

    return nil
}

// Bad: Missing error handling and cleanup
func processFileBad(filename string) {
    file, _ := os.Open(filename)
    data, _ := io.ReadAll(file)
    // File never closed, errors ignored
}
```
*/