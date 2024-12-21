package middleware

import (
	"DnDSim/db"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type contextKey string

const (
	userIDKey contextKey = "userID"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("session_id")
			if err != nil {
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			sessionId := cookie.Value
			session, err := db.GetSessionByID(sessionId)
			if err != nil {
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			if session.ExpiresAt.Before(time.Now()) {
				db.DeleteSession(sessionId)
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			err = db.UpdateSessionExpiration(sessionId)
			if err != nil {
				return c.String(http.StatusInternalServerError, "Internal Server Error")
			}

			ctx := context.WithValue(c.Request().Context(), userIDKey, session.UserID)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}

	}
}
