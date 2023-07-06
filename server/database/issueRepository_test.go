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
		Status:      "TODO",
	},
	{
		ID:          2,
		Title:       "Make more fun",
		Description: "Lägg in mycket färger, appen skall vara rolig!",
		Created:     time.Date(2023, time.June, 28, 14, 3, 0, 0, time.Local),
		Updated:     time.Date(2023, time.June, 28, 15, 20, 0, 0, time.Local),
		ProjectID:   1,
		Status:      "In Progress",
	},
	{
		ID:          3,
		Title:       "Update damage system",
		Description: "Damage system must be more complicated",
		Created:     time.Date(2023, time.June, 29, 13, 12, 3, 0, time.Local),
		Updated:     time.Date(2023, time.June, 29, 16, 12, 0, 0, time.Local),
		ProjectID:   2,
		Status:      "Finished",
	},
	{
		ID:          4,
		Title:       "Receptmodell",
		Description: "",
		Created:     time.Date(2023, time.June, 30, 10, 0, 0, 0, time.Local),
		Updated:     time.Date(2023, time.June, 30, 15, 0, 34, 0, time.Local),
		ProjectID:   3,
		Status:      "TODO",
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

func TestCreateIssue(t *testing.T) {
	tests := []struct {
		name              string
		body              models.IssueCreateForm
		wantedTitle       string
		wantedDescription string
		wantedProject     int
		wantedStatus      string
		wantedError       error
	}{
		{
			name: "title&description",
			body: models.IssueCreateForm{
				Title:       "new issue",
				Description: "description",
				ProjectID:   1,
			},
			wantedTitle:       "new issue",
			wantedDescription: "description",
			wantedProject:     1,
			wantedStatus:      "TODO",
			wantedError:       nil,
		},
		{
			name: "title_no_description",
			body: models.IssueCreateForm{
				Title:     "new issue",
				ProjectID: 2,
			},
			wantedTitle:       "new issue",
			wantedDescription: "",
			wantedProject:     2,
			wantedStatus:      "TODO",
			wantedError:       nil,
		},
		{
			name: "invalid_project",
			body: models.IssueCreateForm{
				Title:     "invalid project",
				ProjectID: 6,
			},
			wantedTitle:       "",
			wantedDescription: "",
			wantedProject:     0,
			wantedStatus:      "",
			wantedError:       ErrIssueCreate,
		},
		{
			name: "missing_title",
			body: models.IssueCreateForm{
				Description: "description",
				ProjectID:   2,
			},
			wantedTitle:       "",
			wantedDescription: "",
			wantedProject:     0,
			wantedStatus:      "",
			wantedError:       ErrIssueCreate,
		},
		{
			name: "missing_project",
			body: models.IssueCreateForm{
				Title:       "invalid project",
				Description: "description",
			},
			wantedTitle:       "",
			wantedDescription: "",
			wantedProject:     0,
			wantedStatus:      "",
			wantedError:       ErrIssueCreate,
		},
	}

	repo := getIssueRepository(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Cleanup(func() {
				repo.CustomQuery("DELETE FROM issues WHERE issue_id > 4")
				repo.CustomQuery("ALTER TABLE issues AUTO_INCREMENT = 5")
			})

			issue, err := repo.CreateIssue(test.body)
			if err != test.wantedError {
				t.FailNow()
			}

			if issue.Title != test.wantedTitle || issue.Description != test.wantedDescription || issue.ProjectID != test.wantedProject || issue.Status != test.wantedStatus {
				t.FailNow()
			}
		})
	}
}
