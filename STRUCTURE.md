# Sitemapper Project Structure

This document explains the directory structure and architecture of the Sitemapper project.

## Directory Layout

```
sitemapper/
├── cmd/
│   └── server/
│       └── main.go                    # Application entry point
│
├── internal/                          # Private application code
│   ├── config/
│   │   └── config.go                  # Configuration loader
│   │
│   ├── models/                        # Domain models (database-agnostic)
│   │   ├── category.go
│   │   ├── changeset.go
│   │   ├── entry.go
│   │   ├── entry_categories.go
│   │   ├── release.go
│   │   ├── report.go
│   │   ├── report_job.go
│   │   └── user.go
│   │
│   ├── repositories/                  # Repository interfaces
│   │   └── interfaces.go              # All repository contracts
│   │
│   ├── database/                      # Database implementations
│   │   ├── database.go                # Database factory
│   │   │
│   │   ├── postgres/                  # PostgreSQL adapter
│   │   │   ├── postgres.go
│   │   │   ├── entry_repository.go
│   │   │   ├── report_repository.go
│   │   │   ├── user_repository.go
│   │   │   ├── category_repository.go
│   │   │   ├── report_job_repository.go
│   │   │   └── release_repository.go
│   │   │
│   │   ├── mysql/                     # MySQL adapter
│   │   │   ├── mysql.go
│   │   │   └── stubs.go
│   │   │
│   │   ├── sqlite/                    # SQLite adapter
│   │   │   ├── sqlite.go
│   │   │   └── stubs.go
│   │   │
│   │   └── memory/                    # In-memory implementation
│   │       ├── memory.go
│   │       ├── entry_repository.go
│   │       └── stubs.go
│   │
│   ├── services/                      # Business logic layer
│   │   ├── sitemap_service.go
│   │   ├── report_service.go
│   │   ├── job_service.go
│   │   └── category_service.go
│   │
│   └── handlers/                      # HTTP handlers (controllers)
│       ├── sitemap_handler.go
│       ├── report_handler.go
│       ├── user_handler.go
│       └── middleware/
│           ├── auth.go
│           └── logging.go
│
├── pkg/                               # Public, reusable packages
│   ├── sitemap/
│   │   ├── parser.go
│   │   └── validator.go
│   ├── http/
│   │   └── client.go
│   └── scheduler/
│       └── scheduler.go
│
├── configs/                           # Configuration files
│   ├── config.yaml
│   └── config.example.yaml
│
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## Architecture Overview

### Layer Separation

```
Handlers (HTTP) → Services (Business Logic) → Repositories (Data Access) → Database (Adapters)
```

1. **Handlers**: Handle HTTP requests/responses, validation, and routing
2. **Services**: Contain business logic, orchestrate operations across repositories
3. **Repositories**: Abstract data access, provide CRUD operations
4. **Database Adapters**: Implement repository interfaces for specific databases

### Key Patterns

#### Repository Pattern

Repository interfaces are defined in `internal/repositories/interfaces.go`:

```go
type EntryRepository interface {
    Create(ctx context.Context, entry *models.Entry) error
    GetByID(ctx context.Context, id string) (*models.Entry, error)
    // ... other methods
}
```

Each database adapter implements these interfaces:
- `internal/database/postgres/` - PostgreSQL implementation
- `internal/database/mysql/` - MySQL implementation
- `internal/database/sqlite/` - SQLite implementation
- `internal/database/memory/` - In-memory implementation (for testing)

#### Adapter Pattern

The database factory (`internal/database/database.go`) creates the appropriate database adapter based on configuration:

```go
func NewDatabase(cfg *config.Config) (repositories.Database, error) {
    switch cfg.DatabaseType {
    case "postgres":
        return postgres.New(cfg.DatabaseURL)
    case "mysql":
        return mysql.New(cfg.DatabaseURL)
    // ...
    }
}
```

### internal/ vs pkg/

- **internal/**: Private code that cannot be imported by other projects
  - Contains business logic specific to Sitemapper
  - Knows about domain models and business rules
  
- **pkg/**: Public utilities that could be reused by other projects
  - Generic, reusable libraries
  - No knowledge of Sitemapper's business domain

## Database Support

The application supports multiple database types through the adapter pattern:

1. **PostgreSQL** - Production-ready, full-featured
2. **MySQL** - Alternative production database
3. **SQLite** - Lightweight, great for development/testing
4. **Memory** - In-memory storage for unit tests

Users can select their database by configuring `database_type` in `configs/config.yaml`.

## Getting Started

### Using Make

```bash
# Install dependencies
make deps

# Run the application
make run

# Build the application
make build

# Run tests
make test

# Format code
make fmt
```

### Using Docker

```bash
# Run with PostgreSQL
docker-compose up

# Run with MySQL
docker-compose --profile mysql up
```

## Configuration

Configuration is managed through YAML files in the `configs/` directory:

1. Copy `config.example.yaml` to `config.yaml`
2. Update with your database settings
3. Run the application

Alternatively, use environment variables (see `.env.example`).

## Testing

The in-memory database adapter (`internal/database/memory/`) makes testing easy:

```go
func TestMyService(t *testing.T) {
    db := memory.New()
    defer db.Close()
    
    service := services.NewReportService(db)
    // Test your service...
}
```

No need to set up a real database for unit tests!

## Adding a New Database Adapter

1. Create a new directory: `internal/database/yourdb/`
2. Implement the `repositories.Database` interface
3. Implement each repository interface
4. Add the adapter to the database factory
5. Update configuration documentation

## Contributing

When adding new features:

1. Define models in `internal/models/`
2. Add repository interfaces in `internal/repositories/`
3. Implement business logic in `internal/services/`
4. Add HTTP handlers in `internal/handlers/`
5. Update all database adapters to support new repositories

