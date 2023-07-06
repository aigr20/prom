package database

import (
	"aigr20/prom/models"
	"database/sql"
	"log"
)

type IssueRepository struct {
	db *sql.DB
}

func NewIssueRepository(db *sql.DB) *IssueRepository {
	return &IssueRepository{
		db: db,
	}
}

func (rep *IssueRepository) Close() error {
	return rep.db.Close()
}

func (rep *IssueRepository) GetIssuesFromProject(projectId int) ([]models.Issue, error) {
	const query = `
	SELECT
		i.issue_id,
		i.issue_title,
		i.issue_description,
		i.creation_date,
		i.last_changed,
		i.project,
		s.status_text
	FROM issues AS i
	JOIN issue_statuses AS s
	ON i.issue_status = s.status_id
	WHERE project = ?`
	rows, err := rep.db.Query(query, projectId)
	if err != nil {
		log.Println(err)
		return nil, ErrIssuesNotFound
	}
	defer rows.Close()

	issues, err := models.ScanIssues(rows)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return issues, nil
}

func (rep *IssueRepository) GetOne(issueId int) (models.Issue, error) {
	const query = `
	SELECT
		i.issue_id,
		i.issue_title,
		i.issue_description,
		i.creation_date,
		i.last_changed,
		i.project,
		s.status_text
	FROM issues AS i
	JOIN issue_statuses AS s
	ON i.issue_status = s.status_id
	WHERE i.issue_id = ?`
	row := rep.db.QueryRow(query, issueId)
	if row.Err() != nil {
		log.Println(row.Err())
		return models.Issue{}, ErrIssueNotFound
	}

	issue, err := models.ScanIssue(row)
	if err != nil {
		log.Println(err)
		return models.Issue{}, ErrIssueNotFound
	}

	return issue, nil
}

func (rep *IssueRepository) CreateIssue(body models.IssueCreateForm) (models.Issue, error) {
	if body.Title == "" || body.ProjectID == 0 {
		return models.Issue{}, ErrIssueCreate
	}

	result, err := rep.db.Exec("INSERT INTO issues (issue_title, issue_description, project) VALUES (?, ?, ?)", body.Title, body.Description, body.ProjectID)
	if err != nil {
		log.Println(err)
		return models.Issue{}, ErrIssueCreate
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return models.Issue{}, ErrIssueCreate
	}
	createdIssue, err := rep.GetOne(int(lastInsertId))
	if err != nil {
		return models.Issue{}, err
	}

	return createdIssue, nil
}

// Should only be used in tests
func (rep *IssueRepository) CustomQuery(query string, args ...any) (sql.Result, error) {
	return rep.db.Exec(query, args...)
}
