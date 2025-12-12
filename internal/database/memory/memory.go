package memory

import (
	"context"
	"fmt"
	"sync"

	"jonopens/sitemapper/internal/models"
	"jonopens/sitemapper/internal/repositories"
)

// Database implements repositories.Database with in-memory storage
// Useful for testing and development
type Database struct {
	entries   map[string]*models.Entry
	reports   map[string]*models.Report
	users     map[string]*models.User
	groupings map[string]*models.Group
	jobs      map[string]*models.ReportJob
	releases  map[string]*models.Release
	mu        sync.RWMutex
}

// New creates a new in-memory database
func New() *Database {
	return &Database{
		entries:   make(map[string]*models.Entry),
		reports:   make(map[string]*models.Report),
		users:     make(map[string]*models.User),
		groupings: make(map[string]*models.Group),
		jobs:      make(map[string]*models.ReportJob),
		releases:  make(map[string]*models.Release),
	}
}

// Entries returns the entry repository
func (d *Database) Entries() repositories.EntryRepository {
	return &EntryRepository{db: d}
}

// Reports returns the report repository
func (d *Database) Reports() repositories.ReportRepository {
	return &ReportRepository{db: d}
}

// Users returns the user repository
func (d *Database) Users() repositories.UserRepository {
	return &UserRepository{db: d}
}

// Groupings returns the grouping repository
func (d *Database) Groupings() repositories.GroupingRepository {
	return &GroupingRepository{db: d}
}

// ReportJobs returns the report job repository
func (d *Database) ReportJobs() repositories.ReportJobRepository {
	return &ReportJobRepository{db: d}
}

// Releases returns the release repository
func (d *Database) Releases() repositories.ReleaseRepository {
	return &ReleaseRepository{db: d}
}

// BeginTx starts a new transaction (no-op for in-memory)
func (d *Database) BeginTx(ctx context.Context) (repositories.Database, error) {
	// For in-memory, we just return the same instance
	// In a real implementation, you might want to implement transaction isolation
	return d, nil
}

// Commit commits the transaction (no-op for in-memory)
func (d *Database) Commit() error {
	return nil
}

// Rollback rolls back the transaction (no-op for in-memory)
func (d *Database) Rollback() error {
	return nil
}

// Close closes the database (no-op for in-memory)
func (d *Database) Close() error {
	return nil
}

// Common error
var ErrNotFound = fmt.Errorf("not found")

