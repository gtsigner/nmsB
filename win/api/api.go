package api

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

var (
	// load dlls
	modpsapi = windows.NewLazySystemDLL("psapi.dll")

	//load func [psapi.dll]
	procGetModuleBaseName  = modpsapi.NewProc("GetModuleBaseNameW")
	procEnumProcessModules = modpsapi.NewProc("EnumProcessModules")
)

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	return e
}

func EnumProcessModules(process windows.Handle, modules []windows.Handle) (int, error) {
	var needed int32

	var first *windows.Handle
	handleSize := unsafe.Sizeof(first)

	if modules != nil {
		first = &modules[0]
	}

	r1, _, e1 := syscall.Syscall6(procEnumProcessModules.Addr(), 4,
		uintptr(process),
		uintptr(unsafe.Pointer(first)),
		handleSize*uintptr(len(modules)),
		uintptr(unsafe.Pointer(&needed)),
		0,
		0,
	)

	if r1 == 0 {
		if e1 != 0 {
			return 0, errnoErr(e1)
		}
		return 0, syscall.EINVAL
	}

	n := int(uintptr(needed) / handleSize)
	return n, nil
}

func GetModuleBaseName(process windows.Handle, module windows.Handle, moduleName *uint16, size uint32) (int, error) {
	r1, _, e1 := syscall.Syscall6(procGetModuleBaseName.Addr(),
		4,
		uintptr(process),
		uintptr(module),
		uintptr(unsafe.Pointer(moduleName)),
		uintptr(size),
		0,
		0,
	)

	if r1 == 0 {
		if e1 != 0 {
			return 0, errnoErr(e1)
		}
		return 0, syscall.EINVAL
	}

	return int(r1), nil
}
