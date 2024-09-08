package main

import (
	"log"
	"net/http"

	"github.com/Oussamabh242/open-sourcerer/dbase"
	"github.com/Oussamabh242/open-sourcerer/handlers"
	"golang.org/x/time/rate"

	// "github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	newsHandler := handlers.NewNewsHandler(db)
	dashboardHandler := handlers.NewDashHandler(db)
	e := echo.New()
  e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(10))))

	e.Static("/static", "static")

	e.GET("/", handlers.MainHandler)
	e.GET("/blog/add", handlers.Middleware(handlers.AddPost, handlers.MainHandler))
	e.POST("/add", handlers.Middleware(blogHandler.CreatePostHandler, handlers.MainHandler))
	e.GET("/blog", blogHandler.GetAllPosts)
	e.GET("/blog/:id", blogHandler.GetPost)
	e.POST("/preview", handlers.PreviewPost)
	e.GET("/login", handlers.Login)
	e.POST("/signin", handlers.Signin)
	e.POST("/subscribe", newsHandler.Subscribe)
	e.GET("/subscription/confirm/:sid", newsHandler.Confirm)
	e.GET("/admin/dashboard", handlers.Middleware(dashboardHandler.Dashboard, handlers.MainHandler))
	e.DELETE("/admin/delete/:id", handlers.Middleware(dashboardHandler.DeletePost, handlers.MainHandler))
	e.GET("/admin/view/update/:id", handlers.Middleware(dashboardHandler.UpdateView, handlers.MainHandler))
	e.PUT("/admin/update/:id", handlers.Middleware(dashboardHandler.Update, handlers.MainHandler))
	e.Logger.Fatal(e.Start(":3000"))
}
