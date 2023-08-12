package testdata

import (
	"aigr20/prom/models"
	"time"
)

var SampleProjects = []models.Project{
	{
		ID:   1,
		Name: "Glada Schemat",
		Tags: []models.IssueTag{
			{Text: "important", Color: "#ff0000"},
			{Text: "bug", Color: "#0000ff"},
		},
		Created: time.Date(2023, time.June, 24, 10, 0, 0, 0, time.Local),
		Updated: time.Date(2023, time.June, 24, 10, 15, 0, 0, time.Local),
	},
	{
		ID:   2,
		Name: "Max Power",
		Tags: []models.IssueTag{
			{Text: "bug", Color: "#0000ff"},
		},
		Created: time.Date(2023, time.June, 24, 9, 10, 0, 0, time.Local),
		Updated: time.Date(2023, time.June, 24, 16, 13, 12, 0, time.Local),
	},
	{
		ID:      3,
		Name:    "Mina Recept",
		Tags:    []models.IssueTag{},
		Created: time.Date(2023, time.June, 25, 15, 22, 35, 0, time.Local),
		Updated: time.Date(2023, time.June, 26, 12, 0, 14, 0, time.Local),
	},
}
