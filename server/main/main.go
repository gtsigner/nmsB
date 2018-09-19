package main

import (
	"log"

	"../../server"
)

var (
	/*VERSION the version of the server*/
	VERSION string
	/*RELEASE the release date of the server*/
	RELEASE string
)

func main() {
	err := server.Run(VERSION, RELEASE)
	if err != nil {
		log.Panicf("fail to start server, because %s", err)
	}
}
