package oops

/*
Struct Embedding in Go

This file demonstrates struct embedding and its various use cases in Go.
These examples are for educational purposes and may contain code that would trigger
linter errors if actually run.

1. BASIC EMBEDDING
=================
Basic struct embedding allows one struct to inherit fields and methods from another:

Example:
```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person      // Embedded Person
    Salary float64
}

func main() {
    emp := Employee{
        Person: Person{
            Name: "John",
            Age:  30,
        },
        Salary: 50000,
    }

    // Access embedded fields directly
    fmt.Println(emp.Name)  // "John"
    fmt.Println(emp.Age)   // 30
    fmt.Println(emp.Salary) // 50000
}
```

2. METHOD PROMOTION
==================
Methods of embedded types are promoted to the containing type:

Example:
```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return "Some sound"
}

type Dog struct {
    Animal
    Breed string
}

func main() {
    dog := Dog{
        Animal: Animal{Name: "Rex"},
        Breed:  "German Shepherd",
    }

    // Method promotion
    fmt.Println(dog.Speak()) // "Some sound"
}
```

3. METHOD OVERRIDING
===================
Methods can be overridden in the containing type:

Example:
```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return "Some sound"
}

type Dog struct {
    Animal
    Breed string
}

// Override Speak method
func (d Dog) Speak() string {
    return "Woof!"
}

func main() {
    dog := Dog{
        Animal: Animal{Name: "Rex"},
        Breed:  "German Shepherd",
    }

    fmt.Println(dog.Speak()) // "Woof!"
}
```

4. MULTIPLE EMBEDDING
====================
A struct can embed multiple types:

Example:
```go
type Logger struct {
    Level string
}

func (l Logger) Log(msg string) {
    fmt.Printf("[%s] %s\n", l.Level, msg)
}

type Metrics struct {
    Count int
}

func (m Metrics) Increment() {
    m.Count++
}

type Service struct {
    Logger
    Metrics
    Name string
}

func main() {
    svc := Service{
        Logger:  Logger{Level: "INFO"},
        Metrics: Metrics{Count: 0},
        Name:    "MyService",
    }

    svc.Log("Service started")
    svc.Increment()
}
```

5. EMBEDDING INTERFACES
======================
Structs can embed interfaces:

Example:
```go
type Writer interface {
    Write([]byte) error
}

type FileWriter struct {
    Writer
    FilePath string
}

type ConsoleWriter struct{}

func (c ConsoleWriter) Write(data []byte) error {
    fmt.Println(string(data))
    return nil
}

func main() {
    fw := FileWriter{
        Writer:   ConsoleWriter{},
        FilePath: "output.txt",
    }

    fw.Write([]byte("Hello, World!"))
}
```

6. EMBEDDING POINTERS
====================
Structs can embed pointers to other types:

Example:
```go
type Config struct {
    Host string
    Port int
}

type Server struct {
    *Config  // Embedded pointer
    Name     string
}

func main() {
    config := &Config{
        Host: "localhost",
        Port: 8080,
    }

    server := Server{
        Config: config,
        Name:   "MyServer",
    }

    fmt.Println(server.Host) // "localhost"
    fmt.Println(server.Port) // 8080
}
```

7. EMBEDDING BEST PRACTICES
==========================
Guidelines for effective embedding:

1. Use embedding for "is-a" relationships
2. Keep the embedding hierarchy shallow
3. Be aware of method promotion
4. Consider interface embedding for flexibility
5. Document embedded types
6. Avoid ambiguous field access
7. Use embedding for code reuse

Example:
```go
// Good: Clear "is-a" relationship
type Animal struct {
    Name string
}

type Dog struct {
    Animal  // Dog is an Animal
    Breed string
}

// Bad: Unclear relationship
type Logger struct {
    Level string
}

type Service struct {
    Logger  // Service is not a Logger
    Name string
}

// Better: Use composition
type Service struct {
    logger Logger
    Name   string
}
```

8. EMBEDDING PATTERNS
====================
Common embedding patterns in Go:

Example:
```go
// 1. Decorator Pattern
type Logger struct {
    Level string
}

func (l Logger) Log(msg string) {
    fmt.Printf("[%s] %s\n", l.Level, msg)
}

type Service struct {
    Logger
    Name string
}

// 2. Mixin Pattern
type Validator struct {
    Errors []string
}

func (v *Validator) AddError(err string) {
    v.Errors = append(v.Errors, err)
}

type User struct {
    Validator
    Name string
}

// 3. Trait Pattern
type Printable struct {
    Format string
}

func (p Printable) Print() {
    fmt.Printf(p.Format, p)
}

type Document struct {
    Printable
    Content string
}
```

9. EMBEDDING PERFORMANCE
========================
Performance considerations for embedding:

1. Embedded fields are accessed directly
2. Method promotion has no runtime cost
3. Memory layout is optimized
4. Interface embedding has minimal overhead
5. Pointer embedding can save memory

Example:
```go
// Memory layout is contiguous
type Animal struct {
    Name string
    Age  int
}

type Dog struct {
    Animal
    Breed string
}

// Memory layout:
// [Name][Age][Breed]
// All fields are directly accessible
```

10. INTERVIEW TIPS
=================
Common interview questions about embedding:

1. Q: What is struct embedding in Go?
   A: Struct embedding allows one struct to inherit fields and methods from another struct, promoting code reuse and composition.

2. Q: How does method promotion work with embedding?
   A: Methods of embedded types are automatically promoted to the containing type, allowing them to be called directly on the containing type.

3. Q: Can you override promoted methods?
   A: Yes, you can override promoted methods by defining a method with the same name and signature in the containing type.

4. Q: What are the benefits of embedding?
   A: Benefits include code reuse, method promotion, and the ability to create complex types from simpler ones.

5. Q: What are the drawbacks of embedding?
   A: Drawbacks include potential method name collisions, ambiguous field access, and the risk of creating overly complex types.

6. Q: When should you use embedding vs. composition?
   A: Use embedding for "is-a" relationships and when you want method promotion. Use composition (fields) for "has-a" relationships.

7. Q: Can you embed interfaces?
   A: Yes, you can embed interfaces in structs, which is useful for implementing the interface while adding additional fields.
*/