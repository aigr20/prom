package api

import (
	"aigr20/prom/database"
	"aigr20/prom/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestAPI(t *testing.T) *API {
	db, err := database.CreateConnection("prom_test", "prom_tester", "tester")
	if err != nil {
		t.Error("Failed on database connection")
	}

	return NewAPI(db)
}

func projectsFromBody(t *testing.T, body []byte) map[string][]models.Project {
	projects := make(map[string][]models.Project, 0)
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&projects)
	if err != nil {
		t.Fail()
		return nil
	}

	return projects
}

func TestProjectsRoute(t *testing.T) {
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
