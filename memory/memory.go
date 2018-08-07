package memory

import (
	"unsafe"
)

var (
	NULL = uintptr(unsafe.Pointer(nil))
)
