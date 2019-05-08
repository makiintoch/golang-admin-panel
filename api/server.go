package api

import (
	"log"
	"net/http"
)

// Start runs ListenAndServe on the http.Server
func Start() {
	log.Println("Starting server...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
