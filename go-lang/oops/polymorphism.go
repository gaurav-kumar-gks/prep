package oops

/*
1. INTERFACE-BASED POLYMORPHISM
==============================
Go achieves polymorphism through interfaces:

Example:
```go
type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    animals := []Animal{
        Dog{Name: "Rex"},
        Cat{Name: "Whiskers"},
    }

    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

2. RUNTIME TYPE ASSERTION
========================
Type assertions allow runtime type checking:

Example:
```go
func ProcessAnimal(a Animal) {
    switch v := a.(type) {
    case Dog:
        fmt.Printf("Processing dog: %s\n", v.Name)
    case Cat:
        fmt.Printf("Processing cat: %s\n", v.Name)
    default:
        fmt.Println("Unknown animal type")
    }
}

func main() {
    animals := []Animal{
        Dog{Name: "Rex"},
        Cat{Name: "Whiskers"},
    }

    for _, animal := range animals {
        ProcessAnimal(animal)
    }
}
```

3. INTERFACE COMPOSITION
=======================
Interfaces can be composed to create more specific behaviors:

Example:
```go
type Speaker interface {
    Speak() string
}

type Mover interface {
    Move() string
}

type Animal interface {
    Speaker
    Mover
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

func (d Dog) Move() string {
    return "Running"
}

func main() {
    dog := Dog{Name: "Rex"}
    var animal Animal = dog

    fmt.Println(animal.Speak()) // "Woof!"
    fmt.Println(animal.Move())  // "Running"
}
```

4. IMPLICIT INTERFACE IMPLEMENTATION
==================================
Types implicitly implement interfaces:

Example:
```go
type Stringer interface {
    String() string
}

type Person struct {
    Name string
}

// Person implicitly implements Stringer
func (p Person) String() string {
    return p.Name
}

func main() {
    p := Person{Name: "John"}
    var s Stringer = p // Valid because Person has String() method
    fmt.Println(s.String())
}
```

5. POLYMORPHIC FUNCTIONS
=======================
Functions can accept interface types:

Example:
```go
type Processor interface {
    Process() string
}

type DataProcessor struct {
    Data string
}

func (dp DataProcessor) Process() string {
    return "Processing " + dp.Data
}

type FileProcessor struct {
    Filename string
}

func (fp FileProcessor) Process() string {
    return "Processing file " + fp.Filename
}

func ProcessItem(p Processor) string {
    return p.Process()
}

func main() {
    dp := DataProcessor{Data: "test data"}
    fp := FileProcessor{Filename: "test.txt"}

    fmt.Println(ProcessItem(dp)) // "Processing test data"
    fmt.Println(ProcessItem(fp)) // "Processing file test.txt"
}
```

6. POLYMORPHIC COLLECTIONS
=========================
Collections can store different types through interfaces:

Example:
```go
type Storage interface {
    Store() string
}

type Database struct {
    Name string
}

func (db Database) Store() string {
    return "Storing in database " + db.Name
}

type Cache struct {
    Size int
}

func (c Cache) Store() string {
    return fmt.Sprintf("Storing in cache of size %d", c.Size)
}

func main() {
    storages := []Storage{
        Database{Name: "MySQL"},
        Cache{Size: 1024},
    }

    for _, storage := range storages {
        fmt.Println(storage.Store())
    }
}
```

7. POLYMORPHISM BEST PRACTICES
=============================
Guidelines for effective polymorphism:

1. Keep interfaces small and focused
2. Design interfaces for behavior, not data
3. Accept interfaces, return structs
4. Use interface composition for complex behaviors
5. Document interface requirements
6. Consider interface segregation
7. Use type assertions carefully



9. POLYMORPHISM PERFORMANCE
==========================
Performance considerations for polymorphism:

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
Common interview questions about polymorphism:

1. Q: How does Go implement polymorphism?
   A: Go implements polymorphism through interfaces. Types implicitly implement interfaces by providing the required methods.

2. Q: What is the difference between polymorphism in Go and other languages?
   A: Go uses interface-based polymorphism rather than inheritance-based polymorphism. It's more explicit and type-safe.

3. Q: How do you achieve runtime polymorphism in Go?
   A: Runtime polymorphism is achieved through interfaces and type assertions, allowing different types to be treated uniformly.

4. Q: What are the benefits of Go's approach to polymorphism?
   A: Benefits include simplicity, type safety, explicit interfaces, and the ability to implement interfaces without modifying existing types.

5. Q: What are the drawbacks of Go's approach to polymorphism?
   A: Drawbacks include slight performance overhead for interface calls and the need for type assertions in some cases.

6. Q: How do you handle multiple interfaces in Go?
   A: Multiple interfaces can be handled through interface composition, where one interface embeds others.

7. Q: What are the best practices for using polymorphism in Go?
   A: Keep interfaces small, design for behavior, accept interfaces but return structs, and use interface composition for complex behaviors.
*/