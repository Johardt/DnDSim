package main

import (
	"DnDSim/db"
	"DnDSim/routes"
	"flag"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitializeDB("main.db")
	defer db.CloseDB()

	e := echo.New()
	routes.RegisterRoutes(e)

	port := "8080"
	flagSet := flag.NewFlagSet("port", flag.ExitOnError)
	flagSet.StringVar(&port, "port", "8080", "Port to run the server on")
	flagSet.StringVar(&port, "p", "8080", "Port to run the server on (shorthand)")
	flagSet.Parse(os.Args[1:])

	log.Printf("Server running on localhost:%s\n", port)
	e.Logger.Fatal(e.Start("localhost:" + port))
}
