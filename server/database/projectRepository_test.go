package database

import (
	"aigr20/prom/models"
	"testing"
	"time"
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
			want: models.Project{
				ID:      1,
				Name:    "Glada Schemat",
				Created: time.Date(2023, time.June, 24, 10, 0, 0, 0, time.Local),
				Updated: time.Date(2023, time.June, 24, 10, 15, 0, 0, time.Local),
			},
		},
		{
			name: "project_2",
			want: models.Project{
				ID:      2,
				Name:    "Max Power",
				Created: time.Date(2023, time.June, 24, 9, 10, 0, 0, time.Local),
				Updated: time.Date(2023, time.June, 24, 16, 13, 12, 0, time.Local),
			},
		},
		{
			name: "project_3",
			want: models.Project{
				ID:      3,
				Name:    "Mina Recept",
				Created: time.Date(2023, time.June, 25, 15, 22, 35, 0, time.Local),
				Updated: time.Date(2023, time.June, 26, 12, 0, 14, 0, time.Local),
			},
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
