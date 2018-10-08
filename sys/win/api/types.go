package api

import (
	"unsafe"
)

// WIdnows Types: https://github.com/JamesHovious/w32/blob/master/typedef.go

type MemoryInformation struct {
	BaseAddress       unsafe.Pointer
	AllocationBase    unsafe.Pointer
	AllocationProtect uint32
	RegionSize        uintptr
	State             uint32
	Protect           uint32
	Type              uint32
}

type ModuleInfo struct {
	BaseOfDll   unsafe.Pointer
	SizeOfImage uint32
	EntryPoint  unsafe.Pointer
}
