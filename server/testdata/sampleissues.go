package testdata

import (
	"aigr20/prom/models"
	"time"
)

var SampleIssues = []models.Issue{
	{
		ID:          1,
		Title:       "Make accessible",
		Description: "Appen måste gå att använda av alla!",
		Estimate:    5,
		Created:     time.Date(2023, time.June, 28, 14, 0, 0, 0, time.Local),
		Updated:     time.Date(2023, time.June, 28, 15, 15, 0, 0, time.Local),
		ProjectID:   1,
		Status:      "TODO",
		Tags: []models.IssueTag{
			{
				ID:    1,
				Text:  "important",
				Color: "#ff0000",
			},
			{
				ID:    2,
				Text:  "bug",
				Color: "#0000ff",
			},
		},
	},
	{
		ID:          2,
		Title:       "Make more fun",
		Description: "Lägg in mycket färger, appen skall vara rolig!",
		Estimate:    3,
		Created:     time.Date(2023, time.June, 28, 14, 3, 0, 0, time.Local),
		Updated:     time.Date(2023, time.June, 28, 15, 20, 0, 0, time.Local),
		ProjectID:   1,
		Status:      "In Progress",
		Tags: []models.IssueTag{
			{
				ID:    1,
				Text:  "important",
				Color: "#ff0000",
			},
		},
	},
	{
		ID:          3,
		Title:       "Update damage system",
		Description: "Damage system must be more complicated",
		Estimate:    2,
		Created:     time.Date(2023, time.June, 29, 13, 12, 3, 0, time.Local),
		Updated:     time.Date(2023, time.June, 29, 16, 12, 0, 0, time.Local),
		ProjectID:   2,
		Status:      "Finished",
		Tags: []models.IssueTag{
			{
				ID:    2,
				Text:  "bug",
				Color: "#0000ff",
			},
		},
	},
	{
		ID:          4,
		Title:       "Receptmodell",
		Description: "",
		Estimate:    1,
		Created:     time.Date(2023, time.June, 30, 10, 0, 0, 0, time.Local),
		Updated:     time.Date(2023, time.June, 30, 15, 0, 34, 0, time.Local),
		ProjectID:   3,
		Status:      "TODO",
		Tags:        []models.IssueTag{},
	},
}
