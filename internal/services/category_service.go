package services

import (
	"context"

	"jonopens/sitemapper/internal/models"
	"jonopens/sitemapper/internal/repositories"
)

// CategoryService handles category management
type CategoryService struct {
	db repositories.Database
}

// NewCategoryService creates a new category service
func NewCategoryService(db repositories.Database) *CategoryService {
	return &CategoryService{db: db}
}

// CreateCategory creates a new category
func (s *CategoryService) CreateCategory(ctx context.Context, category *models.Category) error {
	return s.db.Categories().Create(ctx, category)
}

// GetCategory retrieves a category by ID
func (s *CategoryService) GetCategory(ctx context.Context, id string) (*models.Category, error) {
	return s.db.Categories().GetByID(ctx, id)
}

// ListCategories lists all categories
func (s *CategoryService) ListCategories(ctx context.Context) ([]*models.Category, error) {
	return s.db.Categories().List(ctx)
}

// CategorizeURLs automatically categorizes URLs based on path segments
func (s *CategoryService) CategorizeURLs(ctx context.Context, urls []string) (map[string][]string, error) {
	// TODO: Implement URL categorization logic
	return nil, nil
}

