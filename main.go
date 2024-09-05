package main

import (
	"log"
	"net/http"

	"github.com/Oussamabh242/open-sourcerer/dbase"
	"github.com/Oussamabh242/open-sourcerer/handlers"
	"github.com/Oussamabh242/open-sourcerer/views/blogview"
	"github.com/Oussamabh242/open-sourcerer/views/home"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

const (
	OK = http.StatusOK
)

func main() {
	db, err := dbase.NewConnection("dev.db")
	if err != nil {
		log.Fatal("error happend while connecting to the data base")
	}

	blogHandler := handlers.NewHandler(db)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return Render(c, http.StatusOK, home.Index("oussama", "21"))
	})
	e.GET("/blog", func(c echo.Context) error {
		return Render(c, http.StatusOK, blogview.Main())
	})

	e.POST("/add", blogHandler.CreatePostHandler)

	e.Logger.Fatal(e.Start(":3000"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
