package main

import (
	"github.com/Oussamabh242/open-sourcerer/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return Render(c, http.StatusOK, views.Index("oussama", "21"))
	})

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
