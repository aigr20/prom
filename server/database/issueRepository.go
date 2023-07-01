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
