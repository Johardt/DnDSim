package handlers

import (
	"DnDSim/db"
	"DnDSim/views/common"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserRoutes(g *echo.Group) {
	g.POST("", handleUserPost)
	g.GET("/:email", handleUserGet)
	g.POST("/username", handleUsername)
	g.POST("/email", handleUserEmail)
	g.POST("/password", handleUserPassword)
}

func handleUserPost(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	if err := ValidateUsername(username); err != nil {
		log.Println(err.Error())
		return RenderTempl(c, http.StatusUnprocessableEntity, common.Form("register-form", "/users", "Register"))
	}
	if err := ValidateEmail(email); err != nil {
		log.Println(err.Error())
		return RenderTempl(c, http.StatusUnprocessableEntity, common.Form("register-form", "/users", "Register"))
	}
	if err := ValidatePassword(password); err != nil {
		log.Println(err.Error())
		return RenderTempl(c, http.StatusUnprocessableEntity, common.Form("register-form", "/users", "Register"))
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error processing password: "+err.Error())
	}

	err = db.CreateUser(username, email, hashedPassword)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error creating user: "+err.Error())
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
			return c.String(http.StatusInternalServerError, "Internal Server Error: "+err.Error())
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

func handleUsername(c echo.Context) error {
	username := c.FormValue("username")

	if err := ValidateUsername(username); err != nil {
		log.Println(err.Error())
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.String(http.StatusOK, "")
}

func handleUserEmail(c echo.Context) error {
	email := c.FormValue("email")

	if err := ValidateEmail(email); err != nil {
		log.Println(err.Error())
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.String(http.StatusOK, "")
}

func handleUserPassword(c echo.Context) error {
	password := c.FormValue("password")

	if err := ValidatePassword(password); err != nil {
		log.Println(err.Error())
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.String(http.StatusOK, "")
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func VerifyPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
