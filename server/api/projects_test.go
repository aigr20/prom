package api

import (
	"aigr20/prom/database"
	"aigr20/prom/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func getTestAPI(t *testing.T) *API {
	db, err := database.CreateConnection("prom_test", "prom_tester", "tester")
	if err != nil {
		t.Error("Failed on database connection")
	}

	gin.DefaultWriter, _ = os.Open(os.DevNull)
	gin.DefaultErrorWriter, _ = os.Open(os.DevNull)
	return NewAPI(db)
}

func projectsFromBody(t *testing.T, body []byte) map[string][]models.Project {
	projects := make(map[string][]models.Project)
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&projects)
	if err != nil {
		t.Fail()
		return nil
	}

	return projects
}

func projectFromBody(t *testing.T, body []byte) models.Project {
	data := make(map[string]models.Project)
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&data)
	if err != nil {
		t.Fail()
		return models.Project{}
	}

	return data["data"]
}

func issuesFromBody(t *testing.T, body []byte) []models.Issue {
	data := make(map[string][]models.Issue)
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&data)
	if err != nil {
		t.Fail()
		return []models.Issue{}
	}

	return data["data"]
}

func TestAllProjectsRoute(t *testing.T) {
	api := getTestAPI(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/projects/all", nil)
	api.Router.ServeHTTP(w, req)

	t.Run("statuscode", func(t *testing.T) {
		if w.Code != 200 {
			t.Fail()
		}
	})
	t.Run("content", func(t *testing.T) {
		projects := projectsFromBody(t, w.Body.Bytes())
		if len(projects["data"]) != 3 {
			t.Fail()
		}
	})
}

func TestOneProjectRoute(t *testing.T) {
	api := getTestAPI(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/projects/2", nil)
	api.Router.ServeHTTP(w, req)

	t.Run("statuscode", func(t *testing.T) {
		if w.Code != 200 {
			t.Fail()
		}
	})
	t.Run("content", func(t *testing.T) {
		project := projectFromBody(t, w.Body.Bytes())
		if project.ID != 2 {
			t.Fail()
		}
	})
}

func TestProjectIssuesRoute(t *testing.T) {
	api := getTestAPI(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/projects/1/issues", nil)
	api.Router.ServeHTTP(w, req)

	t.Run("statuscode", func(t *testing.T) {
		if w.Code != 200 {
			t.Fail()
		}
	})
	t.Run("content", func(t *testing.T) {
		issues := issuesFromBody(t, w.Body.Bytes())
		if len(issues) != 2 {
			t.Fail()
		}
		for i, issue := range issues {
			if issue.ID != i+1 {
				t.Fail()
			}
		}
	})
}

func TestCreateProjectRoute(t *testing.T) {
	api := getTestAPI(t)

	tests := []struct {
		name       string
		body       models.ProjectCreateForm
		wantedCode int
		wantedName string
	}{
		{
			name:       "valid request",
			body:       models.ProjectCreateForm{Name: "test project"},
			wantedCode: http.StatusCreated,
			wantedName: "test project",
		},
		{
			name:       "invalid request",
			body:       models.ProjectCreateForm{},
			wantedCode: http.StatusBadRequest,
			wantedName: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Cleanup(func() {
				api.ProjectRepo.CustomQuery("DELETE FROM projects WHERE project_id > 3")
				api.ProjectRepo.CustomQuery("ALTER TABLE projects AUTO_INCREMENT = 4")
			})

			w := httptest.NewRecorder()
			bodyMarshal, err := json.Marshal(test.body)
			if err != nil {
				t.Error(err)
			}
			reader := bytes.NewReader(bodyMarshal)
			req, _ := http.NewRequest("POST", "/projects/create", reader)
			api.Router.ServeHTTP(w, req)

			if w.Code != test.wantedCode {
				t.FailNow()
			}
			if w.Code == http.StatusCreated {
				project := projectFromBody(t, w.Body.Bytes())
				if project.Name != test.wantedName {
					t.FailNow()
				}
			}
		})
	}
}
