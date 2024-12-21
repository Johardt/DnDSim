package handlers

import (
	"DnDSim/db"
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func RegisterSessionRoutes(g *echo.Group) {
	g.POST("", handleSessionPost)
	g.DELETE("", handleSessionDelete)
}

func handleSessionPost(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := db.GetUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.HTML(http.StatusUnauthorized, "Invalid email or password.")
		}
		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}

	if VerifyPassword(password, user.Password) != nil {
		return c.HTML(http.StatusUnauthorized, "Invalid email or password.")
	}

	if db.SessionExists(user.ID) {
		return c.Redirect(http.StatusSeeOther, "/index")
	}

	id, err := db.CreateSession(user.ID)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}

	c.SetCookie(&http.Cookie{
		Name:     "session",
		Value:    id,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	return c.Redirect(http.StatusSeeOther, "/index")
}

func handleSessionDelete(c echo.Context) error {
	cookie, err := c.Cookie("session")
	if err != nil {
		return c.HTML(http.StatusUnauthorized, "Unauthorized")
	}
	err = db.DeleteSession(cookie.Value)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.Redirect(http.StatusSeeOther, "/index")
}
