package main

// #cgo CFLAGS: -g -Wall
// #include "fix-main.hpp"
import "C"

import (
	"io/ioutil"
	"log"
	"time"
)

func main() {
	Attach()
}

//export Attach
func Attach() {

	//C.aatest()

	now := time.Now()

	d1 := []byte(now.String())
	err := ioutil.WriteFile("/Temp/test.txt", d1, 0644)
	if err != nil {
		log.Panicln(err)
	}

}
