package main

import (
	"DnDSim/handlers"
	"DnDSim/views"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", templ.Handler(views.IndexPage()))
	http.Handle("/register", templ.Handler(views.RegisterPage()))

	handlers.RegisterUserRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
