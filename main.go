package main

import (
	"crypto/tls"
	"DnDSim/db"
	"DnDSim/routes"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitializeDB("main.db")
	defer db.CloseDB()

	e := echo.New()
	routes.RegisterRoutes(e)

	port := "8080"
	certFile := "cert.pem"
	keyFile := "key.pem"

	flagSet := flag.NewFlagSet("port", flag.ExitOnError)
	flagSet.StringVar(&port, "port", "8080", "Port to run the server on")
	flagSet.StringVar(&port, "p", "8080", "Port to run the server on (shorthand)")
	flagSet.StringVar(&certFile, "cert", "cert.pem", "Path to the TLS certificate file")
	flagSet.StringVar(&keyFile, "key", "key.pem", "Path to the TLS key file")
	flagSet.Parse(os.Args[1:])

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		log.Fatalf("TLS certificate file not found: %s", certFile)
	}
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		log.Fatalf("TLS key file not found: %s", keyFile)
	}

	log.Printf("Server running on localhost:%s\n", port)
	e.Logger.Fatal(e.StartTLS("localhost:"+port, certFile, keyFile))
}
