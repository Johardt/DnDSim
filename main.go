package main

import (
	"DnDSim/db"
	"DnDSim/handlers"
	"DnDSim/views"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

func main() {
	db.InitializeDB("users.db")
	defer db.CloseDB()

	http.Handle("/", templ.Handler(views.BasePage()))
	http.Handle("/index", templ.Handler(views.IndexPage()))
	http.Handle("/login", templ.Handler(views.LoginForm()))
	http.Handle("/register", templ.Handler(views.RegisterPage()))

	handlers.RegisterUserRoutes()

	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
