package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection(dbName string, dbUser string, dbPassword string) (*sql.DB, error) {
	if dbUser == "" || dbPassword == "" {
		log.Printf("Database user or password unset. Received user: %s & password: %s. Set them using the --user and --password flags.\n", dbUser, dbPassword)
		return nil, ErrConnectionFailed
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?parseTime=true&loc=%s", dbUser, dbPassword, dbName, time.Local))
	if err != nil {
		log.Printf("Connecting to database failed with error: %s\n", err.Error())
		return nil, ErrConnectionFailed
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		log.Printf("Could not ping the database: %s\n", err.Error())
		return nil, ErrConnectionFailed
	}

	return db, nil
}
