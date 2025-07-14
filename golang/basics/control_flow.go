package basics

import (
	"fmt"
)

func DemonstrateIfElse() {
	x := 10
	if y := 20; x < y {
		fmt.Println("x is less than y")
	}
	
	if z := 30; z > 20 {
		fmt.Println("z in if")
		// z is available here
	} else if z < 10 {
		fmt.Println("z in elseif")
	} else {
		fmt.Println("z in else, z =", z)
	}
}

func DemonstrateForLoops() {
	for i := 0; i < 3; i++ {
		if i == 1 {
			continue
		}
		fmt.Printf("%d ", i)
		if i == 2 {
			fmt.Println()
		}
	}
	
	i := 0
	for i < 3 {
		fmt.Printf("i = %d\n", i)
		i++
	}
	fmt.Println()
	
	for {
		fmt.Println("infinite loop with break")
		break
	}
	
	// Range variables are reused in each iteration (be careful with pointers)
	fmt.Println("Range-based for loop: ")
	slice := []int{1, 2, 3}
	for index, value := range slice {
		fmt.Printf("slice[%d] = %d\n", index, value)
	}
	
	fmt.Println("\nRange with map:")
	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Printf("m[%s] = %d\n", key, value)
	}
	
	fmt.Println("\nRange with string:")
	for i, r := range "Hello" {
		fmt.Printf("str[%d] = %c\n", i, r)
	}
	
	fmt.Println("\nNested loops with labels:")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
	
	fmt.Println("\nRange with channel:")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 3
	close(ch)
	
	for v := range ch {
		fmt.Printf("Received: %d\n", v)
	}
	
}

func DemonstrateSwitch() {
	fmt.Println("\nSwitch with multiple cases and fallthrough:")
	switch 2 {
	case 1, 2, 3:
		fmt.Println("One, Two, or Three")
		fallthrough
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Other")
	}

	fmt.Println("\nSwitch with no condition:")
	switch {
	case 1 == 1:
		fmt.Println("1 is equal to 1")
	case 2 == 1:
		fmt.Println("2 is equal to 1")
	default:
		fmt.Println("Neither condition is true")
	}
	
	fmt.Println("\nSwitch with expression:")
	switch x, y := 10, 5; {
	case x < y:
		fmt.Printf("x: {%d} is less than y: {%d}", x, y)
	case x > y:
		fmt.Printf("x: {%d} is more than y: {%d}", x, y)
	default:
		fmt.Println("Neither condition is true")
	}
	
	fmt.Println("\nSwitch with initialization:")
	switch x := 10; x {
	case 10:
		fmt.Println("x is 10")
	default:
		fmt.Println("x is not 10")
	}
	
	fmt.Println("\nSwitch with type assertion:")
	var i interface{} = "hello"
	switch v := i.(type) {
	case string:
		fmt.Printf("It's a string: %s\n", v)
	case int:
		fmt.Printf("It's an int: %d\n", v)
	default:
		fmt.Printf("It's something else: %v\n", v)
	}
}


/*
Defer Statements:
   - Defer pushes a function call onto a stack
   - Defer is commonly used for cleanup operations
   - Defer can be used to modify return values
   - Defer statements are executed even if a panic occurs

Panic and Recover:
   - Panic is used for exceptional conditions
   - Recover can only be used inside deferred functions
   - Recover returns nil if no panic is in progress
   - Panic and recover are not for normal error handling
   - Panic and recover are used for truly exceptional situations
   - The runtime prints a stack trace when a panic occurs
*/
func DemonstrateDefer() {
	// defers executed in lifo
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")
	
	defer func() {
		fmt.Println("This deferred function will be executed second last")
	}()
	
	msg := "Hello"
	defer fmt.Println("Deferred:", msg) // will print "Hello" as msg is evaluated immediately
	msg = "World" // this will not affect the deferred function
	fmt.Println("Current:", msg)
	
	fmt.Println("\nDefer with return values:")
	x := func(t int) (result int) {
		defer func(x int) {
			result = x*5
		}(t)
		return t*2 // This will be overridden by the deferred function
	}(10)
	fmt.Println("Result changed by defer ", x)
	
	fmt.Println("\nDefer with panic:")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	
	panic("Something went wrong!")
}

// Panic and recover
func DemonstratePanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Outer recover: %v\n", r)
		}
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("inner panic recover: %v\n", r)
				panic("inner panicked after recovery")
			} else {
				fmt.Println("No panic to recover from")
			}
		}()
		
		panic("inner panic")
	}()
	
	panic("outer panic") // this won't be reached as 
	// inner panic will be recovered and re-panic 
	// and outer panic will be caught
} 