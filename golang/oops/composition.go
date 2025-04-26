package oops

/*
1. BASIC COMPOSITION
===================
Go favors composition over inheritance. Here's how to compose types:

Example:
```go
// Basic types
type Animal struct {
    Name string
    Age  int
}

type Dog struct {
    Animal      // Embedding Animal
    Breed string
}

type Cat struct {
    Animal      // Embedding Animal
    Color string
}

// Usage
func main() {
    dog := Dog{
        Animal: Animal{
            Name: "Rex",
            Age:  3,
        },
        Breed: "German Shepherd",
    }

    cat := Cat{
        Animal: Animal{
            Name: "Whiskers",
            Age:  2,
        },
        Color: "Orange",
    }

    fmt.Println(dog.Name)  // Accessing embedded field
    fmt.Println(cat.Age)   // Accessing embedded field
}
```

2. METHOD PROMOTION
==================
Methods of embedded types are promoted to the containing type:

Example:
```go
// Methods on Animal
func (a Animal) Speak() string {
    return "Some sound"
}

func (a Animal) GetInfo() string {
    return fmt.Sprintf("%s is %d years old", a.Name, a.Age)
}

// Dog-specific method
func (d Dog) Speak() string {
    return "Woof!"
}

// Usage
func main() {
    dog := Dog{
        Animal: Animal{Name: "Rex", Age: 3},
        Breed:  "German Shepherd",
    }

    fmt.Println(dog.Speak())    // "Woof!" (Dog's method)
    fmt.Println(dog.GetInfo())  // "Rex is 3 years old" (promoted from Animal)
}
```

3. COMPOSITION WITH INTERFACES
=============================
Composition can be used with interfaces:

Example:
```go
type Speaker interface {
    Speak() string
}

type Mover interface {
    Move() string
}

type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return "Some sound"
}

func (a Animal) Move() string {
    return "Moving"
}

type Dog struct {
    Animal
    Breed string
}

// Dog overrides Speak but inherits Move
func (d Dog) Speak() string {
    return "Woof!"
}

// Usage
func main() {
    dog := Dog{
        Animal: Animal{Name: "Rex"},
        Breed:  "German Shepherd",
    }

    var speaker Speaker = dog
    var mover Mover = dog

    fmt.Println(speaker.Speak()) // "Woof!"
    fmt.Println(mover.Move())    // "Moving"
}
```

4. MULTIPLE COMPOSITION
======================
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

// Usage
func main() {
    svc := Service{
        Logger:  Logger{Level: "INFO"},
        Metrics: Metrics{Count: 0},
        Name:    "MyService",
    }

    svc.Log("Service started")  // Uses Logger's method
    svc.Increment()             // Uses Metrics' method
}

BEST PRACTICES
================
Guidelines for effective composition:

1. Keep embedded types simple and focused
2. Use interfaces for flexibility
3. Avoid deep embedding hierarchies
4. Consider method promotion carefully
5. Use composition for code reuse
6. Keep the composition shallow
7. Document the composition relationships


COMMON PITFALLS
==================
Watch out for these composition pitfalls:

1. Method name collisions
2. Ambiguous field access
3. Circular dependencies
4. Over-composition
5. Hidden dependencies


PERFORMANCE CONSIDERATIONS
============================
Composition impact on performance:

1. Method promotion has no runtime cost
2. Embedded fields are accessed directly
3. Interface composition has minimal overhead
4. Memory layout is optimized
5. No virtual method table needed

10. INTERVIEW TIPS
=================
Common interview questions about composition:

1. Q: What is composition in Go?
   A: Composition is Go's primary mechanism for code reuse, where types can embed other types to inherit their fields and methods.

2. Q: How does composition differ from inheritance?
   A: Composition is more explicit, flexible, and maintainable than inheritance. It allows for better code reuse without the drawbacks of inheritance.

3. Q: What is method promotion?
   A: Method promotion allows methods of embedded types to be called directly on the containing type, as if they were defined on it.

4. Q: When should you use composition?
   A: Use composition when you want to reuse code, implement interfaces, or create complex types from simpler ones.

5. Q: What are the benefits of composition?
   A: Benefits include code reuse, flexibility, maintainability, and avoiding the problems of inheritance.
*/