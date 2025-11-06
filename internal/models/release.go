package models // domain models

import "time"

// Release is an annotation on the timeline used to track significant events or changes
// It helps correlate sitemap changes with deployments or other events in timeseries analysis
type Release struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	ReleaseDate time.Time `json:"release_date"`
	Version     *string   `json:"version,omitempty"`
	ReleaseNotes *string  `json:"release_notes,omitempty"`
	
	// Optional: link to a report schedule for context in timeseries comparisons
	ReportScheduleID *string `json:"report_schedule_id,omitempty"`
	
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
