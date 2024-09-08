package dbase

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewConnection(conn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
