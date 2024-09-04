package main

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
  "github.com/Oussamabh242/open-sourcerer/views"

)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		data := map[string]interface{}{
			"Name":  "oussama",
			"Title": "Home",
			"Age":   21,
		}
		return c.Render(http.StatusOK, "index.html", data)
	})
	e.GET("/sup", func(c echo.Context) error {

		return Render(c, http.StatusAccepted, views.Home())
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
