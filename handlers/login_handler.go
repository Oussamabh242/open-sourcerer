package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/Oussamabh242/open-sourcerer/views/home"
	"github.com/Oussamabh242/open-sourcerer/views/login"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	return Render(c, 200, login.Login())
}

func Signin(c echo.Context) error {

	email := c.FormValue("email")
	password := c.FormValue("password")
	token, err := auth(email, password)
	if err != nil {
		log.Println(err)
		return c.HTML(400, "error happend")
	}
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.HTML(http.StatusOK, "")

}

var key = []byte("secret")

func CreateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = os.Getenv("THISUSER")
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func auth(email, password string) (string, error) {
	OSemail := []byte(os.Getenv("Email"))
	if bcrypt.CompareHashAndPassword(OSemail, []byte(email)) != nil {
		return "", fmt.Errorf("No account found")
	}
	OSpassword := os.Getenv("Password")
	if bcrypt.CompareHashAndPassword([]byte(OSpassword), []byte(password)) != nil {
		return "", fmt.Errorf("Wrong Password")
	}
	token, err := CreateToken()
	if err != nil {
		return "", fmt.Errorf("Error generating jwt token")
	}
	return token, nil

}

func writeCookie(c echo.Context) error {
	return c.String(http.StatusOK, "write a cookie")
}

func MainHandler(c echo.Context) error {
	log.Println("this is Stop")
	return Render(c, http.StatusOK, home.Index())
}

func Middleware(next, stop echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			log.Println(err)
			c.Response().Header().Set("Hx-Redirect", "/")
			return stop(c)
		}
		_, err = DecodeToken(cookie.Value)
		if err != nil {
			log.Println(err)
			c.Response().Header().Set("Hx-Redirect", "/")
			return stop(c)
		}
		return next(c)
	}
}

func DecodeToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {

		return nil, fmt.Errorf("invalid token")

	}
	if claims["user"] != os.Getenv("THISUSER") {
		return nil, fmt.Errorf("Wrong User")
	}
	return token, nil

}
