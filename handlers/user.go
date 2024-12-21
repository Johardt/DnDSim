package handlers

import (
	"DnDSim/db"
	"DnDSim/views"
	"database/sql"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const MIN_PASSWORD_LENGTH = 3

// func RegisterUserRoutes() {
// 	http.HandleFunc("/users", handleUsers)
// 	http.HandleFunc("/users/email", handleUserEmail)
// 	http.HandleFunc("/users/password", handleUserPassword)
// }

func RegisterUserRoutes(g *echo.Group) {
	g.POST("", handleUserPost)
	g.GET("/:email", handleUserGet)
	g.POST("/email", handleUserEmail)
	g.POST("/password", handleUserPassword)
}

// TODO duplicate code and error handling is horrible
func handleUserPost(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Validate input
	if !isValidEmail(email) {
		// TODO this should also be a 422, but should not replace the form
		return c.String(http.StatusBadRequest, "Invalid email address.")
	}
	if !isValidPassword(password) {
		// TODO same here
		return c.String(http.StatusBadRequest, "Password too short! Minimum 12 characters.")
	}

	// Check if user already exists
	// TODO check for username as well
	_, err := db.GetUserByEmail(email)
	if err == nil {
		return c.String(http.StatusConflict, "Email already registered.")
	} else if err != sql.ErrNoRows {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error processing password.")
	}

	err = db.CreateUser(username, email, hashedPassword)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create user.")
	}

	return c.String(http.StatusCreated, "User created successfully!")
}

func handleUserGet(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.String(http.StatusBadRequest, "Email query parameter is required.")
	}

	user, err := db.GetUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "User not found.")
		} else {
			return c.String(http.StatusInternalServerError, "Internal Server Error.")
		}
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
	return c.String(http.StatusOK, responseString)
}

func isValidEmail(email string) bool {
	validDomains := `com|org|net|de|nl`
	regexPattern := `^[^@]+@[^@.]+\.(` + validDomains + `)$`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(email)
}

func handleUserEmail(c echo.Context) error {
	email := c.FormValue("email")

	if !isValidEmail(email) {
		log.Println("Invalid email address provided.")
		return RenderTempl(c, http.StatusUnprocessableEntity,
			views.UserInputField("email", email, "Invalid email address provided."),
		)
	}

	return RenderTempl(c, http.StatusOK, views.UserInputField("email", email, ""))
}

func handleUserPassword(c echo.Context) error {
	password := c.FormValue("password")

	if !isValidPassword(password) {
		return RenderTempl(c, http.StatusUnprocessableEntity,
			views.UserPasswordField(password, "Password too short! Minimum "+strconv.Itoa(MIN_PASSWORD_LENGTH)+" characters."),
		)
	}

	return RenderTempl(c, http.StatusOK, views.UserPasswordField(password, ""))
}

func isValidPassword(password string) bool {
	return len(password) >= MIN_PASSWORD_LENGTH
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func VerifyPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
