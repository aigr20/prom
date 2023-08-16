package models

import (
	"database/sql"
	"log"
)

type IssueTag struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Color string `json:"color"`
}

func ScanTags(rows *sql.Rows) ([]IssueTag, error) {
	tags := make([]IssueTag, 0)
	for rows.Next() {
		var tag IssueTag
		err := rows.Scan(&tag.ID, &tag.Text, &tag.Color)
		if err != nil {
			log.Println(err)
			continue
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

type CreateTagBody struct {
	Project int    `json:"projectId" binding:"required"`
	Text    string `json:"text" binding:"required"`
	Color   string `json:"color" binding:"required"`
}
