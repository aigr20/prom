package models

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

var ErrColumnCount = errors.New("unexpected column count when scanning")

type Project struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Tags    []IssueTag `json:"tags"`
	Created time.Time  `json:"createdAt"`
	Updated time.Time  `json:"updatedAt"`
}

func (project *Project) Equals(other Project) bool {
	tagsMatch := true
	if len(project.Tags) == len(other.Tags) {
		for i := range project.Tags {
			if project.Tags[i] != other.Tags[i] {
				tagsMatch = false
				break
			}
		}
	} else {
		tagsMatch = false
	}

	return project.ID == other.ID &&
		project.Name == other.Name &&
		project.Created == other.Created &&
		tagsMatch
}

func ScanProjects(rows *sql.Rows) ([]Project, error) {
	if cols, _ := rows.Columns(); len(cols) > 4 {
		return scanProjectsWithTags(rows)
	} else if len(cols) == 4 {
		return simpleProjectScan(rows)
	}

	return []Project{}, ErrColumnCount
}

func simpleProjectScan(rows *sql.Rows) ([]Project, error) {
	projects := make([]Project, 0)
	for rows.Next() {
		var project Project
		project.Tags = []IssueTag{}
		err := rows.Scan(&project.ID, &project.Name, &project.Created, &project.Updated)
		if err != nil {
			log.Println(err)
			continue
		}

		projects = append(projects, project)
	}
	return projects, nil
}

func scanProjectsWithTags(rows *sql.Rows) ([]Project, error) {
	projects := make([]Project, 0)
	for rows.Next() {
		var project Project
		var tag IssueTag
		err := rows.Scan(&project.ID, &project.Name, &project.Created, &project.Updated, &tag.ID, &tag.Text, &tag.Color)
		if err != nil {
			log.Println(err)
			continue
		}

		if scannedProject, hasScanned := projectHasBeenScanned(projects, project.ID); !hasScanned {
			project.Tags = make([]IssueTag, 0)
			if tag.Text != "" && tag.Color != "" {
				project.Tags = append(project.Tags, tag)
			}
			projects = append(projects, project)
		} else {
			if tag.Text != "" && tag.Color != "" {
				scannedProject.Tags = append(scannedProject.Tags, tag)
			}
		}
	}
	return projects, nil
}

func projectHasBeenScanned(projects []Project, id int) (*Project, bool) {
	for i := range projects {
		if projects[i].ID == id {
			return &projects[i], true
		}
	}
	return nil, false
}

type ProjectCreateForm struct {
	Name string `json:"name" binding:"required"`
}
