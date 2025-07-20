package basics

import (
	"errors"
	"fmt"
)

/*
ERROR HANDLING
=============

This file demonstrates Go's error handling, including:
- Error creation and checking
- Custom error types
- Error wrapping
- Error handling patterns
- Panic and recover

INTERNALS AND ADVANCED CONCEPTS:
------------------------------
1. Error Interface:
   - The error interface is defined as: type error interface { Error() string }
   - Any type that implements the Error() string method satisfies the error interface
   - Errors are values, not exceptions
   - Errors are returned as the last return value of functions
   - Errors should be handled by the caller

2. Error Creation:
   - errors.New() creates a new error with a string message
   - fmt.Errorf() creates a formatted error message
   - Custom error types can be created by implementing the error interface
   - Error wrapping (Go 1.13+) allows adding context to errors
   - errors.Unwrap() and errors.Is() provide error inspection

3. Error Handling Patterns:
   - Check errors immediately after function calls
   - Return errors up the call stack
   - Add context to errors when returning them
   - Use errors.Is() to check for specific errors
   - Use errors.As() to extract error details
   - Use defer to ensure cleanup in error cases

4. Panic and Recover:
   - Panic is used for exceptional conditions
   - Recover can only be used inside deferred functions
   - Panic and recover are not for normal error handling
   - Panic and recover are used for truly exceptional situations
   - The runtime prints a stack trace when a panic occurs

5. Best Practices:
   - Always check errors
   - Don't ignore errors
   - Add context to errors
   - Use custom error types for specific error cases
   - Use error wrapping to preserve the error chain
   - Use errors.Is() and errors.As() for error checking
*/

// Basic error creation and checking
func DemonstrateBasicErrors() {
	fmt.Println("\n=== Basic Error Creation and Checking ===")
	
	// Create errors
	err1 := errors.New("this is an error")
	err2 := fmt.Errorf("this is a formatted error: %d", 42)
	
	// Check errors
	if err1 != nil {
		fmt.Printf("Error 1: %v\n", err1)
	}
	
	if err2 != nil {
		fmt.Printf("Error 2: %v\n", err2)
	}
	

}


// ValidationError represents a validation error
type ValidationError struct {
	Field string
	Issue string
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Issue)
}

// NotFoundError represents a not found error
type NotFoundError struct {
	Resource string
	ID       string
}

// Error implements the error interface
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %s not found", e.Resource, e.ID)
}

// Custom error types
func DemonstrateCustomErrors() {
	fmt.Println("\n=== Custom Error Types ===")
	
	// Create custom errors
	validationErr := &ValidationError{
		Field: "email",
		Issue: "invalid format",
	}
	
	notFoundErr := &NotFoundError{
		Resource: "user",
		ID:       "123",
	}
	
	// Check custom errors
	if err := validationErr; err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}
	
	if err := notFoundErr; err != nil {
		fmt.Printf("Not found error: %v\n", err)
	}
	
	// Function that returns a custom error
	user, err := findUser("456")
	if err != nil {
		switch e := err.(type) {
		case *NotFoundError:
			fmt.Printf("User not found: %v\n", e)
		case *ValidationError:
			fmt.Printf("Validation error: %v\n", e)
		default:
			fmt.Printf("Unknown error: %v\n", err)
		}
	} else {
		fmt.Printf("User found: %v\n", user)
	}
}

// Mock function that returns a custom error
func findUser(id string) (string, error) {
	if id == "123" {
		return "John Doe", nil
	}
	return "", &NotFoundError{
		Resource: "user",
		ID:       id,
	}
}

// Error wrapping (Go 1.13+)
func DemonstrateErrorWrapping() {
	baseErr := errors.New("base error")
	wrappedErr := fmt.Errorf("failed to process request: %w", baseErr)
	// check for sentinel errors using errors.Is(err, targetErr)
	if errors.Is(wrappedErr, baseErr) {
		fmt.Println("The error is wrapped")
	}
	
	unwrappedErr := errors.Unwrap(wrappedErr)
	if unwrappedErr != nil {
		fmt.Printf("Unwrapped error: %v\n", unwrappedErr)
	}
	
	// Function that wraps errors
	result, err := processData("invalid data")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		
		// check for custom errors using errors.As(err, &targetErr)
		var validationErr *ValidationError
		if errors.As(err, &validationErr) {
			fmt.Printf("Validation error: %v\n", validationErr)
		}
	} else {
		fmt.Printf("Result: %v\n", result)
	}
}

func processData(data string) (string, error) {
	if data == "invalid data" {
		validationErr := &ValidationError{
			Field: "data",
			Issue: "invalid format",
		}
		return "", fmt.Errorf("failed to process data: %w", validationErr)
	}
	return "processed data", nil
}
