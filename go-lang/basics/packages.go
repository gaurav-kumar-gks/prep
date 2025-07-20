package basics

import (
	"flag"
	"fmt"
)

/*
Packages in Go


Common Pitfalls:
- Circular dependencies
- Package initialization order
- Package visibility
- Package naming conflicts
- Package version conflicts

Package Structure Examples
=========================

1. Standard Library Structure:
   /net/http
   ├── client.go      # HTTP client implementation
   ├── server.go      # HTTP server implementation
   ├── request.go     # Request handling
   ├── response.go    # Response handling
   ├── transport.go   # Transport layer
   └── internal/      # Internal implementation details
       ├── httputil/  # HTTP utilities
       └── socks/     # SOCKS proxy support

2. Web Application Structure:
   /myapp
   ├── cmd/                    # Application entry points
   │   ├── server/            # Server binary
   │   │   └── main.go
   │   └── cli/               # CLI binary
   │       └── main.go
   ├── internal/              # Private application code
   │   ├── auth/             # Authentication
   │   ├── database/         # Database access
   │   └── service/          # Business logic
   ├── pkg/                   # Public libraries
   │   ├── client/           # Client library
   │   └── utils/            # Utility functions
   ├── api/                   # API definitions
   │   └── proto/            # Protocol buffers
   ├── configs/               # Configuration files
   ├── deployments/           # Deployment configs
   └── test/                  # Test files

3. Microservice Structure:
   /service
   ├── cmd/                   # Service entry point
   │   └── main.go
   ├── internal/              # Private service code
   │   ├── domain/           # Domain models
   │   ├── repository/       # Data access
   │   ├── service/          # Business logic
   │   └── transport/        # API handlers
   ├── pkg/                   # Shared libraries
   │   ├── logger/           # Logging
   │   └── metrics/          # Metrics
   └── api/                   # API definitions
       └── proto/            # gRPC/Protobuf

4. Library Structure:
   /mylib
   ├── mylib.go              # Main package file
   ├── types.go              # Type definitions
   ├── errors.go             # Error definitions
   ├── utils.go              # Utility functions
   └── examples/             # Example usage
       └── example_test.go

5. CLI Tool Structure:
   /tool
   ├── cmd/                  # Command implementations
   │   ├── root.go          # Root command
   │   ├── init.go          # Init command
   │   └── build.go         # Build command
   ├── internal/             # Private implementation
   │   ├── config/          # Configuration
   │   └── runner/          # Command runner
   └── pkg/                  # Public packages
       └── utils/           # Utility functions

Common Package Patterns:
----------------------
1. cmd/ pattern:
   - Entry points for binaries
   - Minimal code, mostly initialization
   - Each subdirectory is a separate binary

2. internal/ pattern:
   - Private application code
   - Cannot be imported by other projects
   - Contains business logic and implementation

3. pkg/ pattern:
   - Public libraries
   - Can be imported by other projects
   - Contains reusable code

4. api/ pattern:
   - API definitions
   - Protocol buffers
   - OpenAPI/Swagger specs

5. test/ pattern:
   - Test files
   - Test utilities
   - Integration tests

Package Naming Conventions:
-------------------------
1. Use lowercase, single-word names
2. Avoid underscores or mixedCaps
3. Use descriptive but concise names
4. Avoid generic names like 'util' or 'common'
5. Use domain-specific names

Package Organization Rules:
-------------------------
1. Keep packages focused and cohesive
2. Avoid circular dependencies
3. Use interfaces for abstraction
4. Keep package APIs stable
5. Document public APIs
6. Use internal/ for private code
7. Use pkg/ for public code
8. Use cmd/ for binaries
*/

// ExportedConstant is visible to other packages
const ExportedConstant = 42

// unexportedConstant is not visible to other packages
const unexportedConstant = 43

// packageInitialized is set during package initialization
var packageInitialized bool

// init function runs during package initialization
func init() {
	packageInitialized = true
	fmt.Println("Package initialized")
	flag.Bool("verbose", false, "Enable verbose output")
	// Set a default value for the flag
	flag.Parse()

}