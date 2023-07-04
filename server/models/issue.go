package models

import "time"

type Issue struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Created     time.Time `json:"createdAt"`
	Updated     time.Time `json:"updatedAt"`
	ProjectID   int       `json:"-"`
}
