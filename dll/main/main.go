package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"log"

	"../../config"
	"../../dll"
	"../../message/json"
)

var (
	VERSION string
	RELEASE string
)

//export Init
func Init(cParameter *C.char) {
	// cpnvert CString to go string
	parameter := C.GoString(cParameter)
	// execute the dll
	err := dll.Run(VERSION, RELEASE, parameter)
	// check for error
	if err != nil {
		log.Panicln(err)
	}
}

func goRun() error {
	log.Println("executing dll via go run wrapper...")
	// load the config
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// convert config to json string
	data, err := json.Encode(cfg)
	if err != nil {
		return err
	}

	// convert the config to parameter
	cParameter := C.CString(data)
	// start the Dll
	Init(cParameter)

	return nil
}

// template function
func main() {
	err := goRun()
	if err != nil {
		log.Panicln(err)
	}
}
