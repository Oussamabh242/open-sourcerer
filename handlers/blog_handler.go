package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Oussamabh242/open-sourcerer/dbase"
	"github.com/labstack/echo/v4"
)

type Blog struct {
	ID         int
	Content    string
	Created_at string
	View_count int
}

type BlogHandler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) BlogHandler {
	return BlogHandler{
		db: db,
	}
}

func (b BlogHandler) CreatePostHandler(c echo.Context) error {
	GetBlogById(b.db, 1)
	return c.HTML(418, "<h1> Too short </h1>")

	post := c.FormValue("body")
	if len(post) < 100 {
		return c.HTML(418, "<h1> Too short </h1>")
	}
	err := dbase.NewBlogPost(b.db, post)
	if err != nil {
		log.Fatal(fmt.Sprintf("error while inserting new article to the db %s", err))
		return c.HTML(500, "<h1> error when inserting </h1>")
	}
	return c.HTML(201, "<h1> added successfully </h1>")
}

const singlePost = `
SELECT * FROM articles 
WHERE id = ? ;
`

func GetBlogById(db *sql.DB, id int) {
	row := db.QueryRow(singlePost, id)
	var post Blog

	err := row.Scan(&post.ID, &post.Content, &post.Created_at, &post.View_count)
	if err != nil {
		fmt.Printf("error : %s", err)
	}
	fmt.Println(post)

}
