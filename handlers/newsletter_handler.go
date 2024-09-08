package handlers

import (
	"database/sql"
	"fmt"
	"log"

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
	go SendConfirmSubscription(email, "http://"+host+"/subscription/confirm/"+id)

	return c.String(200, "Verification email sent! Please check your inbox or spam folder to confirm.")

}

func (n NewsHandler) Confirm(c echo.Context) error {
	sid := c.Param("sid")
	user, err := dbase.GetEmail(n.db, sid)
	if err != nil {
		log.Println(err)
		return c.HTML(404, "<h1>Confirmation Link might be expired please subscribe again </h1>")
	}
	fmt.Println(user.Email)
	err = dbase.ConfirmSub(n.db, user.Email)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Some Thing Wrong Happebd")
	}
	return c.HTML(404, "<h1>You are now subscribed to the newsletter </h1>")

}
