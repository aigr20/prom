package models

import (
	"database/sql"
	"log"
	"time"
)

type Issue struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Estimate    int        `json:"estimate"`
	Created     time.Time  `json:"createdAt"`
	Updated     time.Time  `json:"updatedAt"`
	ProjectID   int        `json:"-"`
	Status      string     `json:"status"`
	Tags        []IssueTag `json:"tags"`
}

// Compares issues without comparing the Updated field
func (issue *Issue) LenientEquals(other Issue) bool {
	tagsMatch := true
	if len(issue.Tags) == len(other.Tags) {
		for i, tag := range issue.Tags {
			if tag.Text != other.Tags[i].Text || tag.Color != other.Tags[i].Color {
				tagsMatch = false
				break
			}
		}
	} else {
		tagsMatch = false
	}

	return issue.ID == other.ID &&
		issue.Title == other.Title &&
		issue.Description == other.Description &&
		issue.Estimate == other.Estimate &&
		issue.Created == other.Created &&
		issue.ProjectID == other.ProjectID &&
		issue.Status == other.Status &&
		tagsMatch
}

func ScanIssues(rows *sql.Rows) ([]Issue, error) {
	issues := make([]Issue, 0)
	for rows.Next() {
		var issue Issue
		var tag IssueTag
		err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Estimate, &issue.Created, &issue.Updated, &issue.ProjectID, &issue.Status, &tag.ID, &tag.Text, &tag.Color)
		if err != nil {
			log.Println(err)
			continue
		}

		if scannedIssue, isScanned := issueHasBeenScanned(issues, issue.ID); !isScanned {
			issue.Tags = make([]IssueTag, 0)
			if tag.Text != "" && tag.Color != "" {
				issue.Tags = append(issue.Tags, tag)
			}
			issues = append(issues, issue)
		} else {
			if tag.Text != "" && tag.Color != "" {
				scannedIssue.Tags = append(scannedIssue.Tags, tag)
			}
		}
	}

	return issues, nil
}

func issueHasBeenScanned(issues []Issue, id int) (*Issue, bool) {
	for i := range issues {
		if issues[i].ID == id {
			return &issues[i], true
		}
	}
	return nil, false
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

type AddIssueTagsBody struct {
	IssueID int   `json:"issueId" binding:"required"`
	Tags    []int `json:"tags" binding:"required"`
}

type UpdateIssueBody struct {
	IssueID int                    `json:"issueId" binding:"required"`
	Updates map[string]interface{} `json:"updates"`
}
