package main

import (
	"DnDSim/db"
	"DnDSim/handlers"
	"DnDSim/middleware"
	"DnDSim/views"
	"flag"
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

	http.Handle("/play", middleware.Auth(templ.Handler(views.GameSelector())))

	handlers.RegisterUserRoutes()
	handlers.RegisterSessionRoutes()

	port := "8080"
	flagSet := flag.NewFlagSet("port", flag.ExitOnError)
	flagSet.StringVar(&port, "port", "8080", "Port to run the server on")
	flagSet.StringVar(&port, "p", "8080", "Port to run the server on (shorthand)")
	flagSet.Parse(os.Args[1:])

	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
