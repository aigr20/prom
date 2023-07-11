package models

import (
	"database/sql"
	"log"
	"time"
)

type Project struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"createdAt"`
	Updated time.Time `json:"updatedAt"`
}

func ScanProjects(rows *sql.Rows) ([]Project, error) {
	projects := make([]Project, 0)
	for rows.Next() {
		var project Project
		err := rows.Scan(&project.ID, &project.Name, &project.Created, &project.Updated)
		if err != nil {
			log.Println(err)
			continue
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func ScanProject(row *sql.Row) (Project, error) {
	if err := row.Err(); err != nil {
		log.Println(err)
		return Project{}, err
	}
	var project Project
	err := row.Scan(&project.ID, &project.Name, &project.Created, &project.Updated)
	if err != nil {
		log.Println(err)
		return Project{}, err
	}
	return project, nil
}

type ProjectCreateForm struct {
	Name string `json:"name" binding:"required"`
}
