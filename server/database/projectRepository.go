package database

import (
	"aigr20/prom/models"
	"database/sql"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (rep *ProjectRepository) Close() error {
	return rep.db.Close()
}

func (rep *ProjectRepository) GetAll() ([]models.Project, error) {
	rows, err := rep.db.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := make([]models.Project, 0)
	for rows.Next() {
		var project models.Project
		err = rows.Scan(&project.ID, &project.Name, &project.Created, &project.Updated)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}
