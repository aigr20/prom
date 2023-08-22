package models

import (
	"database/sql"
	"log"
	"time"
)

type Sprint struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Start    time.Time `json:"start_date"`
	End      time.Time `json:"end_date"`
	Issues   []Issue   `json:"issues"`
	Finished bool      `json:"-"`
	Current  bool      `json:"-"`
}

func (sprint Sprint) Equals(other Sprint) bool {
	issuesEqual := true
	if len(sprint.Issues) == len(other.Issues) {
		for i := range other.Issues {
			if !sprint.Issues[i].Equals(other.Issues[i]) {
				issuesEqual = false
				break
			}
		}
	} else {
		issuesEqual = false
	}

	return issuesEqual
}

func ScanSprint(rows *sql.Rows) Sprint {
	var sprint Sprint
	sprint.Issues = make([]Issue, 0)

	for rows.Next() {
		var issue Issue
		var tag IssueTag

		err := rows.Scan(
			&sprint.ID,
			&sprint.Name,
			&sprint.Start,
			&sprint.End,
			&sprint.Finished,
			&sprint.Current,
			&issue.ID,
			&issue.Title,
			&issue.Description,
			&issue.Estimate,
			&issue.Created,
			&issue.Updated,
			&issue.ProjectID,
			&issue.Status,
			&tag.ID,
			&tag.Text,
			&tag.Color,
		)
		if err != nil {
			log.Println(err)
			continue
		}

		if scannedIssue, scanned := issueHasBeenScanned(sprint.Issues, issue.ID); !scanned {
			issue.Tags = make([]IssueTag, 0)
			if tag.Text != "" && tag.Color != "" {
				issue.Tags = append(issue.Tags, tag)
			}
			sprint.Issues = append(sprint.Issues, issue)
		} else if scannedIssue != nil && scanned {
			if tag.Text != "" && tag.Color != "" {
				scannedIssue.Tags = append(scannedIssue.Tags, tag)
			}
		}
	}

	return sprint
}
