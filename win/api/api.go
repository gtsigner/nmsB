package api

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// load dlls
	modpsapi    = windows.NewLazySystemDLL("psapi.dll")
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	//load func [kernel32.dll]
	procReadProcessMemory = modkernel32.NewProc("ReadProcessMemory")

	//load func [advapi32.dll]
	procLookupPrivilegeValue  = modadvapi32.NewProc("LookupPrivilegeValueW")
	procAdjustTokenPrivileges = modadvapi32.NewProc("AdjustTokenPrivileges")

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

func LookupPrivilegeValue(systemname *uint16, name *uint16, luid *LUID) error {
	r1, _, e1 := syscall.Syscall(procLookupPrivilegeValue.Addr(),
		3,
		uintptr(unsafe.Pointer(systemname)),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(luid)))
	if r1 == 0 {
		if e1 != 0 {
			return errnoErr(e1)
		}
		return syscall.EINVAL

	}
	return nil
}

func AdjustTokenPrivileges(token windows.Token, disableAllPrivileges bool, newstate *TOKEN_PRIVILEGES, buflen uint32, prevstate *TOKEN_PRIVILEGES, returnlen *uint32) (uint32, error) {
	var _p0 uint32
	if disableAllPrivileges {
		_p0 = 1
	} else {
		_p0 = 0
	}

	r0, _, e1 := syscall.Syscall6(procAdjustTokenPrivileges.Addr(),
		6,
		uintptr(token),
		uintptr(_p0),
		uintptr(unsafe.Pointer(newstate)),
		uintptr(buflen),
		uintptr(unsafe.Pointer(prevstate)),
		uintptr(unsafe.Pointer(returnlen)))

	if r0 == 0 {
		if e1 != 0 {
			return 0, errnoErr(e1)
		}
		return 0, syscall.EINVAL

	}
	return uint32(r0), nil
}
