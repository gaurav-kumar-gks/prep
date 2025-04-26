package oops

/*
Advanced Interface Concepts in Go

This file demonstrates advanced interface concepts and patterns in Go.
These examples are for educational purposes and may contain code that would trigger
linter errors if actually run.

1. INTERFACE COMPOSITION
=======================
Interfaces can be composed of other interfaces:

Example:
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

// Implementation
type File struct {
    // ... file fields
}

func (f *File) Read(p []byte) (n int, err error) {
    // Implementation
    return len(p), nil
}

func (f *File) Write(p []byte) (n int, err error) {
    // Implementation
    return len(p), nil
}

// Usage
func main() {
    var file ReadWriter = &File{}
    // file can be used as both Reader and Writer
}
```

2. INTERFACE SATISFACTION
========================
Types implicitly satisfy interfaces:

Example:
```go
type Stringer interface {
    String() string
}

type Person struct {
    Name string
}

// Person implicitly satisfies Stringer
func (p Person) String() string {
    return p.Name
}

// Usage
func main() {
    p := Person{Name: "John"}
    var s Stringer = p // Valid because Person has String() method
    fmt.Println(s.String())
}
```

3. EMPTY INTERFACE
=================
The empty interface can hold values of any type:

Example:
```go
func PrintAny(v interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", v, v)
}

func main() {
    PrintAny(42)           // int
    PrintAny("hello")      // string
    PrintAny([]int{1, 2})  // []int
    PrintAny(struct{}{})   // struct{}
}
```

4. TYPE ASSERTIONS
=================
Type assertions extract the concrete value from an interface:

Example:
```go
func main() {
    var i interface{} = "hello"

    // Type assertion
    s := i.(string)
    fmt.Println(s) // "hello"

    // Type assertion with ok check
    if s, ok := i.(string); ok {
        fmt.Println(s) // "hello"
    }

    // Type assertion to wrong type
    if n, ok := i.(int); !ok {
        fmt.Println("Not an int")
    }
}
```

5. TYPE SWITCHES
===============
Type switches handle multiple types:

Example:
```go
func PrintType(v interface{}) {
    switch x := v.(type) {
    case int:
        fmt.Printf("Integer: %d\n", x)
    case string:
        fmt.Printf("String: %s\n", x)
    case []int:
        fmt.Printf("Slice: %v\n", x)
    default:
        fmt.Printf("Unknown type: %T\n", x)
    }
}

func main() {
    PrintType(42)
    PrintType("hello")
    PrintType([]int{1, 2, 3})
    PrintType(struct{}{})
}
```

6. INTERFACE VALUES
==================
Interface values contain both type and value:

Example:
```go
type Animal interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    var a Animal

    // Interface value is nil
    fmt.Println(a == nil) // true

    // Interface value with nil concrete value
    var d *Dog
    a = d
    fmt.Println(a == nil) // false

    // Interface value with non-nil concrete value
    a = &Dog{}
    fmt.Println(a == nil) // false
}
```

7. INTERFACE BEST PRACTICES
==========================
Guidelines for effective interface design:

1. Keep interfaces small and focused
2. Accept interfaces, return structs
3. Interface names should be er-nouns
4. Prefer composition over large interfaces
5. Design interfaces for behavior, not data

Example:
```go
// Good: Small, focused interface
type Writer interface {
    Write([]byte) error
}

// Bad: Large, unfocused interface
type BigInterface interface {
    Read() error
    Write() error
    Delete() error
    Update() error
    // ... many more methods
}

// Good: Accept interfaces
func ProcessData(w Writer, data []byte) error {
    return w.Write(data)
}

// Bad: Accept concrete types
func ProcessData(w *File, data []byte) error {
    return w.Write(data)
}
```

8. INTERFACE PATTERNS
====================
Common interface patterns in Go:

Example:
```go
// 1. Interface Segregation
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

// 2. Dependency Injection
type Logger interface {
    Log(msg string)
}

type Service struct {
    logger Logger
}

func NewService(logger Logger) *Service {
    return &Service{logger: logger}
}

// 3. Strategy Pattern
type Strategy interface {
    Execute() string
}

type Context struct {
    strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
    c.strategy = s
}

func (c *Context) ExecuteStrategy() string {
    return c.strategy.Execute()
}
```

9. INTERFACE PERFORMANCE
========================
Performance considerations for interfaces:

1. Interface method calls have slight overhead
2. Type assertions can be expensive
3. Interface values use more memory
4. Interface composition has minimal cost
5. Empty interfaces are less efficient

Example:
```go
// More efficient: Direct method call
type Dog struct{}
func (d Dog) Bark() string { return "Woof!" }

// Less efficient: Interface method call
type Animal interface {
    Bark() string
}

func main() {
    // Direct call
    d := Dog{}
    _ = d.Bark()

    // Interface call
    var a Animal = d
    _ = a.Bark()
}
```

10. INTERVIEW TIPS
=================
Common interview questions about interfaces:

1. Q: What is an interface in Go?
   A: An interface is a type that defines a set of methods. Any type that implements these methods implicitly satisfies the interface.

2. Q: How does Go handle interface implementation?
   A: Go uses implicit interface implementation. A type satisfies an interface by implementing all its methods, without explicitly declaring it.

3. Q: What is the empty interface?
   A: The empty interface (interface{}) can hold values of any type, making it useful for generic programming in Go.

4. Q: What are type assertions?
   A: Type assertions extract the concrete value from an interface. They can be used with the "ok" idiom to safely check types.

5. Q: What are type switches?
   A: Type switches allow handling multiple types in a switch statement, using the special syntax switch x := v.(type).

6. Q: What is interface composition?
   A: Interface composition allows creating new interfaces by combining existing ones, promoting code reuse and flexibility.

7. Q: What are the best practices for interface design?
   A: Keep interfaces small, accept interfaces but return structs, use interface names as er-nouns, and design for behavior rather than data.
*/