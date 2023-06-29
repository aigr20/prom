package models

import (
	"time"
)

type Project struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"createdAt"`
	Updated time.Time `json:"updatedAt"`
}
