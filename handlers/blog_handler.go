package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Oussamabh242/open-sourcerer/dbase"
	"github.com/Oussamabh242/open-sourcerer/views/blogview"
	"github.com/labstack/echo/v4"
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

type BlogHandler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) BlogHandler {
	return BlogHandler{
		db: db,
	}
}

func (b BlogHandler) CreatePostHandler(c echo.Context) error {

	content := c.FormValue("body")
	title := c.FormValue("title")
	if len(content) < 100 {
		return c.HTML(418, "<h1> Too short </h1>")
	}
	err := dbase.NewBlogPost(b.db, title, content)
	if err != nil {
		log.Fatal(fmt.Sprintf("error while inserting new article to the db %s", err))
		return c.HTML(500, "<h1> error when inserting </h1>")
	}
	return c.HTML(201, "<h1> added successfully </h1>")
}

func PreviewPost(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("body")
	strTime := time.Now().Format("02 Jan 2006")
	post := fmt.Sprintf("# %s\n\n> ###### %s\n\n%s", title, strTime, content)
	return Render(c, 200, blogview.BlogPost(post))
}

func (b BlogHandler) GetPost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
		resp := "<h1> You Should Provide a number </h1>"
		return c.HTML(400, resp)
	}
	post, err := GetBlogById(b.db, id)
	if err == DB_NO_RECORD {
		resp := "<h1> 404 Not found </h1>"
		return c.HTML(404, resp)
	}

	return Render(c, 200, blogview.BlogPost(post.Content))

}

func (b BlogHandler) GetAllPosts(c echo.Context) error {
	posts, err := GetAllPostsDB(b.db)
	if err != nil {
		fmt.Println(err)
		resp := "<h1>Internal Server error </h1>"
		return c.HTML(500, resp)
	}

	return Render(c, 200, blogview.All(posts))

}

const singlePost = `
SELECT content FROM articles 
WHERE id = ? ;
`

func GetBlogById(db *sql.DB, id int) (Blog, error) {
	row := db.QueryRow(singlePost, id)
	var post Blog

	err := row.Scan(&post.Content)
	if err == sql.ErrNoRows {
		fmt.Println(err)

		return Blog{}, DB_NO_RECORD
	} else if err != nil {
		fmt.Println("db error", err)
		return Blog{}, err
	}
	return post, nil

}

const AllPosts = `SELECT id ,title , created_at from articles order by id desc; `

func GetAllPostsDB(db *sql.DB) ([]blogview.Overview, error) {
	rows, err := db.Query(AllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var allposts []blogview.Overview

	for rows.Next() {
		var post blogview.Overview
		if err := rows.Scan(&post.ID, &post.Title, &post.Created_at); err != nil {
			return allposts, err
		}
		allposts = append(allposts, post)
	}

	if err := rows.Err(); err != nil {
		return allposts, nil
	}
	return allposts, nil
}

// GetBlogById(b.db, 1)
// 	return c.HTML(418, "<h1> Too short </h1>")
