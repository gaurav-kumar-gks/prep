package basics

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*

PRINT FORMATS
============
%v  - Default format for any value
%+v - Default format with field names for structs
%#v - Go syntax representation
%T  - Type of the value
%%  - Literal percent sign

For specific types:
%d, %b, %o, %x / %X, %f  - decimal, binary, octal, hexa, float
%e, %E  - Scientific notation (lowercase)
%s  - String
%p  - Pointer address
%c  - Character (rune)
%U  - Unicode format

DEFAULT VALUES
=============
- int: 0
- float: 0.0
- string: ""
- bool: false
- pointer: nil
- interface: nil
- slice: nil
- map: nil
- channel: nil
- function: nil
- struct: zero value for each field
- array: zero value for each element

var zeroSlize []string // zeroSlice == nil {
var zeroInt int // zeroInt == 0 {
var zeroMap map[string]int // zeroMap == nil
var zeroChan chan int // zeroChan == nil
var zeroFunc func() //zeroFunc == nil
var zeroInterface interface{} // zeroInterface == nil
var zeroStruct struct{} // zeroStruct == (struct{}{}) {
zeroArray [3]int // zeroArray == [3]int{0, 0, 0} {
zeroPointer *int // zeroPointer == nil {

INTERNALS AND ADVANCED CONCEPTS:
------------------------------
1. Memory Layout:
   - Go uses a static type system with compile-time type checking
   - Types have a fixed size in memory (except for slices, maps, channels, interfaces)
   - Alignment requirements affect struct field ordering and padding
   - Zero values are automatically assigned to uninitialized variables
   - Go's type system is nominal (based on declared names) rather than structural
   - No implicit type conversion between numeric types (must be explicit)
   - Interface satisfaction is checked at compile time
   - Type assertions and switches use runtime type information

2. Memory Management:
   - Small values are passed by value (copied)
   - Large values use pointers internally
   - Slices, maps, channels are reference types (contain pointers)
   - Garbage collection handles memory management

*/

var (
	// Package-level variables
	var1 = 100
	var2 = "package level"
)

// Basic types and their sizes
func DemonstrateTypes() {
	
	const (
		pi     = 3.14159
		e      = 2.71828
		answer int = 42
	)
	
	const (
		Sunday = iota // 0
		Monday // 1
	)
	const (
		a = iota * 2 // 0
		b // 2
		c // 4
	)

	var (
		i8  int8    = 127
		i16 int16   = 32767
		i32 int32   = 2147483647
		i64 int64   = 9223372036854775807
		i   int     = 9223372036854775807
		// f32 float32 = 3.14
		// f64 float64 = 3.141592653589793
		// c   complex64 = 1 + 2i
		// b   bool    = true
		// e string = "world"
		// f bool   = false
		// j, k = 1, 2
	)

	// block scope
	{
		x := 1
		fmt.Printf("Block scope x: %v\n", x)
	}
	// fmt.Println(x) // Error: x undefined
	
	// Shadowing
	x := 1
	{
		x := 2
		fmt.Printf("Shadowed x: %v\n", x)
	}
	fmt.Printf("Original x: %v\n", x)
	
	// Package-level variables
	fmt.Printf("Package-level var1: %v\n", var1)
	fmt.Printf("Package-level var2: %v\n", var2)
	
	fmt.Printf("int8: %d bytes, range: %d to %d\n", unsafe.Sizeof(i8), -128, 127)
	fmt.Printf("int16: %d bytes, range: %d to %d\n", unsafe.Sizeof(i16), -32768, 32767)
	fmt.Printf("int32: %d bytes, range: %d to %d\n", unsafe.Sizeof(i32), -2147483648, 2147483647)
	fmt.Printf("int64: %d bytes, range: %d to %d\n", unsafe.Sizeof(i64), -9223372036854775808, 9223372036854775807)
	fmt.Printf("uint8: %d bytes, range: %d to %d\n", unsafe.Sizeof(uint8(0)), 0, 255)
	fmt.Printf("int: %d bytes (platform dependent)\n", unsafe.Sizeof(i))
	
	// String and rune
	str := "Hello, 世界"
	r := '世'
	fmt.Printf("String: %s, length: %d, runes: %d\n", str, len(str), len([]rune(str)))
	fmt.Printf("Rune: %c, value: %d\n", r, r)
	
	// Boolean
	bo := true
	fmt.Printf("Boolean: %v, size: %d bytes\n", bo, unsafe.Sizeof(bo))
	
	// Demonstrate memory layout and alignment
	type ExampleStruct struct {
		a int8    // 1 byte
		b int32   // 4 bytes
		c int8    // 1 byte
		d int64   // 8 bytes
	}
	
	es := ExampleStruct{}
	fmt.Printf("Struct size: %d bytes (includes padding)\n", unsafe.Sizeof(es))
	fmt.Printf("Field offsets: a=%d, b=%d, c=%d, d=%d\n",
		unsafe.Offsetof(es.a),
		unsafe.Offsetof(es.b),
		unsafe.Offsetof(es.c),
		unsafe.Offsetof(es.d))
	
}

func DemonstrateTypeConversion() {
	// 1. Numeric conversions must be explicit
	// 2. No implicit conversion between numeric types
	// 3. No conversion between pointer types
	// 4. No conversion between interface types
	// 5. No conversion between struct types
	// 6. No conversion between array types of different sizes
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(i)
	fmt.Printf("int: %d -> float64: %f -> uint: %d\n", i, f, u)
	
	str := "Hello"
	bytes := []byte(str)
	runes := []rune(str)
	fmt.Printf("String: %s\nBytes: %v\nRunes: %v\n", str, bytes, runes)
	

	// num to str and vice-versa
	// fmt.Sprintf is not the best practice for converting string to int
	// strconv.Atoi and strconv.Itoa are preferred
	// val, err := strconv.Atoi("42") // Convert string to int
	// s := strconv.Itoa(42) // Convert int to string

	numStr := fmt.Sprintf("%d", i)
	fmt.Printf("Number to string: %s\n", numStr)
	
	// Demonstrate unsafe.Pointer conversions (advanced)
	// WARNING: This is unsafe and should be used with extreme caution
	var x int = 42
	ptr := unsafe.Pointer(&x)
	// Convert to uintptr for pointer arithmetic
	addr := uintptr(ptr)
	fmt.Printf("Memory address: %x\n", addr)	
}

func DemonstrateTypeAliases() {
	fmt.Println("\n=== Type Aliases and Custom Types ===")
	
	// Custom types can have methods, but type aliases cannot
	// This is a key difference between type aliases and custom types
	
	// Type alias (just another name for the same type)
	type Celsius = float64
	type Fahrenheit = float64
	i := 2
	x := Celsius(i) // Type alias allows direct conversion
	var c Celsius = 100
	var f Fahrenheit = 212
	fmt.Println("Type alias: ", x)

	// Custom type (new type with same underlying type)
	type Kelvin float64
	var k Kelvin = 373.15
	
	// Type alias allows direct conversion
	var temp float64 = float64(c)
	// Custom type requires explicit conversion
	var tempK float64 = float64(k)
	
	fmt.Printf("Celsius: %v°C\nFahrenheit: %v°F\nKelvin: %vK\n", 
		c, f, k)
	fmt.Printf("As float64: %v, %v\n", temp, tempK)
	
	// Type checking
	fmt.Printf("Type of Celsius: %v\n", reflect.TypeOf(c))
	fmt.Printf("Type of Fahrenheit: %v\n", reflect.TypeOf(f))
	fmt.Printf("Type of Kelvin: %v\n", reflect.TypeOf(k))
	
	// Demonstrate type embedding
	type Person struct {
		Name string
		Age  int
	}
	
	type Employee struct {
		Person      // Embedded type
		EmployeeID  int
		Department  string
	}
	
	emp := Employee{
		Person: Person{Name: "John", Age: 30},
		EmployeeID: 12345,
		Department: "Engineering",
	}
	
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Employee Name: %s\n", emp.Name) // Access embedded field
}

func DemonstrateTypeAssertions() {
	var i interface{} = "hello"
	
	s, ok := i.(string)
	if ok {
		fmt.Printf("Successfully asserted to string: %s\n", s)
	} else {
		fmt.Println("Failed to assert to string")
	}
	
	switch v := i.(type) {
	case string:
		fmt.Printf("It's a string: %s\n", v)
	case int, int32, int64:
		fmt.Printf("It's an int: %d\n", v)
	case bool:
		fmt.Printf("It's a bool: %v\n", v)
	case nil:
		fmt.Println("It's nil")
	case []int, []string:
		fmt.Printf("It's a slice: %v\n", v)
	case map[string]int:
		fmt.Printf("It's a map: %v\n", v)
	default:
		fmt.Printf("It's something else: %v\n", v)
	}
	
	// Demonstrate interface type assertions
	// Interfaces are implemented implicitly
	// type Writer interface {
	// 	Write([]byte) (int, error)
	// }

	// type StringWriter struct{}
	// func (sw StringWriter) Write(data []byte) (int, error) {
	// 	fmt.Println(string(data))
	// 	return len(data), nil
	// }

	// var w Writer = StringWriter{}
	
	// if sw, ok := w.(StringWriter); ok {
	// 	fmt.Println("Successfully asserted to StringWriter")
	// 	sw.Write([]byte("Hello from StringWriter"))
	// }
	// if sw, ok := w.(Writer); ok {
	// 	fmt.Println("Successfully asserted to Writer")
	// 	sw.Write([]byte("Hello from Writer"))
	// } else {
	// 	fmt.Println("Failed to assert to Writer")
	// }

	// ww := &StringWriter{}
	// // type assertions can only be used with interface types
	// if sw, ok := interface{}(ww).(*StringWriter); ok {
	// 	fmt.Println("Successfully asserted to *StringWriter")
	// 	sw.Write([]byte("Hello from *StringWriter"))
	// } else {
	// 	fmt.Println("Failed to assert to *StringWriter")
	// }

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
	var rw interface{} = struct{}{} // Some value that might implement ReadWriter
	switch rw.(type) {
	case Reader:
		fmt.Println("Implements Reader")
	case Writer:
		fmt.Println("Implements Writer")
	case ReadWriter:
		fmt.Println("Implements both Reader and Writer")
	default:
		fmt.Println("Does not implement any of these interfaces")
	}
}
