package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"syscall"
	"time"
)

import "C"

func write(fpath string) {
	now := time.Now()
	writeData(fpath, now.String())
}

func writeData(fpath string, data string) {
	d1 := []byte(data)
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

//export Init
func Init(C.CString) {

	content := fmt.Sprintf(", %s", , s)
	writeData("/Temp/main.txt", content)

}

func main() {}
