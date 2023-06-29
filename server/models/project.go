package models

import "time"

type Project struct {
	ID      int
	Name    string
	Created time.Time
	Updated time.Time
}
