package models

import (
	"database/sql"
	"log"
	"time"
)

type Issue struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Created     time.Time `json:"createdAt"`
	Updated     time.Time `json:"updatedAt"`
	ProjectID   int       `json:"-"`
	Status      string    `json:"status"`
}

func ScanIssues(rows *sql.Rows) ([]Issue, error) {
	issues := make([]Issue, 0)
	for rows.Next() {
		var issue Issue
		err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Created, &issue.Updated, &issue.ProjectID, &issue.Status)
		if err != nil {
			log.Println(err)
			continue
		}
		issues = append(issues, issue)
	}
	return issues, nil
}

func ScanIssue(row *sql.Row) (Issue, error) {
	if err := row.Err(); err != nil {
		return Issue{}, err
	}
	var issue Issue
	err := row.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Created, &issue.Updated, &issue.ProjectID, &issue.Status)
	if err != nil {
		return Issue{}, err
	}
	return issue, nil
}

type IssueCreateForm struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	ProjectID   int    `json:"project" binding:"required"`
}
