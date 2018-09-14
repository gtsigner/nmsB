package main

import "C"

import (
	"io/ioutil"
	"time"
)

func main() {
	Dllmain()
}

func Dllmain() {

	now := time.Now()

	d1 := []byte(now.String())
	ioutil.WriteFile("C:\\test.txt", d1, 0644)

}
