package models // domain models

import "time"

// EntryType indicates the type of entry (URL or Sitemap)
type EntryType string

const (
	EntryTypeURL     EntryType = "url"     // Regular URL entry from a sitemap
	EntryTypeSitemap EntryType = "sitemap" // Sitemap URL from a sitemap index
)

// SelectionReason indicates why an entry was stored in the database
type SelectionReason string

const (
	SelectionReasonFullStorage SelectionReason = "full_storage" // sitemap under threshold
	SelectionReasonSampled     SelectionReason = "sampled"      // part of stratified sample
	SelectionReasonOutlier     SelectionReason = "outlier"      // error/down/invalid
	SelectionReasonBoundary    SelectionReason = "boundary"     // min/max URL per grouping
)

// Entry represents a single URL from a sitemap
type Entry struct {
	ID         string    `json:"id"`
	ReportID   string    `json:"report_id"`
	GroupingID *string   `json:"grouping_id,omitempty"` // nullable if ungrouped
	Type       EntryType `json:"type"`                  // url or sitemap
	URL        string    `json:"url"`

	// Sitemap metadata fields
	LastModified *time.Time `json:"last_modified,omitempty"`
	ChangeFreq   *string    `json:"change_freq,omitempty"`
	Priority     *float64   `json:"priority,omitempty"`

	// Validation fields
	IsValid         bool    `json:"is_valid"`
	ValidationError *string `json:"validation_error,omitempty"`

	// Liveness check fields (null if not checked)
	HTTPStatusCode    *int       `json:"http_status_code,omitempty"`
	IsLive            *bool      `json:"is_live,omitempty"`
	ResponseTimeMs    *int       `json:"response_time_ms,omitempty"`
	LivenessCheckedAt *time.Time `json:"liveness_checked_at,omitempty"`
	LivenessError     *string    `json:"liveness_error,omitempty"`

	// Sampling metadata
	SelectionReason SelectionReason `json:"selection_reason"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

