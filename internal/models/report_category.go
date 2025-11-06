package models // domain models

import "time"

// ReportGrouping links a report to a grouping and stores aggregate data
type ReportGrouping struct {
	ID         string `json:"id"`
	ReportID   string `json:"report_id"`
	GroupingID string `json:"grouping_id"`

	// Aggregate counts (always accurate, from full sitemap processing)
	TotalEntryCount   int `json:"total_entry_count"`
	StoredEntryCount  int `json:"stored_entry_count"`
	LiveEntryCount    int `json:"live_entry_count"`
	DownEntryCount    int `json:"down_entry_count"`
	ValidEntryCount   int `json:"valid_entry_count"`
	InvalidEntryCount int `json:"invalid_entry_count"`

	// Boundary information for this grouping
	MinURL *string `json:"min_url,omitempty"` // alphabetically first URL
	MaxURL *string `json:"max_url,omitempty"` // alphabetically last URL

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}