package dbase

import (
	"database/sql"
)

const insertPost = `
INSERT INTO articles(content)
VALUES(?);
`

func NewBlogPost(db *sql.DB, post string) error {
	_, err := db.Exec(insertPost, post)
	if err != nil {
		return err
	}
	return nil
}
