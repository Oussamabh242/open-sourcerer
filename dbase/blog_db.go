package dbase

import (
	"github.com/Oussamabh242/open-sourcerer/views/blogview"

	"database/sql"
	"fmt"
	"log"
	"time"
)

var (
	DB_NO_RECORD = fmt.Errorf("No db Record")
)

type Blog struct {
	ID         int
	Content    string
	Title      string
	Created_at string
	View_count int
}

func formatDate(t time.Time) string {
	// Format the time.Time value to "02 Jan 2006" format
	return t.Format("02 Jan 2006")
}

const insertPost = `
INSERT INTO articles(title ,content , created_at)
VALUES($1 , $2 , $3);
`

func NewBlogPost(db *sql.DB, title, content string) error {
	t := time.Now()
	strTime := formatDate(t)
	_, err := db.Exec(insertPost, title, fmt.Sprintf("# %s\n\n> ###### %s\n\n%s", title, strTime, content), strTime)
	if err != nil {
		return err
	}
	return nil
}

const singlePost = `
SELECT title , content FROM articles 
WHERE id = $1 ;
`

func GetBlogById(db *sql.DB, id int) (Blog, error) {
	row := db.QueryRow(singlePost, id)
	var post Blog

	err := row.Scan(&post.Title, &post.Content)
	if err == sql.ErrNoRows {
		log.Println(err)

		return Blog{}, DB_NO_RECORD
	} else if err != nil {
		log.Println("db error", err)
		return Blog{}, err
	}
	post.ID = id
	return post, nil

}

const AllPosts = `SELECT id ,title , created_at, view_count from articles order by id desc; `

func GetAllPostsDB(db *sql.DB) ([]blogview.Overview, error) {
	rows, err := db.Query(AllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var allposts []blogview.Overview

	for rows.Next() {
		var post blogview.Overview
		if err := rows.Scan(&post.ID, &post.Title, &post.Created_at, &post.View_count); err != nil {
			return allposts, err
		}
		allposts = append(allposts, post)
	}

	if err := rows.Err(); err != nil {
		return allposts, nil
	}
	return allposts, nil
}

const NotifPost = `
SELECT id   FROM articles 
WHERE title =  $1 ;
`

func GetPostId(db *sql.DB, title string) (Blog, error) {
	row := db.QueryRow(NotifPost, title)
	var post Blog

	err := row.Scan(&post.ID)
	if err == sql.ErrNoRows {
		log.Println(err)

		return Blog{}, DB_NO_RECORD
	} else if err != nil {
		log.Println("db error", err)
		return Blog{}, err
	}
	post.Title = title
	return post, nil

}

const DELETE_POST = `
DELETE FROM articles where id = $1 ;
`

func DeletePost(db *sql.DB, id int) error {
	_, err := db.Exec(DELETE_POST, id)
	return err
}

const UPDATE_POST = `
UPDATE articles 
SET title = $1 , content = $2
WHERE id = $3
;
`

func UpdatePost(db *sql.DB, title string, content string, id int) error {
	_, err := db.Exec(UPDATE_POST, title, content, id)
	return err
}

const INCREMENT_VIEW_COUNT = `
UPDATE articles 
SET view_count = view_count+1
WHERE id = $1
;
`

func IncrementViews(db *sql.DB, id int) {
	_, err := db.Exec(INCREMENT_VIEW_COUNT, id)
	if err != nil {
		log.Println("error incrementing view_count , ", err)
	}
}
