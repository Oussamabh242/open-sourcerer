package handlers

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/Oussamabh242/open-sourcerer/dbase"
)

type NewsHandler struct {
	db *sql.DB
}

func NewNewsHandler(db *sql.DB) NewsHandler {

	return NewsHandler{
		db: db,
	}
}

func (n NewsHandler) Subscribe(c echo.Context) error {
	host := c.Request().Host
	email := c.FormValue("email")
	id := uuid.New().String()
	dbase.AddConfirm(n.db, email, id)
	go SendConfirmSubscription(email, "https://"+host+"/"+id)

	return c.String(200, "Verification email sent! Please check your inbox or spam folder to confirm.")

}
