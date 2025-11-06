package repositories

import (
	"context"

	"jonopens/sitemapper/internal/models"
)

// Database aggregates all repositories and provides transaction support
type Database interface {
	Entries() EntryRepository
	Reports() ReportRepository
	Users() UserRepository
	Categories() CategoryRepository
	ReportJobs() ReportJobRepository
	Releases() ReleaseRepository
	
	// Transaction support
	BeginTx(ctx context.Context) (Database, error)
	Commit() error
	Rollback() error
	Close() error
}

// EntryRepository defines the contract for entry data access
type EntryRepository interface {
	Create(ctx context.Context, entry *models.Entry) error
	GetByID(ctx context.Context, id string) (*models.Entry, error)
	List(ctx context.Context, filters EntryFilters) ([]*models.Entry, error)
	Update(ctx context.Context, entry *models.Entry) error
	Delete(ctx context.Context, id string) error
	CountByType(ctx context.Context, entryType models.EntryType) (int, error)
}

// ReportRepository defines the contract for report data access
type ReportRepository interface {
	Create(ctx context.Context, report *models.Report) error
	GetByID(ctx context.Context, id string) (*models.Report, error)
	GetByUserID(ctx context.Context, userID string) ([]*models.Report, error)
	List(ctx context.Context, filters ReportFilters) ([]*models.Report, error)
	Update(ctx context.Context, report *models.Report) error
	Delete(ctx context.Context, id string) error
}

// UserRepository defines the contract for user data access
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}

// CategoryRepository defines the contract for category data access
type CategoryRepository interface {
	Create(ctx context.Context, category *models.Category) error
	GetByID(ctx context.Context, id string) (*models.Category, error)
	List(ctx context.Context) ([]*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id string) error
}

// ReportJobRepository defines the contract for report job data access
type ReportJobRepository interface {
	Create(ctx context.Context, job *models.ReportJob) error
	GetByID(ctx context.Context, id string) (*models.ReportJob, error)
	List(ctx context.Context, filters JobFilters) ([]*models.ReportJob, error)
	Update(ctx context.Context, job *models.ReportJob) error
	Delete(ctx context.Context, id string) error
}

// ReleaseRepository defines the contract for release data access
type ReleaseRepository interface {
	Create(ctx context.Context, release *models.Release) error
	GetByID(ctx context.Context, id string) (*models.Release, error)
	List(ctx context.Context, filters ReleaseFilters) ([]*models.Release, error)
	Update(ctx context.Context, release *models.Release) error
	Delete(ctx context.Context, id string) error
}

// Filter types for querying
type EntryFilters struct {
	ReportID string
	Type     *models.EntryType
	Limit    int
	Offset   int
}

type ReportFilters struct {
	UserID string
	Limit  int
	Offset int
}

type JobFilters struct {
	Status string
	UserID string
	Limit  int
	Offset int
}

type ReleaseFilters struct {
	UserID string
	Limit  int
	Offset int
}

