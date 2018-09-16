package main

// #include <windows.h>
// extern void Attach();
// static inline void call_attach() {
//  Attach();
// }
// BOOL WINAPI DllMain(
// 	HINSTANCE hinstDLL, // handle to DLL module
//  DWORD fdwReason, // reason for calling function
//  LPVOID lpReserved) // reserved
// {
//  Attach();
//	return TRUE;
// }
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

	now := time.Now()

	d1 := []byte(now.String())
	err := ioutil.WriteFile("/Temp/test.txt", d1, 0644)
	if err != nil {
		log.Panicln(err)
	}

}
