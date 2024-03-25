package main

import (
	"log"
	"net/http"

	"github.com/ayush/ide/route"
)

func main() {
	route.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
