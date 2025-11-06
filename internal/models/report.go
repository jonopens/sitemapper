package models // domain models

import "time"

// SamplingStrategy describes how entries were stored
type SamplingStrategy string

const (
	SamplingStrategyNone       SamplingStrategy = "none"       // all entries stored
	SamplingStrategyStratified SamplingStrategy = "stratified" // sampled by grouping
	SamplingStrategyRandom     SamplingStrategy = "random"     // random sample
)

// Report is a summary generated from a processed sitemap
type Report struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`

	// Entry counts (always accurate totals from sitemap)
	EntryCount        int `json:"entry_count"`
	StoredEntryCount  int `json:"stored_entry_count"`
	ValidEntryCount   int `json:"valid_entry_count"`
	InvalidEntryCount int `json:"invalid_entry_count"`
	LiveEntryCount    int `json:"live_entry_count"`
	DownEntryCount    int `json:"down_entry_count"`

	// Grouping info
	GroupingCount  int `json:"grouping_count"`
	UngroupedCount int `json:"ungrouped_count"`

	// Sitemap structure info
	ChildSitemapCount int `json:"child_sitemap_count"`

	// Sampling metadata
	IsFullyStored    bool             `json:"is_fully_stored"`
	SamplingStrategy SamplingStrategy `json:"sampling_strategy"`
	SamplingRate     *float64         `json:"sampling_rate,omitempty"` // e.g., 0.1 for 10%

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}