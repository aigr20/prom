package models

import (
	"database/sql"
	"log"
	"time"
)

type Issue struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Estimate    int       `json:"estimate"`
	Created     time.Time `json:"createdAt"`
	Updated     time.Time `json:"updatedAt"`
	ProjectID   int       `json:"-"`
	Status      string    `json:"status"`
}

// Compares issues without comparing the Updated field
func (issue *Issue) LenientEquals(other Issue) bool {
	return issue.ID == other.ID &&
		issue.Title == other.Title &&
		issue.Description == other.Description &&
		issue.Estimate == other.Estimate &&
		issue.Created == other.Created &&
		issue.ProjectID == other.ProjectID &&
		issue.Status == other.Status
}

func ScanIssues(rows *sql.Rows) ([]Issue, error) {
	issues := make([]Issue, 0)
	for rows.Next() {
		var issue Issue
		err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Estimate, &issue.Created, &issue.Updated, &issue.ProjectID, &issue.Status)
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
	err := row.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Estimate, &issue.Created, &issue.Updated, &issue.ProjectID, &issue.Status)
	if err != nil {
		return Issue{}, err
	}
	return issue, nil
}

type IssueCreateForm struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	Estimate    int    `json:"estimate,omitempty"`
	ProjectID   int    `json:"project" binding:"required"`
}

type UpdateIssueStatusBody struct {
	IssueID   int    `json:"issueId" binding:"required"`
	NewStatus string `json:"newStatus" binding:"required"`
}

type UpdateEstimateBody struct {
	IssueID     int `json:"issueId" binding:"required"`
	NewEstimate int `json:"newEstimate" binding:"required"`
}
