package main

import (
	"DnDSim/db"
	"DnDSim/handlers"
	"DnDSim/views"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	db.InitializeDB("users.db")
	defer db.CloseDB()

	http.Handle("/", templ.Handler(views.BasePage()))
	http.Handle("/index", templ.Handler(views.IndexPage()))
	http.Handle("/register", templ.Handler(views.RegisterPage()))

	handlers.RegisterUserRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
