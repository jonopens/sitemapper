package services

import (
	"context"

	"jonopens/sitemapper/internal/repositories"
	"jonopens/sitemapper/pkg/sitemap"
)

// SitemapService handles sitemap processing business logic
type SitemapService struct {
	db     repositories.Database
	parser *sitemap.Parser
}

// NewSitemapService creates a new sitemap service
func NewSitemapService(db repositories.Database) *SitemapService {
	return &SitemapService{
		db:     db,
		parser: sitemap.NewParser(),
	}
}

// ProcessSitemap parses and processes a sitemap
func (s *SitemapService) ProcessSitemap(ctx context.Context, userID string, data []byte) error {
	// TODO: Implement sitemap processing logic
	// 1. Parse sitemap using parser
	// 2. Categorize entries
	// 3. Store in database
	return nil
}

// CompareSitemaps compares two sitemaps and returns differences
func (s *SitemapService) CompareSitemaps(ctx context.Context, oldID, newID string) error {
	// TODO: Implement comparison logic
	return nil
}

