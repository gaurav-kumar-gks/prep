package basics

import "fmt"

/*
FUNCTIONS AND METHODS
====================

This file demonstrates Go's functions and methods, including:
- Function declarations and calls
- Function types and values
- Methods and receivers
- Variadic functions
- Anonymous functions and closures
- Function literals

INTERNALS AND ADVANCED CONCEPTS:
------------------------------
1. Function Declarations:
   - Go functions can have multiple return values
   - Return values can be named
   - Functions are first-class citizens
   - Functions can be passed as arguments
   - Functions can be returned from other functions

2. Method Receivers:
   - Methods can have value or pointer receivers
   - Value receivers make a copy of the value
   - Pointer receivers can modify the value
   - Methods are just functions with a receiver
   - The receiver type must be defined in the same package

3. Function Types:
   - Functions are types in Go
   - Function types can be used as parameters
   - Function types can be used as return values
   - Function types can be used as variables
   - Function types can be used as fields in structs

4. Closures:
   - Closures capture variables from their enclosing scope
   - Closures can modify captured variables
   - Closures are commonly used for callbacks
   - Closures can be used to create function factories
   - Closures can be used to implement iterators

5. Deferred Function Execution:
   - Defer statements execute functions when the surrounding function returns
   - Deferred functions are executed in LIFO order
   - Deferred functions can access and modify return values
   - Deferred functions can recover from panics
   - Deferred functions are commonly used for cleanup
*/

// Basic function declaration
func DemonstrateBasicFunctions() {
	// sum := add(1, 2, 3, 4, 5)
	numbers := []int{1, 2, 3, 4, 5}
	sum := addSlice(numbers...)
	fmt.Println("sum: ", sum)
}

func swap(x, y int) (int, int) {
	return y, x
}

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}

// Variadic function
func add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func addSlice(numbers ...int) int {
	return add(numbers...)
}

// Function types
func DemonstrateFunctionTypes() {
	fmt.Println("\n=== Function Types ===")
	
	// Function type declaration
	type Operation func(int, int) int
	
	// Function as parameter
	applyOperation := func(op Operation, x, y int) int {
		return op(x, y)
	}

	// Function values
	add := func(x, y int) int { return x + y }
	subtract := func(x, y int) int { return x - y }
	multiply := func(x, y int) int { return x * y }

	// Function as return value
	getOperation := func(op string) Operation {
		switch op {
		case "add":
			return add
		case "subtract":
			return subtract
		case "multiply":
			return multiply
		default:
			return add
		}
	}
	
	// Use function as return value
	op := getOperation("add")
	fmt.Printf("getOperation(\"add\")(2, 3) = %d\n", applyOperation(op, 2, 3))
	
	// Function type as struct field
	type Calculator struct {
		operation Operation
	}
	
	calc := Calculator{operation: add}
	fmt.Printf("Calculator{operation: add}(2, 3) = %d\n", calc.operation(2, 3))
}

// Rectangle represents a rectangle with width and height
type Rectangle struct {
	width  int
	height int
}

// Area returns the area of the rectangle
func (r Rectangle) Area() int {
	return r.width * r.height
}

// Scale multiplies the width and height of the rectangle by the given factor
func (r *Rectangle) Scale(factor int) {
	r.width *= factor
	r.height *= factor
}

// difference between value and pointer receiver
// Value receiver: a copy of the receiver is passed, receiver is small and copying is cheap
// Pointer receiver: a pointer to the receiver is passed, receiver is large and copying is expensive

// MyInt is a custom type for demonstrating methods on non-struct types
// note: methods on non-struct types are not common in Go
type MyInt int

// Double returns twice the value of m
func (m MyInt) Double() MyInt {
	return m * 2
}

// Triple multiplies the value of m by 3
func (m *MyInt) Triple() {
	*m *= 3
}

// Methods and receivers
func DemonstrateMethods() {
	
	// Create a rectangle
	rect := Rectangle{width: 5, height: 3}
	fmt.Printf("Rectangle{width: 5, height: 3}.Area() = %d\n", rect.Area())
	
	// Scale the rectangle
	rect.Scale(2) // modifies the rectangle as it is a pointer receiver
	fmt.Printf("After Scale(2): Rectangle{width: %d, height: %d}.Area() = %d\n", rect.width, rect.height, rect.Area())
}

// Closures
func DemonstrateClosures() {
	fmt.Println("\n=== Closures ===")
	
	// Basic closure
	// what is a closure?
	// A closure is a function that captures the variables from its surrounding scope
	// and can access them even after the scope has exited.
	// Closures are commonly used for callbacks and event handling
	// Closures can also be used to create function factories
	// and to implement iterators.
	
	incrbyX, decrbyX := func(i int) (func(x int) int, func(x int) int) {
		val := i
		decr := func(x int) int {
			val -= x
			return val
		}
		incr := func(x int) int {
			val += x
			return val
		}
		return incr, decr
	}(10) // initial value of val is 10
	
	fmt.Printf("incrbyX(2) = %dn", incrbyX(1))
	fmt.Printf("decrbyX(2) = %d \n", decrbyX(2))
	
	// Closure as iterator
	makeRange := func(min, max int) func() (int, bool) {
		next := min
		return func() (int, bool) {
			if next > max {
				return 0, false
			}
			value := next
			next++
			return value, true
		}
	}
	
	iter := makeRange(1, 5)
	// iter is a closure that captures the variables min, max, and next
	for {
		value, ok := iter()
		if !ok {
			break
		}
		fmt.Printf("iter() = %d\n", value)
	}
}

// Anonymous functions
func DemonstrateAnonymousFunctions() {
	// Anonymous function - a function without a name
	// here anonymous function is defined and called immediately -
	// called IIFE (Immediately Invoked Function Expression) 
	func() {
		fmt.Println("This is an anonymous function")
	}()
	
	// Anonymous function with parameters
	add := func(x, y int) int {
		fmt.Printf("This is an anonymous function with parameters: %d, %d\n", x, y)
		return x + y
	}
	
	fmt.Printf("add(5, 3) = %d\n", add(5, 3))
} 