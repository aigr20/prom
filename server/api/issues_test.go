package api

import (
	"aigr20/prom/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func issueFromBody(t *testing.T, body []byte) models.Issue {
	data := make(map[string]models.Issue)
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&data)
	if err != nil {
		t.FailNow()
	}

	return data["data"]
}

func TestCreateIssueRoute(t *testing.T) {
	tests := []struct {
		name              string
		body              models.IssueCreateForm
		wantedCode        int
		wantedTitle       string
		wantedDescription string
		wantedProject     int
	}{
		{
			name: "title&description",
			body: models.IssueCreateForm{
				Title:       "new issue",
				Description: "description",
				ProjectID:   1,
			},
			wantedCode:        http.StatusCreated,
			wantedTitle:       "new issue",
			wantedDescription: "description",
			wantedProject:     1,
		},
		{
			name: "title_no_description",
			body: models.IssueCreateForm{
				Title:     "new issue",
				ProjectID: 2,
			},
			wantedCode:        http.StatusCreated,
			wantedTitle:       "new issue",
			wantedDescription: "",
			wantedProject:     2,
		},
		{
			name: "missing_title",
			body: models.IssueCreateForm{
				Description: "description",
				ProjectID:   1,
			},
			wantedCode: http.StatusBadRequest,
		},
		{
			name: "missing_project",
			body: models.IssueCreateForm{
				Title:       "new issue",
				Description: "description",
			},
			wantedCode: http.StatusBadRequest,
		},
		{
			name: "project_doesnt_exist",
			body: models.IssueCreateForm{
				Title:       "new issue",
				Description: "description",
				ProjectID:   100,
			},
			wantedCode: http.StatusBadRequest,
		},
	}

	api := getTestAPI(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Cleanup(func() {
				api.IssueRepo.CustomQuery("DELETE FROM issues WHERE issue_id > 4")
				api.IssueRepo.CustomQuery("ALTER TABLE issues AUTO_INCREMENT = 5")
			})

			w := httptest.NewRecorder()
			bodyMarshal, err := json.Marshal(test.body)
			if err != nil {
				t.Error(err)
			}
			reader := bytes.NewReader(bodyMarshal)
			req, _ := http.NewRequest("POST", "/issues/create", reader)
			api.Router.ServeHTTP(w, req)

			if w.Code != test.wantedCode {
				t.FailNow()
			}

			issue := issueFromBody(t, w.Body.Bytes())
			if issue.Title != test.wantedTitle || issue.Description != test.wantedDescription || issue.ProjectID != test.wantedProject {
				t.FailNow()
			}
		})
	}
}
