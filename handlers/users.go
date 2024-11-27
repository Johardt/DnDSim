package handlers

import (
	"DnDSim/views"
	"log"
	"net/http"
	"regexp"
)

func RegisterUserRoutes() {
	http.HandleFunc("/users", handleUserCreation)

	http.HandleFunc("/users/email", handleUserEmail)
	http.HandleFunc("/users/password", handleUserPassword)
}

func handleUserCreation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		if isValidEmail(email) && isValidPassword(password) {
			w.Write([]byte("User created successfully!"))
			return
		} else {
			http.Error(w, "Invalid email or password", http.StatusUnprocessableEntity)
		}
	}
}

func handleUserEmail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	if !isValidEmail(email) {
		log.Printf("Invalid email address provided: %s\n", email)
		views.UserInputField(
			"email",
			email,
			"Invalid email address provided.",
		).Render(r.Context(), w)
		return
	}

	views.UserInputField(
		"email",
		r.FormValue("email"),
		"",
	).Render(r.Context(), w)
}

func isValidEmail(email string) bool {
	validDomains := `com|org|net|de|nl`
	regexPattern := `^[^@]+@[^@.]+\.(` + validDomains + `)$`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(email)
}

func handleUserPassword(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")

	if !isValidPassword(password) {
		views.UserPasswordField(password, "Password too short! Minimum 12 characters.").Render(r.Context(), w)
		return
	}

	views.UserPasswordField(password, "").Render(r.Context(), w)
}

func isValidPassword(password string) bool {
	regexPattern := `^.{12,}$`
	re := regexp.MustCompile(regexPattern)

	return re.MatchString(password)
}
