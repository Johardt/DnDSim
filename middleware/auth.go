package middleware

import (
	"DnDSim/db"
	"context"
	"net/http"
	"time"
)

type contextKey string

const (
	userIDKey contextKey = "userID"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sessionID := cookie.Value
		session, err := db.GetSessionByID(sessionID)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if session.ExpiresAt.Before(time.Now()) {
			db.DeleteSession(sessionID)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		err = db.UpdateSessionExpiration(sessionID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, session.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
