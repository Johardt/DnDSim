package handlers

import (
	"DnDSim/db"
	"database/sql"
	"net/http"
	"time"
)

// TODO Set cookie to secure (HTTPS only)

func RegisterSessionRoutes() {
	http.HandleFunc("/sessions", handleSessions)
}

func handleSessions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleSessionPost(w, r)
	case http.MethodDelete:
		handleSessionDelete(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handleSessionPost(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := db.GetUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email or password.", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if VerifyPassword(password, user.Password) != nil {
		http.Error(w, "Invalid email or password.", http.StatusUnauthorized)
		return
	}

	id, err := db.CreateSession(user.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    id,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
}

func handleSessionDelete(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	err = db.DeleteSession(cookie.Value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
