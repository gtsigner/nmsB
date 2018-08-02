package main

import (
	"./win/process"
	"log"
)

var (
	VERSION string
	RELEASE string
)

func main() {
	log.Println(VERSION, RELEASE)

	process.PS()

}
