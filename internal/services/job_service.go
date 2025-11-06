package services

import (
	"context"

	"jonopens/sitemapper/internal/models"
	"jonopens/sitemapper/internal/repositories"
	"jonopens/sitemapper/pkg/scheduler"
)

// JobService handles background job scheduling and execution
type JobService struct {
	db        repositories.Database
	scheduler *scheduler.Scheduler
}

// NewJobService creates a new job service
func NewJobService(db repositories.Database) *JobService {
	return &JobService{
		db:        db,
		scheduler: scheduler.New(),
	}
}

// CreateJob creates a new report job
func (s *JobService) CreateJob(ctx context.Context, job *models.ReportJob) error {
	return s.db.ReportJobs().Create(ctx, job)
}

// ScheduleJob schedules a recurring sitemap check
func (s *JobService) ScheduleJob(ctx context.Context, jobID string, cronExpr string) error {
	// TODO: Implement job scheduling logic
	return nil
}

// ExecuteJob executes a report job
func (s *JobService) ExecuteJob(ctx context.Context, jobID string) error {
	// TODO: Implement job execution logic
	return nil
}

