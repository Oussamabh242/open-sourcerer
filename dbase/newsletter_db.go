package dbase

import (
	"database/sql"
)

const INSERT_CONFIRM = `INSERT INTO confirmations(id , email) values(? ,?)`

func AddConfirm(db *sql.DB, email, id string) error {
	_, err := db.Exec(INSERT_CONFIRM, id, email)
	return err
}
