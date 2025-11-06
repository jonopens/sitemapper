package models // domain models

import "time"

type JobSchedule string

const (
	JobScheduleImmediate JobSchedule = "immediate"
	JobScheduleDaily     JobSchedule = "daily"
	JobScheduleWeekly    JobSchedule = "weekly"
	JobScheduleMonthly   JobSchedule = "monthly"
)

// ReportSchedule defines a recurring schedule for generating reports
type ReportSchedule struct {
	ID     string      `json:"id"`
	UserID string      `json:"user_id"`
	Name   string      `json:"name"`
	Schedule JobSchedule `json:"schedule"`

	// Job configuration (used to create new jobs on schedule)
	CompressionFormat          *CompressionFormat `json:"compression_format,omitempty"`
	SourceLocation             string             `json:"source_location"` // URL or file path
	JobType                    JobType            `json:"job_type"`
	ShouldCheckEntryLiveness   bool               `json:"should_check_entry_liveness"`
	ShouldCheckForValidEntries bool               `json:"should_check_for_valid_entries"`
	MaxStoredEntries           int                `json:"max_stored_entries"`

	// Schedule management
	IsActive  bool       `json:"is_active"`
	LastRunAt *time.Time `json:"last_run_at,omitempty"`
	NextRunAt *time.Time `json:"next_run_at,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s JobSchedule) AsCronExpression() *string {
	switch s {
	case JobScheduleDaily:
		cron := "0 0 * * *" // every day at midnight
		return &cron
	case JobScheduleWeekly:
		cron := "0 0 * * 0" // every Sunday at midnight
		return &cron
	case JobScheduleMonthly:
		cron := "0 0 1 * *" // every month on the first day at midnight
		return &cron
	}
	return nil // no cron, immediate
}
