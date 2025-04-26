package testing

/*
Advanced Testing Techniques

This file demonstrates advanced testing techniques in Go.
These examples are for educational purposes and may contain code that would trigger
linter errors if actually run.

1. PROPERTY-BASED TESTING
========================
Property-based testing tests properties that should hold for all inputs.

Example:
```go
// Using quick package for property-based testing
import "testing/quick"

func TestAddCommutative(t *testing.T) {
    // Test that addition is commutative: a + b = b + a
    f := func(a, b int) bool {
        return Add(a, b) == Add(b, a)
    }

    if err := quick.Check(f, nil); err != nil {
        t.Error(err)
    }
}

func TestAddIdentity(t *testing.T) {
    // Test that 0 is the identity element for addition: a + 0 = a
    f := func(a int) bool {
        return Add(a, 0) == a
    }

    if err := quick.Check(f, nil); err != nil {
        t.Error(err)
    }
}
```

2. FUZZ TESTING
==============
Fuzz testing automatically generates inputs to find edge cases.

Example:
```go
// Using Go 1.18+ fuzz testing
func FuzzReverse(f *testing.F) {
    // Add seed corpus
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)
    }

    f.Fuzz(func(t *testing.T, s string) {
        rev := Reverse(s)
        doubleRev := Reverse(rev)

        if s != doubleRev {
            t.Errorf("Before: %q, after: %q", s, doubleRev)
        }

        if utf8.ValidString(s) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

3. INTEGRATION TESTING
=====================
Integration tests verify that components work together correctly.

Example:
```go
// integration_test.go
package mypackage

import (
    "context"
    "testing"
    "time"
)

func TestIntegration(t *testing.T) {
    // Skip in short mode
    if testing.Short() {
        t.Skip("Skipping integration test in short mode")
    }

    // Setup test environment
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Initialize components
    db := setupTestDB(t)
    defer db.Close()

    service := NewService(db)

    // Test the integration
    result, err := service.Process(ctx, "test-data")
    if err != nil {
        t.Fatalf("Process failed: %v", err)
    }

    // Verify the result
    if result.Status != "success" {
        t.Errorf("Expected status 'success', got '%s'", result.Status)
    }
}
```

4. PERFORMANCE TESTING
=====================
Performance tests verify that code meets performance requirements.

Example:
```go
func TestPerformance(t *testing.T) {
    // Skip in short mode
    if testing.Short() {
        t.Skip("Skipping performance test in short mode")
    }

    // Setup
    data := generateLargeDataSet()

    // Measure performance
    start := time.Now()
    result := ProcessData(data)
    duration := time.Since(start)

    // Verify performance requirements
    if duration > 100*time.Millisecond {
        t.Errorf("ProcessData took %v, want < 100ms", duration)
    }

    // Verify correctness
    if !isValidResult(result) {
        t.Error("ProcessData produced invalid result")
    }
}
```

5. TEST HOOKS
============
Test hooks allow setup and teardown for tests.

Example:
```go
func TestMain(m *testing.M) {
    // Setup before all tests
    setup()

    // Run tests
    code := m.Run()

    // Teardown after all tests
    teardown()

    // Exit with the same code as the tests
    os.Exit(code)
}

func setup() {
    // Initialize test environment
    // Create temporary files, start services, etc.
}

func teardown() {
    // Clean up test environment
    // Remove temporary files, stop services, etc.
}
```

6. TEST HELPERS
==============
Test helpers reduce code duplication in tests.

Example:
```go
// assertEqual is a test helper
func assertEqual(t *testing.T, expected, actual interface{}, msg string) {
    t.Helper() // Marks this function as a test helper
    if expected != actual {
        t.Errorf("%s: expected %v, got %v", msg, expected, actual)
    }
}

func TestWithHelper(t *testing.T) {
    result := Add(1, 2)
    assertEqual(t, 3, result, "Add(1, 2)")
}
```

7. TEST COVERAGE TOOLS
=====================
Advanced coverage tools provide detailed coverage information.

Example:
```go
// Run with: go test -coverprofile=coverage.out
// View with: go tool cover -html=coverage.out -o coverage.html
// View with: go tool cover -func=coverage.out

// Generate coverage report with specific packages
// go test -coverprofile=coverage.out ./...
// go tool cover -html=coverage.out -o coverage.html
```

8. TESTING HTTP HANDLERS
=======================
Testing HTTP handlers requires a test server.

Example:
```go
func TestHTTPHandler(t *testing.T) {
    // Create a test server
    ts := httptest.NewServer(http.HandlerFunc(MyHandler))
    defer ts.Close()

    // Make a request to the test server
    resp, err := http.Get(ts.URL + "/path")
    if err != nil {
        t.Fatalf("Failed to make request: %v", err)
    }
    defer resp.Body.Close()

    // Check status code
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
    }

    // Check response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

    expected := `{"status":"success"}`
    if string(body) != expected {
        t.Errorf("Expected body %q, got %q", expected, string(body))
    }
}
```

9. TESTING DATABASE OPERATIONS
============================
Testing database operations requires a test database.

Example:
```go
func TestDatabaseOperations(t *testing.T) {
    // Create a test database
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("Failed to open test database: %v", err)
    }
    defer db.Close()

    // Create tables
    _, err = db.Exec(`
        CREATE TABLE users (
            id INTEGER PRIMARY KEY,
            name TEXT
        )
    `)
    if err != nil {
        t.Fatalf("Failed to create table: %v", err)
    }

    // Test database operations
    repo := NewUserRepository(db)

    // Test insert
    user := &User{Name: "John"}
    err = repo.Create(user)
    if err != nil {
        t.Errorf("Create failed: %v", err)
    }

    // Test select
    found, err := repo.GetByID(user.ID)
    if err != nil {
        t.Errorf("GetByID failed: %v", err)
    }
    if found.Name != user.Name {
        t.Errorf("Expected name %q, got %q", user.Name, found.Name)
    }
}
```

10. TESTING CONCURRENT CODE WITH RACE DETECTOR
===========================================
The race detector helps find race conditions in concurrent code.

Example:
```go
// Run with: go test -race

func TestConcurrentAccess(t *testing.T) {
    var counter int
    var wg sync.WaitGroup

    // Start multiple goroutines
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter++ // Race condition!
        }()
    }

    wg.Wait()

    // The race detector will detect the race condition
    // and report it as a test failure
}
```

11. TESTING WITH MOCK SERVERS
===========================
Mock servers simulate external services for testing.

Example:
```go
func TestWithMockServer(t *testing.T) {
    // Create a mock server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Simulate external service
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"success"}`))
    }))
    defer server.Close()

    // Create a client that uses the mock server
    client := NewClient(server.URL)

    // Test the client
    result, err := client.GetData()
    if err != nil {
        t.Errorf("GetData failed: %v", err)
    }
    if result.Status != "success" {
        t.Errorf("Expected status 'success', got '%s'", result.Status)
    }
}
```

12. TESTING WITH TESTCONTAINERS
=============================
Testcontainers provide real services in containers for testing.

Example:
```go
// Using testcontainers-go
import (
    "context"
    "testing"
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/wait"
)

func TestWithPostgres(t *testing.T) {
    ctx := context.Background()

    // Start PostgreSQL container
    postgres, err := testcontainers.NewGenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: testcontainers.ContainerRequest{
            Image:        "postgres:13",
            ExposedPorts: []string{"5432/tcp"},
            Env: map[string]string{
                "POSTGRES_USER":     "test",
                "POSTGRES_PASSWORD": "test",
                "POSTGRES_DB":       "test",
            },
            WaitingFor: wait.ForLog("database system is ready to accept connections"),
        },
        Started: true,
    })
    if err != nil {
        t.Fatalf("Failed to start container: %v", err)
    }
    defer postgres.Terminate(ctx)

    // Get container host and port
    host, err := postgres.Host(ctx)
    if err != nil {
        t.Fatalf("Failed to get host: %v", err)
    }

    port, err := postgres.MappedPort(ctx, "5432")
    if err != nil {
        t.Fatalf("Failed to get port: %v", err)
    }

    // Connect to the database
    connStr := fmt.Sprintf("postgres://test:test@%s:%s/test?sslmode=disable", host, port.Port())
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        t.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // Test database operations
    // ...
}
```

13. TESTING WITH GOMOCK ADVANCED FEATURES
======================================
Advanced features of GoMock for complex mocking scenarios.

Example:
```go
func TestAdvancedMocking(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockDatabase(ctrl)

    // Expect multiple calls
    mockDB.EXPECT().Get("user1").Return("John", nil).Times(2)
    mockDB.EXPECT().Get("user2").Return("Jane", nil).Times(1)

    // Expect calls in order
    mockDB.EXPECT().Put("key1", "value1").Return(nil).After(
        mockDB.EXPECT().Get("key1").Return("", fmt.Errorf("not found")),
    )

    // Expect calls with matchers
    mockDB.EXPECT().Get(gomock.Any()).Return("default", nil).AnyTimes()

    // Use the mock
    service := NewService(mockDB)
    // ...
}
```

14. TESTING WITH TESTIFY ADVANCED FEATURES
=======================================
Advanced features of Testify for complex assertions.

Example:
```go
func TestAdvancedAssertions(t *testing.T) {
    // Assert that a function panics
    assert.Panics(t, func() {
        panic("expected panic")
    })

    // Assert that a function panics with a specific value
    assert.PanicsWithValue(t, "expected panic", func() {
        panic("expected panic")
    })

    // Assert that a function eventually succeeds
    assert.Eventually(t, func() bool {
        return someCondition()
    }, 5*time.Second, 100*time.Millisecond)

    // Assert that a function never succeeds
    assert.Never(t, func() bool {
        return someCondition()
    }, 5*time.Second, 100*time.Millisecond)

    // Assert that a function returns an error
    assert.Error(t, someFunction())

    // Assert that a function returns a specific error
    assert.ErrorIs(t, someFunction(), io.EOF)

    // Assert that a function returns an error that wraps another error
    assert.ErrorAs(t, someFunction(), &os.PathError{})
}
```

15. BEST PRACTICES
=================
Advanced testing best practices:

1. Use property-based testing for complex logic
2. Use fuzz testing to find edge cases
3. Write integration tests for critical paths
4. Use performance tests for performance-critical code
5. Use test hooks for setup and teardown
6. Use test helpers to reduce code duplication
7. Use coverage tools to ensure adequate test coverage
8. Use mock servers for external dependencies
9. Use testcontainers for real service dependencies
10. Use advanced mocking features for complex scenarios

Example:
```go
// Comprehensive test suite
func TestComprehensive(t *testing.T) {
    // Skip in short mode
    if testing.Short() {
        t.Skip("Skipping comprehensive test in short mode")
    }

    // Setup
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Start test containers
    postgres, redis := setupTestContainers(t, ctx)
    defer postgres.Terminate(ctx)
    defer redis.Terminate(ctx)

    // Create mock server
    mockServer := setupMockServer(t)
    defer mockServer.Close()

    // Initialize components
    db := connectToPostgres(t, postgres)
    cache := connectToRedis(t, redis)
    client := NewClient(mockServer.URL)

    service := NewService(db, cache, client)

    // Run property-based tests
    t.Run("Property", func(t *testing.T) {
        TestAddCommutative(t)
        TestAddIdentity(t)
    })

    // Run fuzz tests
    t.Run("Fuzz", func(t *testing.T) {
        FuzzReverse(t)
    })

    // Run integration tests
    t.Run("Integration", func(t *testing.T) {
        TestIntegration(t)
    })

    // Run performance tests
    t.Run("Performance", func(t *testing.T) {
        TestPerformance(t)
    })
}
```
*/