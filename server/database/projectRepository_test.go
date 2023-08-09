package database

import (
	"aigr20/prom/models"
	"aigr20/prom/testdata"
	"testing"
)

func getProjectRepository(t *testing.T) *ProjectRepository {
	db, err := CreateConnection("prom_test", "prom_tester", "tester")
	if err != nil {
		t.Error("Failed on database connection")
	}

	return NewProjectRepository(db)
}

func TestGetAllCount(t *testing.T) {
	repository := getProjectRepository(t)
	defer repository.Close()
	t.Run("GetAll/length", func(t *testing.T) {
		projects, err := repository.GetAll()
		if err != nil {
			t.Error(err)
		}
		if len(projects) != 3 {
			t.Fail()
		}
	})
}

func TestGetAllContent(t *testing.T) {
	tests := []struct {
		name string
		want models.Project
	}{
		{
			name: "project_1",
			want: testdata.SampleProjects[0],
		},
		{
			name: "project_2",
			want: testdata.SampleProjects[1],
		},
		{
			name: "project_3",
			want: testdata.SampleProjects[2],
		},
	}

	repository := getProjectRepository(t)
	defer repository.Close()
	projects, err := repository.GetAll()
	if err != nil {
		t.Error(err)
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if projects[i] != test.want {
				t.Fail()
			}
		})
	}
}

func TestGetOneContent(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want models.Project
	}{
		{
			name: "project_1",
			id:   1,
			want: testdata.SampleProjects[0],
		},
		{
			name: "project_2",
			id:   2,
			want: testdata.SampleProjects[1],
		},
		{
			name: "project_3",
			id:   3,
			want: testdata.SampleProjects[2],
		},
	}

	repository := getProjectRepository(t)
	defer repository.Close()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			project, err := repository.GetOne(test.id)
			if err != nil {
				t.Error(err)
			}
			if project != test.want {
				t.Fail()
			}
		})
	}
}

func TestCreateProject(t *testing.T) {
	tests := []struct {
		name        string
		body        models.ProjectCreateForm
		wantedId    int
		wantedName  string
		wantedError error
	}{
		{
			name:        "valid_creation",
			body:        models.ProjectCreateForm{Name: "testproject"},
			wantedId:    4,
			wantedName:  "testproject",
			wantedError: nil,
		},
		{
			name:        "missing_name",
			body:        models.ProjectCreateForm{},
			wantedId:    0,
			wantedName:  "",
			wantedError: ErrProjectCreate,
		},
	}
	repo := getProjectRepository(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Cleanup(func() {
				repo.CustomQuery("DELETE FROM projects WHERE project_id > 3")
				repo.CustomQuery("ALTER TABLE projects AUTO_INCREMENT = 4")
			})

			project, err := repo.CreateProject(test.body)
			if err != test.wantedError {
				t.Fail()
				return
			}
			if project.ID != test.wantedId || project.Name != test.wantedName {
				t.Fail()
			}
		})
	}
}
