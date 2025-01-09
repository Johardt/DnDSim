package handlers

import (
	"DnDSim/db"
	"DnDSim/views"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func RegisterSessionRoutes(g *echo.Group) {
	g.POST("", handleSessionPost)
	g.DELETE("", handleSessionDelete)
	g.POST("/check", handleCheck)
	g.GET("/session_buttons", getSessionButtons)
}

func handleSessionPost(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := db.GetUserByName(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.HTML(http.StatusUnauthorized, "Invalid username or password.")
		}
		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}

	if VerifyPassword(password, user.Password) != nil {
		return c.HTML(http.StatusUnauthorized, "Invalid username or password.")
	}

	id, err := db.GetSessionID(user.ID)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}

	c.SetCookie(&http.Cookie{
		Name:     "session",
		Value:    id,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	c.Response().Header().Add("HX-Redirect", "/")
	return c.HTML(http.StatusSeeOther, "")
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

	c.Response().Header().Add("HX-Redirect", "/")
	return c.HTML(http.StatusSeeOther, "Redirecting...")
}

func handleCheck(c echo.Context) error {
	// Check if there is a username form value
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, _ := db.GetUserByName(username)
	if user == nil {
		return c.String(http.StatusUnauthorized, "Invalid username or password.")
	}

	if VerifyPassword(password, user.Password) != nil {
		return c.String(http.StatusUnauthorized, "Invalid username or password.")
	}

	return c.String(http.StatusOK, "")
}

// This function returns a set of buttons depending on who is logged in.
func getSessionButtons(c echo.Context) error {
	cookie, err := c.Cookie("session")
	if err != nil {
		log.Println("No session token, return login and register button")
		return RenderTempl(c, 200, views.AuthButtons())
	}
	session, err := db.GetSessionByID(cookie.Value)
	if err != nil {
		log.Println("Session ID invalid, return login and register buttons")
		return RenderTempl(c, 200, views.AuthButtons())
	}

	log.Println("Session valid")
	user, err := db.GetUserByID(session.UserID)
	if err != nil {
		log.Println("Error getting user by ID: " + err.Error())
		return RenderTempl(c, 200, views.AuthButtons())
	}
	return RenderTempl(c, 200, views.ProfileButtons(user))
}
