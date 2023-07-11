package database

import (
	"aigr20/prom/models"
	"database/sql"
	"log"
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

	projects, err := models.ScanProjects(rows)
	if err != nil {
		return nil, ErrProjectScan
	}

	return projects, nil
}

func (rep *ProjectRepository) GetOne(id int) (models.Project, error) {
	row := rep.db.QueryRow("SELECT * FROM projects WHERE project_id = ?", id)
	if row == nil {
		return models.Project{}, ErrProjectNotFound
	}

	project, err := models.ScanProject(row)
	if err != nil {
		return models.Project{}, ErrProjectScan
	}
	return project, nil
}

func (rep *ProjectRepository) CreateProject(body models.ProjectCreateForm) (models.Project, error) {
	if body.Name == "" {
		return models.Project{}, ErrProjectCreate
	}

	result, err := rep.db.Exec("INSERT INTO projects (project_name) VALUES (?)", body.Name)
	if err != nil {
		log.Println(err)
		return models.Project{}, ErrProjectCreate
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return models.Project{}, ErrProjectCreate
	}

	project, err := rep.GetOne(int(lastId))
	if err != nil {
		log.Println(err)
		return models.Project{}, ErrProjectCreate
	}

	return project, nil
}

// Should only be used in tests
func (rep *ProjectRepository) CustomQuery(query string, args ...any) (sql.Result, error) {
	return rep.db.Exec(query, args...)
}
