package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"jonopens/sitemapper/internal/repositories"
)

// Database implements repositories.Database for SQLite
type Database struct {
	db *sql.DB
	tx *sql.Tx
}

// New creates a new SQLite database connection
func New(connectionString string) (*Database, error) {
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{db: db}, nil
}

// Entries returns the entry repository
func (d *Database) Entries() repositories.EntryRepository {
	return &EntryRepository{db: d.db, tx: d.tx}
}

// Reports returns the report repository
func (d *Database) Reports() repositories.ReportRepository {
	return &ReportRepository{db: d.db, tx: d.tx}
}

// Users returns the user repository
func (d *Database) Users() repositories.UserRepository {
	return &UserRepository{db: d.db, tx: d.tx}
}

// Groupings returns the grouping repository
func (d *Database) Groupings() repositories.GroupingRepository {
	return &GroupingRepository{db: d.db, tx: d.tx}
}

// ReportJobs returns the report job repository
func (d *Database) ReportJobs() repositories.ReportJobRepository {
	return &ReportJobRepository{db: d.db, tx: d.tx}
}

// Releases returns the release repository
func (d *Database) Releases() repositories.ReleaseRepository {
	return &ReleaseRepository{db: d.db, tx: d.tx}
}

// BeginTx starts a new transaction
func (d *Database) BeginTx(ctx context.Context) (repositories.Database, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &Database{db: d.db, tx: tx}, nil
}

// Commit commits the transaction
func (d *Database) Commit() error {
	if d.tx == nil {
		return fmt.Errorf("no transaction to commit")
	}
	return d.tx.Commit()
}

// Rollback rolls back the transaction
func (d *Database) Rollback() error {
	if d.tx == nil {
		return fmt.Errorf("no transaction to rollback")
	}
	return d.tx.Rollback()
}

// Close closes the database connection
func (d *Database) Close() error {
	return d.db.Close()
}

