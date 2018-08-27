package handler

import (
	"log"
	"net/http"
)

func FileServer(directory string) http.Handler {
	// create file handler
	fileServer := http.FileServer(http.Dir(directory))

	log.Printf("creating static file server for directory [ %s ]\n", directory)

	// create new handler func
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// lets handle the file request
		fileServer.ServeHTTP(response, request)
	})
}
