package gotesting

import (
    "math/rand"
    "sync"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/require"
)


// TEST COVERAGE
// ===============
// Go provides tools to measure test coverage.

// Run with: go test -cover
// Run with: go test -coverprofile=coverage.out
// View with: go tool cover -html=coverage.out

// 1. BASIC TESTING
// ===============

func Add(a, b int) int {
    return a + b
}

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

func TestAddUsingTestify(t *testing.T) {
    result := Add(1, 2)
    assert.Equal(t, 3, result, "Add(1, 2) should equal 3")
}



// 2. BENCHMARKING
// ==============
// Go provides benchmarking tools to measure performance.


func somefunction(arr []int) []int {
    time.Sleep(rand.Intn(10) * time.Millisecond) // Simulate some work
}

func BenchmarkSomefunction(b *testing.B) {
    sizes := []int{10, 100, 1000}

    for _, size := range sizes {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            arr := make([]int, size)
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                somefunction(arr)
            }
        })
    }
}


// 3. TESTING INTERFACES
// ====================
// Testing interfaces requires creating mock implementations.


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

// 4. TESTING CONCURRENT CODE
// =========================
// Testing concurrent code requires special handling.

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

// 5. TESTING WITH SUBTESTS
// =======================
// Subtests allow organizing tests hierarchically.

func IsValidEmail(email string) bool {
    // Simple email validation
    return strings.Contains(email, "@") && strings.Contains(email, ".")
}

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
}


// 11. TESTING WITH GOMOCK
// ======================
// GoMock is a mocking framework for Go.


// type Database interface {
//     Get(id string) (string, error)
// }

// type Service struct {
//     db Database
// }

// func NewService(db Database) *Service {
//     return &Service{db: db}
// }

// func (s *Service) GetUser(id string) (string, error) {
//     return s.db.Get(id)
// }

// func TestService(t *testing.T) {
//     ctrl := gomock.NewController(t)
//     defer ctrl.Finish()

//     // create mock db
//     mockDB := mocks.NewMockDatabase(ctrl)
//     mockDB.EXPECT().Get("user1").Return("John", nil)

//     service := NewService(mockDB)
//     name, err := service.GetUser("user1")

//     if err != nil {
//         t.Errorf("GetUser failed: %v", err)
//     }
//     if name != "John" {
//         t.Errorf("GetUser returned %q; want %q", name, "John")
//     }
// }