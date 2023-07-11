package database

import (
	"database/sql"
	"log"
)

type StatusRepository struct {
	db *sql.DB
}

func NewStatusRepository(db *sql.DB) *StatusRepository {
	return &StatusRepository{
		db: db,
	}
}

func (rep *StatusRepository) GetIDByName(name string) (int, error) {
	row := rep.db.QueryRow("SELECT status_id FROM issue_statuses WHERE status_text = ?", name)
	if row.Err() != nil {
		log.Println(row.Err())
		return -1, ErrStatusNotFound
	}

	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Println(err)
		return -1, ErrStatusNotFound
	}
	return id, nil
}
