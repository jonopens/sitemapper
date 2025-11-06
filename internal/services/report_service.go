package services

import (
	"context"

	"jonopens/sitemapper/internal/models"
	"jonopens/sitemapper/internal/repositories"
)

// ReportService handles report generation business logic
type ReportService struct {
	db repositories.Database
}

// NewReportService creates a new report service
func NewReportService(db repositories.Database) *ReportService {
	return &ReportService{db: db}
}

// GenerateReport creates a comprehensive sitemap report
func (s *ReportService) GenerateReport(ctx context.Context, userID string, sitemapID string) (*models.Report, error) {
	// TODO: Implement report generation logic
	// 1. Fetch sitemap entries
	// 2. Calculate metrics
	// 3. Generate report
	// 4. Store report
	return nil, nil
}

// GetReport retrieves a report by ID
func (s *ReportService) GetReport(ctx context.Context, reportID string) (*models.Report, error) {
	return s.db.Reports().GetByID(ctx, reportID)
}

// ListReportsByUser lists all reports for a user
func (s *ReportService) ListReportsByUser(ctx context.Context, userID string) ([]*models.Report, error) {
	return s.db.Reports().GetByUserID(ctx, userID)
}

