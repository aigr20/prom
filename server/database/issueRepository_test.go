package database

import (
	"aigr20/prom/models"
	"aigr20/prom/testdata"
	"testing"
	"time"
)

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
			want:    testdata.SampleIssues[0:2],
		},
		{
			name:    "project_2",
			project: 2,
			want:    testdata.SampleIssues[2:3],
		},
		{
			name:    "project_3",
			project: 3,
			want:    testdata.SampleIssues[3:],
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
				if !issue.LenientEquals(issues[i]) {
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
		wantedEstimate    int
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
			name: "with_estimate",
			body: models.IssueCreateForm{
				Title:     "new issue",
				Estimate:  3,
				ProjectID: 1,
			},
			wantedTitle:    "new issue",
			wantedProject:  1,
			wantedStatus:   "TODO",
			wantedEstimate: 3,
			wantedError:    nil,
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

			if issue.Title != test.wantedTitle || issue.Description != test.wantedDescription || issue.ProjectID != test.wantedProject || issue.Status != test.wantedStatus || issue.Estimate != test.wantedEstimate {
				t.FailNow()
			}
		})
	}
}

func TestUpdateIssue(t *testing.T) {
	tests := []struct {
		name          string
		target        int
		fields        []string
		values        []any
		expectedIssue models.Issue
		expectedError error
	}{
		{
			name:   "success_multifield",
			target: 1,
			fields: []string{"issue_title", "issue_description"},
			values: []any{"new title", "new description"},
			expectedIssue: models.Issue{
				ID:          1,
				Title:       "new title",
				Description: "new description",
				Estimate:    5,
				Created:     time.Date(2023, time.June, 28, 14, 0, 0, 0, time.Local),
				ProjectID:   1,
				Status:      "TODO",
			},
			expectedError: nil,
		},
		{
			name:   "success_onefield",
			target: 1,
			fields: []string{"issue_title"},
			values: []any{"new title"},
			expectedIssue: models.Issue{
				ID:          1,
				Title:       "new title",
				Description: "Appen måste gå att använda av alla!",
				Estimate:    5,
				Created:     time.Date(2023, time.June, 28, 14, 0, 0, 0, time.Local),
				ProjectID:   1,
				Status:      "TODO",
			},
			expectedError: nil,
		},
		{
			name:          "fail_mismatched_field_count_gt",
			target:        1,
			fields:        []string{"issue_title", "issue_description"},
			values:        []any{"test"},
			expectedIssue: models.Issue{},
			expectedError: ErrUpdateFieldCount,
		},
		{
			name:          "fail_mismatched_field_count_lt",
			target:        1,
			fields:        []string{"issue_title"},
			values:        []any{"new title", "new description"},
			expectedIssue: models.Issue{},
			expectedError: ErrUpdateFieldCount,
		},
		{
			name:          "fail_no_fields",
			target:        1,
			fields:        []string{},
			values:        []any{},
			expectedIssue: models.Issue{},
			expectedError: ErrNoFields,
		},
		{
			name:          "fail_issue_doesn't_exist",
			target:        100,
			fields:        []string{"issue_title"},
			values:        []any{"new title"},
			expectedIssue: models.Issue{},
			expectedError: ErrIssueNotFound,
		},
		{
			name:          "fail_illegal_column",
			target:        1,
			fields:        []string{"hihi"},
			values:        []any{"hoho"},
			expectedIssue: models.Issue{},
			expectedError: ErrIllegalFieldName,
		},
	}

	repo := getIssueRepository(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Cleanup(func() {
				original := &testdata.SampleIssues[0]
				repo.CustomQuery("UPDATE issues SET issue_title=?, issue_description=?, last_changed=?, project=?, issue_status=? WHERE issue_id = 1", original.Title, original.Description, original.Updated, original.ProjectID, 1)
			})

			issue, err := repo.UpdateIssue(test.target, test.fields, test.values)
			if err != test.expectedError {
				t.FailNow()
			}
			if !issue.LenientEquals(test.expectedIssue) {
				t.FailNow()
			}
		})
	}
}
