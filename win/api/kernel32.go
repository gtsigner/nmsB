package api

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// load [kernel32.dll]
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	//load func [kernel32.dll]
	procReadProcessMemory = modkernel32.NewProc("ReadProcessMemory")
)

func ReadProcessMemory(handle windows.Handle, address uintptr, size uint64) ([]byte, error) {
	nbr := uintptr(0)
	data := make([]byte, size)

	r1, _, e1 := syscall.Syscall6(procReadProcessMemory.Addr(),
		5,
		uintptr(handle),
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(size),
		uintptr(unsafe.Pointer(&nbr)),
		0)

	if r1 == 0 {
		if e1 != 0 {
			return nil, errnoErr(e1)
		}
		return nil, syscall.EINVAL
	}

	return data, nil
}
