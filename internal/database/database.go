package database

import (
	"fmt"

	"jonopens/sitemapper/internal/config"
	"jonopens/sitemapper/internal/database/memory"
	"jonopens/sitemapper/internal/database/mysql"
	"jonopens/sitemapper/internal/database/postgres"
	"jonopens/sitemapper/internal/database/sqlite"
	"jonopens/sitemapper/internal/repositories"
)

// NewDatabase creates a database instance based on configuration
func NewDatabase(cfg *config.Config) (repositories.Database, error) {
	switch cfg.DatabaseType {
	case "postgres", "postgresql":
		return postgres.New(cfg.DatabaseURL)
	case "mysql":
		return mysql.New(cfg.DatabaseURL)
	case "sqlite", "sqlite3":
		return sqlite.New(cfg.DatabaseURL)
	case "memory":
		return memory.New(), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.DatabaseType)
	}
}

