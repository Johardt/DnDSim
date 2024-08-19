package main

import (
	"github.com/a-h/templ"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", templ.Handler(indexPage()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
