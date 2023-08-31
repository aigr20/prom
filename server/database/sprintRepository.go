package database

import (
	"aigr20/prom/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const baseSprintQuery = `
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
`

type SprintRepository struct {
	db *sql.DB
}

func NewSprintRepository(db *sql.DB) *SprintRepository {
	return &SprintRepository{
		db: db,
	}
}

func (rep *SprintRepository) GetSprint(sprintId int) (models.Sprint, error) {
	const query = "SELECT sprint_id, sprint_name, sprint_start, sprint_end, finished, current FROM sprints WHERE sprint_id = ?"
	rows, err := rep.db.Query(query, sprintId)
	if err != nil {
		log.Println(err)
		return models.Sprint{}, ErrSprintNotFound
	}

	scanResult := models.SimpleSprintScan(rows)
	if len(scanResult) != 1 {
		log.Println(err)
		return models.Sprint{}, ErrSprintNotFound
	}

	return scanResult[0], nil
}

func (rep *SprintRepository) GetSprintWithIssues(sprintId int) (models.Sprint, error) {
	query := fmt.Sprintf("%s WHERE sprint_id = ?", baseSprintQuery)
	rows, err := rep.db.Query(query, sprintId)
	if err != nil {
		log.Println(err)
		return models.Sprint{}, ErrSprintNotFound
	}

	sprint := models.ScanSprint(rows)
	return sprint, nil
}

func (rep *SprintRepository) GetCurrentSprintForProject(projectId int) (models.Sprint, error) {
	query := fmt.Sprintf("%s WHERE sprint_project = ? AND current = TRUE", baseSprintQuery)
	rows, err := rep.db.Query(query, projectId)
	if err != nil {
		log.Println(err)
		return models.Sprint{}, ErrSprintNotFound
	}

	sprint := models.ScanSprint(rows)
	if sprint.Equals(models.Sprint{}) {
		return models.Sprint{}, ErrSprintNotFound
	}

	return sprint, nil
}

func (rep *SprintRepository) CreateSprint(body models.CreateSprintBody) (int, error) {
	transaction, err := rep.db.Begin()
	if err != nil {
		log.Println(err)
		return 0, ErrSprintCreate
	}

	if body.Start.IsZero() {
		body.Start = time.Now()
	}
	if body.End.IsZero() {
		body.End = body.Start.AddDate(0, 0, 7)
	}

	_, err = transaction.Exec("UPDATE sprints SET current = FALSE WHERE project = ?", body.Project)
	if err != nil {
		log.Println(err)
		transaction.Rollback()
		return 0, ErrSprintCreate
	}

	const query = "INSERT INTO sprints (sprint_name, project, sprint_start, sprint_end) VALUES (?, ?, ?, ?)"
	result, err := transaction.Exec(query, body.Name, body.Project, body.Start, body.End)
	if err != nil {
		log.Println(err)
		transaction.Rollback()
		return 0, ErrSprintCreate
	}

	inserted, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, ErrSprintCreate
	}

	err = transaction.Commit()
	if err != nil {
		log.Println(err)
		return 0, ErrSprintCreate
	}

	return int(inserted), nil
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
