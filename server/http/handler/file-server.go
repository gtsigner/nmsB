package handler

import (
	"log"
	"net/http"
	"path/filepath"
)

func toAbsPath(directory string) (string, error) {
	dpath, err := filepath.Abs(directory)
	return dpath, err
}

func FileServer(directory string) (http.Handler, error) {
	// make the directory path absolute
	absDirectory, err := toAbsPath(directory)
	if err != nil {
		return nil, err
	}

	// create file handler
	fileServer := http.FileServer(http.Dir(absDirectory))
	log.Printf("serving static files from directory [ %s ]\n", absDirectory)

	// create new handler func
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// lets handle the file request
		fileServer.ServeHTTP(response, request)
	}), nil
}
