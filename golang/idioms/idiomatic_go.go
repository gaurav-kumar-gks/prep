package idioms

/*
Idiomatic Go

This file demonstrates common idiomatic patterns and best practices in Go.
These examples are for educational purposes and may contain code that would trigger
linter errors if actually run.

1. ERROR HANDLING
================
Proper error handling patterns:

Example:
```go
// Good: Explicit error handling
func processData(data []byte) error {
    if len(data) == 0 {
        return fmt.Errorf("empty data")
    }
    // Process data
    return nil
}

// Bad: Ignoring errors
func processDataBad(data []byte) {
    _ = json.Unmarshal(data, &result)  // Error ignored
}
```

2. INTERFACE DESIGN
==================
Interface design principles:

Example:
```go
// Good: Small, focused interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Bad: Large, monolithic interfaces
type BadInterface interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
    Close() error
    Flush() error
    Reset() error
    // ... many more methods
}
```

3. COMPOSITION
=============
Using composition over inheritance:

Example:
```go
// Good: Using composition
type Logger struct {
    writer io.Writer
}

func (l *Logger) Log(msg string) {
    fmt.Fprintf(l.writer, "%s\n", msg)
}

// Bad: Using inheritance-like patterns
type BadLogger struct {
    io.Writer  // Embedding for inheritance-like behavior
}
```

4. OPTIONS PATTERN
=================
Using functional options:

Example:
```go
type Server struct {
    port    int
    timeout time.Duration
}

type ServerOption func(*Server)

func WithPort(port int) ServerOption {
    return func(s *Server) {
        s.port = port
    }
}

func WithTimeout(timeout time.Duration) ServerOption {
    return func(s *Server) {
        s.timeout = timeout
    }
}

func NewServer(opts ...ServerOption) *Server {
    s := &Server{
        port:    8080,
        timeout: 30 * time.Second,
    }
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

5. CONTEXT USAGE
===============
Proper context usage:

Example:
```go
// Good: Context propagation
func ProcessRequest(ctx context.Context, data []byte) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        // Process data
        return nil
    }
}

// Bad: Context ignored
func ProcessRequestBad(data []byte) error {
    // Context not used
    return nil
}
```

6. CONCURRENCY PATTERNS
======================
Common concurrency patterns:

Example:
```go
// Good: Using channels for coordination
func worker(jobs <-chan int, results chan<- int) {
    for job := range jobs {
        results <- job * 2
    }
}

// Bad: Using shared memory
var counter int
var mutex sync.Mutex

func incrementBad() {
    mutex.Lock()
    counter++
    mutex.Unlock()
}
```

7. NAMING CONVENTIONS
====================
Go naming conventions:

Example:
```go
// Good: Clear, concise names
type User struct {
    ID        int
    Name      string
    CreatedAt time.Time
}

// Bad: Unclear or inconsistent names
type BadStruct struct {
    user_id        int
    UserName       string
    creation_time  time.Time
}
```

8. DOCUMENTATION
===============
Proper documentation practices:

Example:
```go
// Package example provides examples of idiomatic Go code.
package example

// User represents a system user.
// It contains basic user information and metadata.
type User struct {
    // ID is the unique identifier for the user.
    ID int

    // Name is the user's full name.
    Name string

    // CreatedAt is the timestamp when the user was created.
    CreatedAt time.Time
}

// NewUser creates a new user with the given name.
// It returns an error if the name is empty.
func NewUser(name string) (*User, error) {
    if name == "" {
        return nil, fmt.Errorf("name cannot be empty")
    }
    return &User{
        Name:      name,
        CreatedAt: time.Now(),
    }, nil
}
```

9. TESTING
==========
Testing best practices:

Example:
```go
// Good: Table-driven tests
func TestProcess(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {
            name:    "valid input",
            input:   "hello",
            want:    "HELLO",
            wantErr: false,
        },
        {
            name:    "empty input",
            input:   "",
            want:    "",
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Process(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Process() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

10. BEST PRACTICES
=================
General best practices:

1. Use meaningful variable names
2. Keep functions small and focused
3. Handle errors explicitly
4. Use interfaces for abstraction
5. Prefer composition over inheritance
6. Use context for cancellation
7. Write clear documentation
8. Write comprehensive tests
9. Use proper formatting (gofmt)
10. Follow the standard project layout

Example:
```go
// Good: Clear, well-documented code
type Service struct {
    logger *Logger
    db     *Database
}

// NewService creates a new service instance.
func NewService(logger *Logger, db *Database) *Service {
    return &Service{
        logger: logger,
        db:     db,
    }
}

// Process handles the main business logic.
func (s *Service) Process(ctx context.Context, data []byte) error {
    if err := s.validate(data); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }

    if err := s.db.Save(ctx, data); err != nil {
        return fmt.Errorf("failed to save data: %w", err)
    }

    s.logger.Info("data processed successfully")
    return nil
}
```
*/