package models // domain models

import "time"
// a Group is a url segment or a manually assigned Group that defines entries
// a user should be able to override an assigned Group for a specific entry
type Group struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}