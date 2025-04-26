package testing

/*
Go Testing and Benchmarking

This file demonstrates Go's testing and benchmarking capabilities with detailed examples.
These examples are for educational purposes and may contain code that would trigger
linter errors if actually run.

1. BASIC TESTING
===============
Go provides a built-in testing package for writing tests.

Example:
```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 1, 2, 3},
        {"negative numbers", -1, -2, -3},
        {"mixed numbers", 1, -2, -1},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestSubtract(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 3, 2, 1},
        {"negative numbers", -1, -2, 1},
        {"mixed numbers", 1, -2, 3},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Subtract(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

2. TABLE-DRIVEN TESTS
====================
Table-driven tests allow testing multiple cases with a single test function.

Example:
```go
// stringutils.go
package stringutils

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// stringutils_test.go
package stringutils

import "testing"

func TestReverse(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"empty string", "", ""},
        {"single character", "a", "a"},
        {"palindrome", "radar", "radar"},
        {"normal string", "hello", "olleh"},
        {"unicode string", "café", "éfac"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Reverse(tt.input)
            if result != tt.expected {
                t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
            }
        })
    }
}
```

3. BENCHMARKING
==============
Go provides benchmarking tools to measure performance.

Example:
```go
// sort.go
package sort

func BubbleSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}

func QuickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[len(arr)/2]
    var left, right []int

    for _, x := range arr {
        if x < pivot {
            left = append(left, x)
        } else if x > pivot {
            right = append(right, x)
        }
    }

    left = QuickSort(left)
    right = QuickSort(right)

    return append(append(left, pivot), right...)
}

// sort_test.go
package sort

import (
    "math/rand"
    "testing"
)

func BenchmarkBubbleSort(b *testing.B) {
    sizes := []int{10, 100, 1000}

    for _, size := range sizes {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            arr := make([]int, size)
            for i := range arr {
                arr[i] = rand.Intn(1000)
            }

            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                BubbleSort(arr)
            }
        })
    }
}

func BenchmarkQuickSort(b *testing.B) {
    sizes := []int{10, 100, 1000}

    for _, size := range sizes {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            arr := make([]int, size)
            for i := range arr {
                arr[i] = rand.Intn(1000)
            }

            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                QuickSort(arr)
            }
        })
    }
}
```

4. TEST COVERAGE
===============
Go provides tools to measure test coverage.

Example:
```go
// calculator.go
package calculator

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

func Multiply(a, b int) int {
    return a * b
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// calculator_test.go
package calculator

import (
    "testing"
)

func TestAdd(t *testing.T) {
    result := Add(1, 2)
    if result != 3 {
        t.Errorf("Add(1, 2) = %d; want 3", result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(3, 2)
    if result != 1 {
        t.Errorf("Subtract(3, 2) = %d; want 1", result)
    }
}

// Run with: go test -cover
// Run with: go test -coverprofile=coverage.out
// View with: go tool cover -html=coverage.out
```

5. TESTING MAIN PACKAGE
======================
Testing the main package requires special handling.

Example:
```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

// main_test.go
package main

import "testing"

func TestMain(t *testing.T) {
    // This is a placeholder test for the main package
    // Main packages are typically not tested directly
    t.Skip("Main package tests are typically not run")
}
```

6. TESTING INTERFACES
====================
Testing interfaces requires creating mock implementations.

Example:
```go
// storage.go
package storage

type Storage interface {
    Get(key string) (string, error)
    Put(key, value string) error
    Delete(key string) error
}

type FileStorage struct {
    // Implementation details
}

func (fs *FileStorage) Get(key string) (string, error) {
    // Implementation
    return "value", nil
}

func (fs *FileStorage) Put(key, value string) error {
    // Implementation
    return nil
}

func (fs *FileStorage) Delete(key string) error {
    // Implementation
    return nil
}

// storage_test.go
package storage

import "testing"

// MockStorage is a mock implementation of the Storage interface
type MockStorage struct {
    data map[string]string
}

func NewMockStorage() *MockStorage {
    return &MockStorage{
        data: make(map[string]string),
    }
}

func (ms *MockStorage) Get(key string) (string, error) {
    if value, ok := ms.data[key]; ok {
        return value, nil
    }
    return "", fmt.Errorf("key not found: %s", key)
}

func (ms *MockStorage) Put(key, value string) error {
    ms.data[key] = value
    return nil
}

func (ms *MockStorage) Delete(key string) error {
    if _, ok := ms.data[key]; !ok {
        return fmt.Errorf("key not found: %s", key)
    }
    delete(ms.data, key)
    return nil
}

func TestStorage(t *testing.T) {
    storage := NewMockStorage()

    // Test Put
    err := storage.Put("key1", "value1")
    if err != nil {
        t.Errorf("Put failed: %v", err)
    }

    // Test Get
    value, err := storage.Get("key1")
    if err != nil {
        t.Errorf("Get failed: %v", err)
    }
    if value != "value1" {
        t.Errorf("Get returned %q; want %q", value, "value1")
    }

    // Test Delete
    err = storage.Delete("key1")
    if err != nil {
        t.Errorf("Delete failed: %v", err)
    }

    // Test Get after Delete
    _, err = storage.Get("key1")
    if err == nil {
        t.Error("Get after Delete succeeded; want error")
    }
}
```

7. TESTING CONCURRENT CODE
=========================
Testing concurrent code requires special handling.

Example:
```go
// counter.go
package counter

type Counter struct {
    value int
    mu    sync.Mutex
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

// counter_test.go
package counter

import (
    "sync"
    "testing"
)

func TestCounter(t *testing.T) {
    c := &Counter{}

    // Test single increment
    c.Increment()
    if c.Value() != 1 {
        t.Errorf("Counter value = %d; want 1", c.Value())
    }

    // Test concurrent increments
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            c.Increment()
        }()
    }

    wg.Wait()
    if c.Value() != 101 {
        t.Errorf("Counter value = %d; want 101", c.Value())
    }
}
```

8. TESTING WITH SUBTESTS
=======================
Subtests allow organizing tests hierarchically.

Example:
```go
// validator.go
package validator

func IsValidEmail(email string) bool {
    // Simple email validation
    return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func IsValidPassword(password string) bool {
    // Simple password validation
    return len(password) >= 8
}

// validator_test.go
package validator

import "testing"

func TestValidator(t *testing.T) {
    t.Run("Email", func(t *testing.T) {
        t.Run("Valid", func(t *testing.T) {
            if !IsValidEmail("user@example.com") {
                t.Error("IsValidEmail('user@example.com') = false; want true")
            }
        })

        t.Run("Invalid", func(t *testing.T) {
            if IsValidEmail("invalid-email") {
                t.Error("IsValidEmail('invalid-email') = true; want false")
            }
        })
    })

    t.Run("Password", func(t *testing.T) {
        t.Run("Valid", func(t *testing.T) {
            if !IsValidPassword("password123") {
                t.Error("IsValidPassword('password123') = false; want true")
            }
        })

        t.Run("Invalid", func(t *testing.T) {
            if IsValidPassword("short") {
                t.Error("IsValidPassword('short') = true; want false")
            }
        })
    })
}
```

9. TESTING WITH EXAMPLES
=======================
Examples serve as both documentation and tests.

Example:
```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

// math_test.go
package math

import "fmt"

func ExampleAdd() {
    result := Add(1, 2)
    fmt.Println(result)
    // Output: 3
}
```

10. TESTING WITH TESTDATA
========================
Test data can be stored in separate files.

Example:
```go
// parser.go
package parser

func ParseJSON(data []byte) (map[string]interface{}, error) {
    var result map[string]interface{}
    err := json.Unmarshal(data, &result)
    return result, err
}

// parser_test.go
package parser

import (
    "os"
    "testing"
)

func TestParseJSON(t *testing.T) {
    // Read test data from file
    data, err := os.ReadFile("testdata/valid.json")
    if err != nil {
        t.Fatalf("Failed to read test data: %v", err)
    }

    result, err := ParseJSON(data)
    if err != nil {
        t.Errorf("ParseJSON failed: %v", err)
    }

    // Check result
    if result["name"] != "John" {
        t.Errorf("result[\"name\"] = %v; want \"John\"", result["name"])
    }
}
```

11. TESTING WITH GOMOCK
======================
GoMock is a mocking framework for Go.

Example:
```go
// service.go
package service

type Database interface {
    Get(id string) (string, error)
}

type Service struct {
    db Database
}

func NewService(db Database) *Service {
    return &Service{db: db}
}

func (s *Service) GetUser(id string) (string, error) {
    return s.db.Get(id)
}

// service_test.go
package service

import (
    "testing"
    "github.com/golang/mock/gomock"
    "github.com/yourusername/yourproject/mocks"
)

func TestService(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockDatabase(ctrl)
    mockDB.EXPECT().Get("user1").Return("John", nil)

    service := NewService(mockDB)
    name, err := service.GetUser("user1")

    if err != nil {
        t.Errorf("GetUser failed: %v", err)
    }
    if name != "John" {
        t.Errorf("GetUser returned %q; want %q", name, "John")
    }
}
```

12. TESTING WITH TESTIFY
=======================
Testify provides additional testing utilities.

Example:
```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

// math_test.go
package math

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
    result := Add(1, 2)
    assert.Equal(t, 3, result, "Add(1, 2) should equal 3")
}

func TestAddWithRequire(t *testing.T) {
    result := Add(1, 2)
    require.Equal(t, 3, result, "Add(1, 2) should equal 3")
    // This line will not be executed if the above assertion fails
    t.Log("This line will be executed if the assertion passes")
}
*/