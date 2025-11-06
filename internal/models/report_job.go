package models // domain models

import "time"

const (
	ReportJobStatusPending = "pending"
	ReportJobStatusRunning = "running"
	ReportJobStatusCompleted = "completed"
	ReportJobStatusFailed = "failed"
	ReportJobStatusCancelled = "cancelled"
	ReportJobStatusTimedOut = "timed_out"
)

type JobType string

const (
	JobTypeURL JobType = "url"
	JobTypeUpload JobType = "upload"
)

type CompressionFormat string

const (
	CompressionFormatGzip CompressionFormat = "gzip"
	CompressionFormatZip CompressionFormat = "zip"
)

// ReportJob is a job that generates a report from a sitemap
type ReportJob struct {
	ID               string  `json:"id"`
	UserID           string  `json:"user_id"`
	ReportID         *string `json:"report_id,omitempty"`         // populated after report is created
	ReportScheduleID *string `json:"report_schedule_id,omitempty"` // null if ad-hoc job

	// Source configuration
	CompressionFormat  *CompressionFormat `json:"compression_format,omitempty"` // null if uncompressed
	SourceLocation     string             `json:"source_location"`              // URL or file path
	JobType            JobType            `json:"job_type"`

	// Processing configuration
	ShouldCheckEntryLiveness   bool `json:"should_check_entry_liveness"`
	ShouldCheckForValidEntries bool `json:"should_check_for_valid_entries"`
	MaxStoredEntries           int  `json:"max_stored_entries"` // threshold for sampling (0 = store all)

	// Job status
	Status      string     `json:"status"`
	ErrorMessage *string   `json:"error_message,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}