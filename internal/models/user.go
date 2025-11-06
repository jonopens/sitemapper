package models // domain models

import "time"

// a user is a user of the system
type User struct {
	ID string `json:"id"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
	SlackID *string `json:"slack_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}