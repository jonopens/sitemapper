package models // domain models

import "time"

// ReportDiff represents a comparison between two reports
// This can be used to show how a sitemap has changed over time
type ReportDiff struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	BaseReportID    string    `json:"base_report_id"`    // older report
	CompareReportID string    `json:"compare_report_id"` // newer report
	
	// Summary statistics
	EntriesAdded   int `json:"entries_added"`
	EntriesRemoved int `json:"entries_removed"`
	EntriesChanged int `json:"entries_changed"`
	
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

