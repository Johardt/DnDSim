package handlers

import (
	"DnDSim/db"
	"DnDSim/views"
	"database/sql"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUserRoutes() {
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/email", handleUserEmail)
	http.HandleFunc("/users/password", handleUserPassword)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleUserPost(w, r)
	case http.MethodGet:
		handleUserGet(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handleUserPost(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Validate input
	if !isValidEmail(email) {
		http.Error(w, "Invalid email address.", http.StatusBadRequest)
		return
	}
	if !isValidPassword(password) {
		http.Error(w, "Password too short! Minimum 12 characters.", http.StatusBadRequest)
		return
	}

	// Check if user already exists
	_, err := db.GetUserByEmail(email)
	if err == nil {
		http.Error(w, "Email already registered.", http.StatusConflict)
		return
	} else if err != sql.ErrNoRows {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		http.Error(w, "Error processing password.", http.StatusInternalServerError)
		return
	}

	err = db.CreateUser(email, hashedPassword)
	if err != nil {
		http.Error(w, "Failed to create user.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully!"))
}

func handleUserGet(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email query parameter is required.", http.StatusBadRequest)
		return
	}

	user, err := db.GetUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found.", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error.", http.StatusInternalServerError)
		}
		return
	}

	// For security reasons, do not send the password hash
	response := struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
	}{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	responseString := "ID: " + strconv.Itoa(response.ID) + ", Email: " + response.Email + ", Created At: " + response.CreatedAt
	w.Write([]byte(responseString))
}

func isValidEmail(email string) bool {
	validDomains := `com|org|net|de|nl`
	regexPattern := `^[^@]+@[^@.]+\.(` + validDomains + `)$`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(email)
}

func handleUserEmail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	if !isValidEmail(email) {
		log.Println("Invalid email address provided.")
		views.UserInputField(
			"email",
			email,
			"Invalid email address provided.",
		).Render(r.Context(), w)
		return
	}

	views.UserInputField(
		"email",
		email,
		"",
	).Render(r.Context(), w)
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
	return len(password) >= 12
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func VerifyPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
