package handlers

import (
	"database/sql"
	"fmt"
	"github.com/Oussamabh242/open-sourcerer/dbase"
	"github.com/Oussamabh242/open-sourcerer/views/blogview"
	"github.com/Oussamabh242/open-sourcerer/views/layout"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
	"time"
)

type BlogHandler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) BlogHandler {
	return BlogHandler{
		db: db,
	}
}

// Create Post Endpoint
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
	post, err := dbase.GetPostId(b.db, title)
	if err != nil {
		log.Println(post, err)
	} else {

		subs, err := dbase.GetSubscribers(b.db)
		if err != nil {
			log.Println(subs, err)
		} else {

			var strSubs []string
			for _, sub := range subs {
				strSubs = append(strSubs, sub.Email)
			}
			go SendNewPostAdded(strSubs, fmt.Sprintf("http://"+c.Request().Host+"/blog/%d", post.ID), post.Title)

		}

	}

	return c.HTML(201, "<h1> added successfully </h1>")
}

// PreviewPost
func PreviewPost(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("body")
	strTime := time.Now().Format("02 Jan 2006")
	post := fmt.Sprintf("# %s\n\n> ###### %s\n\n%s", title, strTime, content)
	return Render(c, 200, blogview.BlogPost("", post))
}

// ADD POST VIEW
func AddPost(c echo.Context) error {
	return Render(c, 200, blogview.AddPost())
}

// Retrieve Single Post
func (b BlogHandler) GetPost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		resp := "<h1> You Should Provide a number </h1>"
		return c.HTML(400, resp)
	}
	post, err := dbase.GetBlogById(b.db, id)
	if err == dbase.DB_NO_RECORD {
		return Render(c, 404, layout.NotFound())
	}

	return Render(c, 200, blogview.BlogPost(post.Title, post.Content))

}

// See All Posts
func (b BlogHandler) GetAllPosts(c echo.Context) error {
	posts, err := dbase.GetAllPostsDB(b.db)
	if err != nil {
		log.Println(err)
		resp := "<h1>Internal Server error </h1>"
		return c.HTML(500, resp)
	}

	return Render(c, 200, blogview.All(posts))

}
