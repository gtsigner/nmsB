package signature

import (
	"log"
	"unsafe"
)

func Scan(start uintptr, length int, p *Pattern) uintptr {

	for i := 0; i < length; i++ {
		index := start + uintptr(i)
		b := (*byte)(unsafe.Pointer(index))
		log.Println(b)
	}

	return 0
}
