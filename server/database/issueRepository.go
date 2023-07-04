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
	rows, err := rep.db.Query("SELECT * FROM issues WHERE project = ?", projectId)
	if err != nil {
		log.Println(err)
		return nil, ErrIssuesNotFound
	}
	defer rows.Close()

	issues := make([]models.Issue, 0)
	for rows.Next() {
		var issue models.Issue
		err = rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Created, &issue.Updated, &issue.ProjectID)
		if err != nil {
			log.Println(err)
			continue
		}
		issues = append(issues, issue)
	}

	return issues, nil
}

func (rep *IssueRepository) GetOne(issueId int) (models.Issue, error) {
	row := rep.db.QueryRow("SELECT * FROM issues WHERE issue_id = ?", issueId)
	if row.Err() != nil {
		log.Println(row.Err())
		return models.Issue{}, ErrIssueNotFound
	}

	var issue models.Issue
	err := row.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Created, &issue.Updated, &issue.ProjectID)
	if err != nil {
		log.Println(err)
		return models.Issue{}, ErrIssueCreate
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
