package database

import (
	"aigr20/prom/models"
	"testing"
	"time"
)

var sampleIssues = []models.Issue{
	{
		ID:          1,
		Title:       "Make accessible",
		Description: "Appen måste gå att använda av alla!",
		Created:     time.Date(2023, time.June, 28, 14, 0, 0, 0, time.Local),
		Updated:     time.Date(2023, time.June, 28, 15, 15, 0, 0, time.Local),
		ProjectID:   1,
	},
	{
		ID:          2,
		Title:       "Make more fun",
		Description: "Lägg in mycket färger, appen skall vara rolig!",
		Created:     time.Date(2023, time.June, 28, 14, 3, 0, 0, time.Local),
		Updated:     time.Date(2023, time.June, 28, 15, 20, 0, 0, time.Local),
		ProjectID:   1,
	},
	{
		ID:          3,
		Title:       "Update damage system",
		Description: "Damage system must be more complicated",
		Created:     time.Date(2023, time.June, 29, 13, 12, 3, 0, time.Local),
		Updated:     time.Date(2023, time.June, 29, 16, 12, 0, 0, time.Local),
		ProjectID:   2,
	},
	{
		ID:          4,
		Title:       "Receptmodell",
		Description: "",
		Created:     time.Date(2023, time.June, 30, 10, 0, 0, 0, time.Local),
		Updated:     time.Date(2023, time.June, 30, 15, 0, 34, 0, time.Local),
		ProjectID:   3,
	},
}

func getIssueRepository(t *testing.T) *IssueRepository {
	db, err := CreateConnection("prom_test", "prom_tester", "tester")
	if err != nil {
		t.Error("Failed on database connection")
	}

	return NewIssueRepository(db)
}

func TestGetAllForProjectCount(t *testing.T) {
	tests := []struct {
		name    string
		project int
		want    int
	}{
		{
			name:    "project_1",
			project: 1,
			want:    2,
		},
		{
			name:    "project_2",
			project: 2,
			want:    1,
		},
		{
			name:    "project_3",
			project: 3,
			want:    1,
		},
	}

	repo := getIssueRepository(t)
	defer repo.Close()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			issues, err := repo.GetIssuesFromProject(test.project)
			if err != nil {
				t.Error(err)
			}
			if len(issues) != test.want {
				t.Fail()
			}
		})
	}
}

func TestGetAllForProjectContent(t *testing.T) {
	tests := []struct {
		name    string
		project int
		want    []models.Issue
	}{
		{
			name:    "project_1",
			project: 1,
			want:    sampleIssues[0:2],
		},
		{
			name:    "project_2",
			project: 2,
			want:    sampleIssues[2:3],
		},
		{
			name:    "project_3",
			project: 3,
			want:    sampleIssues[3:],
		},
	}

	repo := getIssueRepository(t)
	defer repo.Close()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			issues, err := repo.GetIssuesFromProject(test.project)
			if err != nil {
				t.Error(err)
			}
			for i, issue := range test.want {
				if issue != issues[i] {
					t.Fail()
				}
			}
		})
	}
}
