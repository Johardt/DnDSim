package handlers

import (
	"DnDSim/views"
	"log"
	"net/http"
)

func RegisterUserRoutes() {
	http.HandleFunc("/users/email", handleUserEmail)
	http.HandleFunc("/users/password", handleUserPassword)
}

func handleUserEmail(w http.ResponseWriter, r *http.Request) {
	log.Printf("Email: %s\n", r.FormValue("email"))

	views.UserInputField(
		"Email Address",
		r.FormValue("email"),
		"Invalid email address provided.",
	).Render(r.Context(), w)
}

func handleUserPassword(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")

	// Validate password length
	if len(password) < 8 {
		views.UserPasswordField(password, "Password too short! Minimum 8 characters.").Render(r.Context(), w)
		return
	}

	views.UserPasswordField(password, "").Render(r.Context(), w)
}
