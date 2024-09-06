package dbase

import (
	"database/sql"
	"fmt"
	"time"
)

const insertPost = `
INSERT INTO articles(title ,content , created_at)
VALUES(? , ? , ?);
`

func formatDate(t time.Time) string {
	// Format the time.Time value to "02 Jan 2006" format
	return t.Format("02 Jan 2006")
}

func NewBlogPost(db *sql.DB, title, content string) error {
	t := time.Now()
	strTime := formatDate(t)
	_, err := db.Exec(insertPost, title, fmt.Sprintf("# %s\n\n> ###### %s\n\n%s", title, strTime, content), strTime)
	if err != nil {
		return err
	}
	return nil
}
