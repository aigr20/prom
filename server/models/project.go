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

type ProjectCreateForm struct {
	Name string `json:"name" binding:"required"`
}
