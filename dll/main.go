package main

import (
	"io/ioutil"
	"runtime"
	"time"
)

import "C"

func write(fpath string) {
	now := time.Now()
	d1 := []byte(now.String())
	err := ioutil.WriteFile(fpath, d1, 0644)
	if err != nil {
		//log.Panicln(err)
	}
}

//export ProcessAttached
func ProcessAttached() {
	runtime.LockOSThread()
	go write("/Temp/process-attached.txt")
}

//export ThreadAttached
func ThreadAttached() {
	runtime.LockOSThread()
	write("/Temp/thread-attached.txt")
}

func main() {}
