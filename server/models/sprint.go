package models

import "time"

type Sprint struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Start    time.Time `json:"start_date"`
	End      time.Time `json:"end_date"`
	Finished bool      `json:"-"`
	Current  bool      `json:"-"`
}
