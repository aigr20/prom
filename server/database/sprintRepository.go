package database

import (
	"aigr20/prom/models"
	"database/sql"
	"log"
)

type SprintRepository struct {
	db *sql.DB
}

func NewSprintRepository(db *sql.DB) *SprintRepository {
	return &SprintRepository{
		db: db,
	}
}

func (rep *SprintRepository) GetSprint(sprintId int) (models.Sprint, error) {
	const query = `
	SELECT
		sprint_id,
		sprint_name,
		sprint_start,
		sprint_end,
		finished,
		current,
		issue_id,
		issue_title,
		issue_description,
		estimate,
		creation_date,
		last_changed,
		project,
		status_text,
		tag_id,
		tag_text,
		tag_color
	FROM sprint_issues_v
	WHERE sprint_id = ?`
	rows, err := rep.db.Query(query, sprintId)
	if err != nil {
		log.Println(err)
		return models.Sprint{}, ErrSprintNotFound
	}

	sprint := models.ScanSprint(rows)
	return sprint, nil
}

func (rep *SprintRepository) GetIssues(sprintId int) ([]models.Issue, error) {
	const query = `
	SELECT
		issue_id,
		issue_title,
		issue_description,
		estimate,
		creation_date,
		last_changed,
		project,
		status_text,
		tag_id,
		tag_text,
		tag_color
	FROM sprint_issues_v
	WHERE sprint_id = ?
	`
	rows, err := rep.db.Query(query, sprintId)
	if err != nil {
		log.Println(err)
		return []models.Issue{}, ErrIssuesNotFound
	}

	issues := models.ScanIssues(rows)
	return issues, nil
}
