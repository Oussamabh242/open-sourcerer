package dbase

import (
	"database/sql"
	"log"
)

type PendingSub struct {
	Email string
}

// add the email to the db to be confirmed later
const INSERT_CONFIRM = `INSERT INTO confirmations(id , email) values($1 ,$2)`

func AddConfirm(db *sql.DB, email, id string) error {
	_, err := db.Exec(INSERT_CONFIRM, id, email)
	return err
}

const SUB_CONFIRM = `INSERT INTO subscribe(email) values($1);`

func ConfirmSub(db *sql.DB, email string) error {
	_, err := db.Exec(SUB_CONFIRM, email)
	return err
}

const AFTER_CONFIRM = `DELETE FROM confirmations where email = $1;`

func DeletePendingSub(db *sql.DB, email string) error {
	_, err := db.Exec(AFTER_CONFIRM, email)
	return err
}
const GET_EMAIL_FROM_PENDING = `SELECT email FROM confirmations where id = $1 ;  `

func GetEmail(db *sql.DB, id string) (PendingSub, error) {
	row := db.QueryRow(GET_EMAIL_FROM_PENDING, id)
	var pending PendingSub
	err := row.Scan(&pending.Email)
	if err == sql.ErrNoRows {
		log.Println(err)

		return pending, DB_NO_RECORD
	} else if err != nil {
		log.Println("db error", err)
		return pending, err
	}

	return pending, nil
}

const GET_SUBSCRIBERS = `SELECT email FROM subscribe ;  `

func GetSubscribers(db *sql.DB) ([]PendingSub, error) {
	rows, err := db.Query(GET_SUBSCRIBERS)
	if err != nil {
		return nil, err
	}
	var subscribers []PendingSub
	defer rows.Scan()
	for rows.Next() {
		var sub PendingSub
		if err := rows.Scan(&sub.Email); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, sub)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return subscribers, nil
}
