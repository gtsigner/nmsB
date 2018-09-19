package handler

import (
	"log"
	"net/http"
)

func AccessLog(handler http.Handler) http.Handler {
	// create new handler func to wrap given handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log the request
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		// execute given handler
		handler.ServeHTTP(w, r)
	})
}
