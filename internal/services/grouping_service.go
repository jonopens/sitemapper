package services

import (
	"context"

	"jonopens/sitemapper/internal/models"
	"jonopens/sitemapper/internal/repositories"
)

// GroupingService handles grouping management
type GroupingService struct {
	db repositories.Database
}

// NewGroupingService creates a new grouping service
func NewGroupingService(db repositories.Database) *GroupingService {
	return &GroupingService{db: db}
}

// CreateGrouping creates a new grouping
func (s *GroupingService) CreateGrouping(ctx context.Context, grouping *models.Group) error {
	return s.db.Groupings().Create(ctx, grouping)
}

// GetGrouping retrieves a grouping by ID
func (s *GroupingService) GetGrouping(ctx context.Context, id string) (*models.Group, error) {
	return s.db.Groupings().GetByID(ctx, id)
}

// ListGroupings lists all groupings
func (s *GroupingService) ListGroupings(ctx context.Context) ([]*models.Group, error) {
	return s.db.Groupings().List(ctx)
}

// GroupURLs automatically groups URLs based on path segments
func (s *GroupingService) GroupURLs(ctx context.Context, urls []string) (map[string][]string, error) {
	// TODO: Implement URL grouping logic
	return nil, nil
}

