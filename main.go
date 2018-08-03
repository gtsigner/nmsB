package main

import (
	"./win/modules"
	"./win/process"
	"log"
	"unsafe"
)

var (
	VERSION string
	RELEASE string
)

func main() {
	log.Println(VERSION, RELEASE)

	p, err := process.FindProcess("test.exe")
	if err != nil {
		log.Panicln(err)
	}

	handle, err := process.Open(p.Id)
	if err != nil {
		log.Panicln(err)
	}

	module, err := modules.Find(handle, "test.exe")
	if err != nil {
		log.Panicln(err)
	}

	log.Println(module.Name)

	log.Printf("module: 0x%x\n", module.Handle)

	p1 := (*unsafe.Pointer)(unsafe.Pointer(uintptr(module.Handle) + uintptr(0x5)))
	log.Printf("p1: %x\n", *p1)

	// +0015CEB0 0x18 0x68 0x40 0xF0

}
