package main

import (
	"log"
	"net/http"

	"github.com/gopherguides/training/testing/async/src/httpd"
	"github.com/gopherguides/training/testing/async/src/keys"
)

func main() {
	handler := httpd.NewHandler()
	handler.Store = keys.NewStore()

	log.Println("starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
