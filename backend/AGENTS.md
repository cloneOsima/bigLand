# Agent Guidelines for bigLand Backend

## Build/Lint/Test Commands

### Testing
- **All tests**: `make test` or `go test ./internal/services/ ./internal/handlers/ -v`
- **Single test**: `go test -run TestFunctionName ./internal/handlers/` or `go test -run TestFunctionName ./internal/services/`
- **Specific package**: `go test ./internal/services/ -v` or `go test ./internal/handlers/ -v`

### Building & Running
- **Build**: `go build ./cmd/main.go`
- **Run**: `make run` or `go run ./cmd/main.go`
- **Generate mocks/SQL**: `make generate` (runs sqlc + mockery)

### Code Generation
- **SQL generation**: `sqlc generate`
- **Mock generation**: `mockery`

## Code Style Guidelines

### Go Version & Dependencies
- Go 1.25.0
- Gin framework for HTTP routing
- PostgreSQL with pgx/v5 driver
- sqlc for type-safe SQL queries
- mockery for interface mocks
- testify for assertions and mocking

### Import Organization
```go
import (
    // Standard library
    "context"
    "fmt"

    // Third-party
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"

    // Internal
    "github.com/cloneOsima/bigLand/backend/internal/services"
)
```

### Naming Conventions
- **Functions, Constants, Variables, Interface**: PascalCase for exported, camelCase for unexported
- **Packages**: lowercase (`handlers`, `services`, `repositories`)
- **Handler Function**: Start with Get, Create, Update, Delete (`GetUserInfo`, `CreatePost`)
- **Service Function**: Define our business logic (`Login`, `SignUp`)
- **Repositories Function**: Start with Select, Insert, Update, Delete (`SelectPosts`, `InsertNewAccount`)

### Architecture Patterns
- **Clean Architecture**: handlers → services → repositories
- **Dependency Injection**: Constructor functions (`NewUserHandler`)
- **Interface-based**: All layers define interfaces
- **Context usage**: Pass `context.Context` for timeouts/cancellation

### Error Handling
- Use `log.Fatalf` for fatal startup errors
- Return errors from functions, handle in handlers
- Context timeout handling (`context.DeadlineExceeded`, `context.Canceled`)
- Custom errors should be defined at ./internal/errors/

### Testing Patterns
- **Framework**: testify with mocks
- **Structure**: Table-driven tests with `t.Run`
- **Parallel**: Use `t.Parallel()` in tests
- **Setup**: `gin.SetMode(gin.TestMode)` for HTTP tests
- **Mocking**: mockery-generated mocks in `internal/mocks/`

### SQL & Database
- **Generator**: sqlc with PostgreSQL
- **Types**: UUIDs use `github.com/google/uuid.UUID`
- **Nullables**: Pointers for nullable fields (`*float64`)
- **Location data**: `[]byte` for PostGIS geometries