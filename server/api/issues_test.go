package api

import (
	"aigr20/prom/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func issueFromBody(t *testing.T, body []byte) models.Issue {
	data := make(map[string]models.Issue)
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&data)
	if err != nil {
		return models.Issue{}
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
		wantedStatus      string
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
			wantedStatus:      "TODO",
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
			wantedStatus:      "TODO",
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
				t.Log("Failing code")
				t.FailNow()
			}

			issue := issueFromBody(t, w.Body.Bytes())
			if issue.Title != test.wantedTitle || issue.Description != test.wantedDescription || issue.Status != test.wantedStatus {
				t.Log("failing content")
				t.FailNow()
			}
		})
	}
}

func TestUpdateStatusRoute(t *testing.T) {
	tests := []struct {
		name       string
		body       models.UpdateIssueStatusBody
		wantedCode int
	}{
		{
			name: "success",
			body: models.UpdateIssueStatusBody{
				IssueID:   1,
				NewStatus: "Finished",
			},
			wantedCode: http.StatusNoContent,
		},
		{
			name: "fail_missing_id",
			body: models.UpdateIssueStatusBody{
				NewStatus: "Finished",
			},
			wantedCode: http.StatusBadRequest,
		},
		{
			name: "fail_missing_status",
			body: models.UpdateIssueStatusBody{
				IssueID: 1,
			},
			wantedCode: http.StatusBadRequest,
		},
		{
			name: "fail_invalid_issue",
			body: models.UpdateIssueStatusBody{
				IssueID:   100,
				NewStatus: "Finished",
			},
			wantedCode: http.StatusBadRequest,
		},
		{
			name: "fail_status_doesn't_exist",
			body: models.UpdateIssueStatusBody{
				IssueID:   1,
				NewStatus: "Will do",
			},
			wantedCode: http.StatusBadRequest,
		},
	}

	api := getTestAPI(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Cleanup(func() {
				api.IssueRepo.CustomQuery("UPDATE issues SET issue_status=1, last_changed=? WHERE issue_id = 1", time.Date(2023, time.June, 28, 15, 15, 0, 0, time.Local))
			})

			w := httptest.NewRecorder()
			bodyMarshal, err := json.Marshal(test.body)
			if err != nil {
				t.Error(err)
			}
			reader := bytes.NewReader(bodyMarshal)
			req, _ := http.NewRequest("PATCH", "/issues/status", reader)
			api.Router.ServeHTTP(w, req)

			if w.Code != test.wantedCode {
				t.FailNow()
			}
		})
	}
}

func TestUpdateEstimateRoute(t *testing.T) {
	tests := []struct {
		name       string
		body       models.UpdateEstimateBody
		wantedCode int
	}{
		{
			name:       "success",
			body:       models.UpdateEstimateBody{IssueID: 1, NewEstimate: 1},
			wantedCode: http.StatusNoContent,
		},
		{
			name:       "fail_missing_issue",
			body:       models.UpdateEstimateBody{NewEstimate: 3},
			wantedCode: http.StatusBadRequest,
		},
		{
			name:       "fail_missing_estimate",
			body:       models.UpdateEstimateBody{IssueID: 1},
			wantedCode: http.StatusBadRequest,
		},
	}

	api := getTestAPI(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Cleanup(func() {
				api.IssueRepo.CustomQuery("UPDATE issues SET estimate=5, last_changed=? WHERE issue_id = 1", time.Date(2023, time.June, 28, 15, 15, 0, 0, time.Local))
			})

			w := httptest.NewRecorder()
			bodyMarshal, err := json.Marshal(test.body)
			if err != nil {
				t.Error(err)
			}
			reader := bytes.NewReader(bodyMarshal)
			req, _ := http.NewRequest("PATCH", "/issues/estimate", reader)
			api.Router.ServeHTTP(w, req)

			if w.Code != test.wantedCode {
				t.FailNow()
			}
		})
	}
}
