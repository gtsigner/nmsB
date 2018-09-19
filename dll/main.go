package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"io/ioutil"
	"time"
)

var (
	VERSION string
	RELEASE string
)

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

//export Init
func Init(cParameter *C.char) {
	// cpnvert CString to go string
	parameter := C.GoString(cParameter)

	go writeData("/Temp/main.txt", fmt.Sprintf("%s, version: %s", parameter, VERSION))
}

func main() {}
