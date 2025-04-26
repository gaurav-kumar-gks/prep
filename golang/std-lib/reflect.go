package stdlib

/*
Reflect Package

This file demonstrates the usage of Go's reflect package with detailed examples.
These examples are for educational purposes and may contain code that would trigger
linter errors if actually run.

1. BASIC REFLECTION
=================
Basic reflection operations:

Example:
```go
import (
    "fmt"
    "reflect"
)

// Person represents a person with name and age
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create a person
    person := Person{
        Name: "John Doe",
        Age:  30,
    }
    
    // Get type information
    t := reflect.TypeOf(person)
    fmt.Println("Type:", t.Name()) // Type: Person
    
    // Get value information
    v := reflect.ValueOf(person)
    fmt.Println("Kind:", v.Kind()) // Kind: struct
    
    // Get field information
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        fmt.Printf("Field: %s, Type: %v, Value: %v\n", 
            field.Name, field.Type, value.Interface())
    }
    // Field: Name, Type: string, Value: John Doe
    // Field: Age, Type: int, Value: 30
}
```

2. MODIFYING VALUES
=================
Modifying values using reflection:

Example:
```go
import (
    "fmt"
    "reflect"
)

// Person represents a person with name and age
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create a person
    person := Person{
        Name: "John Doe",
        Age:  30,
    }
    
    // Get a pointer to the person
    p := reflect.ValueOf(&person).Elem()
    
    // Modify the Name field
    nameField := p.FieldByName("Name")
    if nameField.IsValid() && nameField.CanSet() {
        nameField.SetString("Jane Doe")
    }
    
    // Modify the Age field
    ageField := p.FieldByName("Age")
    if ageField.IsValid() && ageField.CanSet() {
        ageField.SetInt(25)
    }
    
    // Print the modified person
    fmt.Printf("Modified person: %+v\n", person)
    // Modified person: {Name:Jane Doe Age:25}
}
```

3. CALLING METHODS
================
Calling methods using reflection:

Example:
```go
import (
    "fmt"
    "reflect"
)

// Person represents a person with name and age
type Person struct {
    Name string
    Age  int
}

// Greet returns a greeting for the person
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, my name is %s", p.Name)
}

// SetAge sets the age of the person
func (p *Person) SetAge(age int) {
    p.Age = age
}

func main() {
    // Create a person
    person := Person{
        Name: "John Doe",
        Age:  30,
    }
    
    // Get type information
    t := reflect.TypeOf(person)
    
    // Call the Greet method
    greetMethod, _ := t.MethodByName("Greet")
    result := greetMethod.Func.Call([]reflect.Value{reflect.ValueOf(person)})
    fmt.Println(result[0].String())
    // Hello, my name is John Doe
    
    // Get a pointer to the person
    p := reflect.ValueOf(&person).Elem()
    
    // Call the SetAge method
    setAgeMethod := p.MethodByName("SetAge")
    setAgeMethod.Call([]reflect.Value{reflect.ValueOf(40)})
    
    // Print the modified person
    fmt.Printf("Modified person: %+v\n", person)
    // Modified person: {Name:John Doe Age:40}
}
```

4. CREATING NEW VALUES
====================
Creating new values using reflection:

Example:
```go
import (
    "fmt"
    "reflect"
)

// Person represents a person with name and age
type Person struct {
    Name string
    Age  int
}

func main() {
    // Get type information
    t := reflect.TypeOf(Person{})
    
    // Create a new Person
    newPerson := reflect.New(t).Elem()
    
    // Set field values
    nameField := newPerson.FieldByName("Name")
    if nameField.IsValid() && nameField.CanSet() {
        nameField.SetString("Alice")
    }
    
    ageField := newPerson.FieldByName("Age")
    if ageField.IsValid() && ageField.CanSet() {
        ageField.SetInt(35)
    }
    
    // Convert back to Person
    person := newPerson.Interface().(Person)
    
    // Print the new person
    fmt.Printf("New person: %+v\n", person)
    // New person: {Name:Alice Age:35}
    
    // Create a slice of Person
    sliceType := reflect.SliceOf(t)
    newSlice := reflect.MakeSlice(sliceType, 3, 3)
    
    // Set values in the slice
    for i := 0; i < newSlice.Len(); i++ {
        elem := newSlice.Index(i)
        elem.FieldByName("Name").SetString(fmt.Sprintf("Person %d", i+1))
        elem.FieldByName("Age").SetInt(int64(20 + i*5))
    }
    
    // Convert back to []Person
    people := newSlice.Interface().([]Person)
    
    // Print the slice
    fmt.Printf("People: %+v\n", people)
    // People: [{Name:Person 1 Age:20} {Name:Person 2 Age:25} {Name:Person 3 Age:30}]
}
```

5. STRUCT TAGS
=============
Working with struct tags using reflection:

Example:
```go
import (
    "fmt"
    "reflect"
)

// Person represents a person with name and age
type Person struct {
    Name string `json:"name" validate:"required"`
    Age  int    `json:"age" validate:"min=0,max=150"`
}

func main() {
    // Create a person
    person := Person{
        Name: "John Doe",
        Age:  30,
    }
    
    // Get type information
    t := reflect.TypeOf(person)
    
    // Get field information
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        
        // Get JSON tag
        jsonTag := field.Tag.Get("json")
        
        // Get validation tag
        validateTag := field.Tag.Get("validate")
        
        fmt.Printf("Field: %s, JSON tag: %s, Validation tag: %s\n", 
            field.Name, jsonTag, validateTag)
    }
    // Field: Name, JSON tag: name, Validation tag: required
    // Field: Age, JSON tag: age, Validation tag: min=0,max=150
}
```

6. TYPE ASSERTIONS AND CONVERSIONS
================================
Type assertions and conversions using reflection:

Example:
```go
import (
    "fmt"
    "reflect"
)

// Animal interface
type Animal interface {
    Speak() string
}

// Dog struct
type Dog struct {
    Name string
}

// Speak method for Dog
func (d Dog) Speak() string {
    return "Woof!"
}

// Cat struct
type Cat struct {
    Name string
}

// Speak method for Cat
func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    // Create a Dog
    dog := Dog{Name: "Rex"}
    
    // Create a Cat
    cat := Cat{Name: "Whiskers"}
    
    // Create a slice of Animal
    animals := []Animal{dog, cat}
    
    // Use reflection to get the concrete type
    for i, animal := range animals {
        t := reflect.TypeOf(animal)
        v := reflect.ValueOf(animal)
        
        fmt.Printf("Animal %d: Type=%v, Value=%v\n", i, t, v)
        
        // Call the Speak method using reflection
        speakMethod := v.MethodByName("Speak")
        result := speakMethod.Call(nil)
        fmt.Printf("Animal %d says: %s\n", i, result[0].String())
    }
    // Animal 0: Type=Dog, Value:{Rex}
    // Animal 0 says: Woof!
    // Animal 1: Type=Cat, Value:{Whiskers}
    // Animal 1 says: Meow!
    
    // Type conversion using reflection
    var animal Animal = dog
    
    // Check if animal is a Dog
    if t := reflect.TypeOf(animal); t == reflect.TypeOf(Dog{}) {
        // Convert to Dog
        dogValue := reflect.ValueOf(animal).Interface().(Dog)
        fmt.Printf("It's a dog named %s\n", dogValue.Name)
    }
    // It's a dog named Rex
}
```

7. BEST PRACTICES
===============
Best practices for using reflection in Go:

Example:
```go
import (
    "fmt"
    "reflect"
)

// Person represents a person with name and age
type Person struct {
    Name string
    Age  int
}

// PrintStruct prints the fields of a struct
func PrintStruct(v interface{}) {
    // Get value information
    val := reflect.ValueOf(v)
    
    // Check if v is a pointer
    if val.Kind() == reflect.Ptr {
        // Dereference the pointer
        val = val.Elem()
    }
    
    // Check if v is a struct
    if val.Kind() != reflect.Struct {
        fmt.Println("Not a struct")
        return
    }
    
    // Get type information
    t := val.Type()
    
    // Print struct name
    fmt.Printf("Struct: %s\n", t.Name())
    
    // Print fields
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := val.Field(i)
        
        fmt.Printf("  %s: %v = %v\n", field.Name, field.Type, value.Interface())
    }
}

func main() {
    // Create a person
    person := Person{
        Name: "John Doe",
        Age:  30,
    }
    
    // Print the person struct
    PrintStruct(person)
    // Struct: Person
    //   Name: string = John Doe
    //   Age: int = 30
    
    // Create a pointer to a person
    personPtr := &Person{
        Name: "Jane Doe",
        Age:  25,
    }
    
    // Print the person pointer
    PrintStruct(personPtr)
    // Struct: Person
    //   Name: string = Jane Doe
    //   Age: int = 25
}
``` 