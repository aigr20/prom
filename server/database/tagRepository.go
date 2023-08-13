package database

import (
	"aigr20/prom/models"
	"database/sql"
	"log"
)

type TagRepository struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{
		db: db,
	}
}

func (rep *TagRepository) FindTagByID(tagId int) (models.IssueTag, bool) {
	row := rep.db.QueryRow("SELECT tag_id, tag_text, tag_color FROM tags WHERE tag_id = ?", tagId)
	if row.Err() != nil {
		log.Println(row.Err())
		return models.IssueTag{}, false
	}

	var tag models.IssueTag
	err := row.Scan(&tag.ID, &tag.Text, &tag.Color)
	if err != nil {
		log.Println(err)
		return models.IssueTag{}, false
	}

	return tag, true
}

func (rep *TagRepository) FindTagInProject(tagId int, project int) (models.IssueTag, bool) {
	const query = `
	SELECT t.tag_id, t.tag_text, t.tag_color FROM tags AS t
	JOIN project_tags AS pts ON pts.tag_id = t.tag_id
	WHERE t.tag_id = ? && pts.project_id = ?`
	row := rep.db.QueryRow(query, tagId, project)
	if row.Err() != nil {
		log.Println(row.Err())
		return models.IssueTag{}, false
	}

	var tag models.IssueTag
	err := row.Scan(&tag.ID, &tag.Text, &tag.Color)
	if err != nil {
		log.Println(err)
		return models.IssueTag{}, false
	}
	return tag, true
}

func (rep *TagRepository) CreateTag(tag models.CreateTagBody) (int, error) {
	res, err := rep.db.Exec("INSERT INTO tags (tag_text, tag_color) VALUES (?, ?)", tag.Text, tag.Color)
	var id int64
	if err != nil {
		log.Println(err)
		row := rep.db.QueryRow("SELECT tag_id FROM tags WHERE tag_text = ? && tag_color = ?", tag.Text, tag.Color)
		if row.Err() != nil {
			log.Println(err)
			return -1, ErrTagCreate
		}
		err = row.Scan(&id)
		if err != nil {
			log.Println(err)
			return -1, ErrTagCreate
		}
	} else {
		id, _ = res.LastInsertId()
	}

	_, err = rep.db.Exec("INSERT INTO project_tags (tag_id, project_id) VALUES (?, ?)", id, tag.Project)
	if err != nil {
		log.Println(err)
		return -1, ErrTagDuplicate
	}

	return int(id), nil
}
