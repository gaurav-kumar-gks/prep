package stdlib

/*
Net Package

This file demonstrates the usage of Go's net package with detailed examples.
These examples are for educational purposes and may contain code that would trigger
linter errors if actually run.

1. TCP SERVER
============
Basic TCP server implementation:

Example:
```go
import (
    "fmt"
    "net"
)

func main() {
    // Listen on TCP port 8080
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        return
    }
    defer listener.Close()

    for {
        // Accept new connections
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

        // Handle connection in a goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Read data
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        return
    }

    // Process data
    data := buffer[:n]

    // Send response
    conn.Write([]byte("Received: " + string(data)))
}
```

2. TCP CLIENT
============
Basic TCP client implementation:

Example:
```go
func main() {
    // Connect to server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        return
    }
    defer conn.Close()

    // Send data
    message := []byte("Hello, Server!")
    _, err = conn.Write(message)
    if err != nil {
        return
    }

    // Read response
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        return
    }

    fmt.Println(string(buffer[:n]))
}
```

3. UDP SERVER
============
UDP server implementation:

Example:
```go
func main() {
    // Create UDP address
    addr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        return
    }

    // Create UDP connection
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        return
    }
    defer conn.Close()

    buffer := make([]byte, 1024)
    for {
        // Read UDP packet
        n, remoteAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            continue
        }

        // Process packet
        data := buffer[:n]

        // Send response
        conn.WriteToUDP([]byte("Received: "+string(data)), remoteAddr)
    }
}
```

4. UDP CLIENT
============
UDP client implementation:

Example:
```go
func main() {
    // Resolve server address
    serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
    if err != nil {
        return
    }

    // Create UDP connection
    conn, err := net.DialUDP("udp", nil, serverAddr)
    if err != nil {
        return
    }
    defer conn.Close()

    // Send data
    message := []byte("Hello, UDP Server!")
    _, err = conn.Write(message)
    if err != nil {
        return
    }

    // Read response
    buffer := make([]byte, 1024)
    n, _, err := conn.ReadFromUDP(buffer)
    if err != nil {
        return
    }

    fmt.Println(string(buffer[:n]))
}
```

5. HTTP SERVER
=============
Basic HTTP server implementation:

Example:
```go
import (
    "fmt"
    "net/http"
)

func main() {
    // Define handler function
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
    })

    // Start server
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        return
    }
}
```

6. HTTP CLIENT
=============
HTTP client implementation:

Example:
```go
func main() {
    // Make GET request
    resp, err := http.Get("http://localhost:8080/World")
    if err != nil {
        return
    }
    defer resp.Body.Close()

    // Read response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return
    }

    fmt.Println(string(body))

    // Make POST request
    data := []byte(`{"message": "Hello"}`)
    resp, err = http.Post(
        "http://localhost:8080/api",
        "application/json",
        bytes.NewBuffer(data),
    )
    if err != nil {
        return
    }
    defer resp.Body.Close()
}
```

7. CUSTOM TRANSPORT
==================
Custom HTTP transport configuration:

Example:
```go
func main() {
    // Create custom transport
    transport := &http.Transport{
        MaxIdleConns:        100,
        IdleConnTimeout:     90 * time.Second,
        DisableCompression:  true,
        TLSHandshakeTimeout: 10 * time.Second,
    }

    // Create client with custom transport
    client := &http.Client{
        Transport: transport,
        Timeout:   30 * time.Second,
    }

    // Use client
    resp, err := client.Get("http://localhost:8080")
    if err != nil {
        return
    }
    defer resp.Body.Close()
}
```

8. DNS RESOLUTION
================
DNS resolution examples:

Example:
```go
func main() {
    // Lookup IP addresses
    ips, err := net.LookupIP("example.com")
    if err != nil {
        return
    }
    for _, ip := range ips {
        fmt.Println(ip)
    }

    // Lookup CNAME
    cname, err := net.LookupCNAME("www.example.com")
    if err != nil {
        return
    }
    fmt.Println(cname)

    // Lookup MX records
    mxRecords, err := net.LookupMX("example.com")
    if err != nil {
        return
    }
    for _, mx := range mxRecords {
        fmt.Printf("Host: %s, Pref: %d\n", mx.Host, mx.Pref)
    }
}
```

9. SOCKET OPTIONS
================
Setting socket options:

Example:
```go
func main() {
    // Create TCP connection
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        return
    }
    defer conn.Close()

    // Get underlying TCP connection
    tcpConn := conn.(*net.TCPConn)

    // Set keep-alive
    tcpConn.SetKeepAlive(true)
    tcpConn.SetKeepAlivePeriod(30 * time.Second)

    // Set read/write deadlines
    tcpConn.SetReadDeadline(time.Now().Add(10 * time.Second))
    tcpConn.SetWriteDeadline(time.Now().Add(10 * time.Second))

    // Set buffer sizes
    tcpConn.SetReadBuffer(8192)
    tcpConn.SetWriteBuffer(8192)
}
```

10. BEST PRACTICES
=================
Guidelines for effective networking:

1. Always close connections
2. Use appropriate timeouts
3. Handle errors properly
4. Use connection pooling when appropriate
5. Implement retry mechanisms
6. Use context for cancellation
7. Set appropriate buffer sizes
8. Handle connection limits
9. Implement proper logging
10. Use TLS for security

Example:
```go
// Good: Proper connection handling with timeouts
func makeRequest(url string) error {
    client := &http.Client{
        Timeout: 30 * time.Second,
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    req = req.WithContext(ctx)
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    return nil
}

// Bad: Missing timeouts and proper cleanup
func makeRequestBad(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    // Missing defer resp.Body.Close()
    return nil
}
```
*/