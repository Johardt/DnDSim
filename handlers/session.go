package handlers

import (
	"DnDSim/db"
	"database/sql"
	"net/http"
)

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
	if err == nil {
		http.Error(w, "Email already registered.", http.StatusConflict)
		return
	} else if err != sql.ErrNoRows {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if VerifyPassword(password, user.Password) != nil {
		http.Error(w, "Incorrect password.", http.StatusUnauthorized)
		return
	} else {

	}
}

func handleSessionDelete(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}