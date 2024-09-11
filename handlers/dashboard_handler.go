package handlers

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/Oussamabh242/open-sourcerer/dbase"
	"github.com/Oussamabh242/open-sourcerer/views/dashboard"
	"github.com/Oussamabh242/open-sourcerer/views/layout"

	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	db *sql.DB
}

func NewDashHandler(db *sql.DB) DashboardHandler {
	return DashboardHandler{
		db: db,
	}
}

func (dh DashboardHandler) Dashboard(c echo.Context) error {
	posts, err := dbase.GetAllPostsDB(dh.db)
	if err != nil {
		log.Println(err)
		return c.HTML(500, "error")
	}
	return Render(c, 200, dashboard.Dashbord(posts))
}

func (dh DashboardHandler) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.HTML(500, "you should provide an number")
	}
	err = dbase.DeletePost(dh.db, id)
	if err != nil {
		log.Println(err)
		return c.HTML(500, "you should provide an number")
	}
	c.Response().Header().Set("HX-Redirect", "/admin/dashboard")
	return c.NoContent(200)

}

func (dh DashboardHandler) UpdateView(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.HTML(500, "you should provide an number")
	}
	post, err := dbase.GetBlogById(dh.db, id)
	dashPost := dashboard.Blog{
		Title:      post.Title,
		View_count: post.View_count,
		ID:         post.ID,
		Content:    post.Content,
		Created_at: post.Created_at,
	}
	if err != nil {
		return Render(c, 404, layout.NotFound())
	}
	c.Response().Header().Set("HX-Redirect", "/admin/view/update/"+c.Param("id"))
	return Render(c, 200, dashboard.Update(dashPost))
}

func (dh DashboardHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.HTML(500, "you should provide an number")
	}
	err = dbase.UpdatePost(dh.db, c.FormValue("title"), c.FormValue("content"), id)
	if err != nil {
		log.Println(err)
		return c.HTML(500, "Error updating the post")
	}
	c.Response().Header().Set("HX-Redirect", "/admin/dashboard")
	return c.NoContent(200)
}
