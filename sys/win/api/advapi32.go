package api

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// load dll [advapi32.dll]
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")

	//load func [advapi32.dll]
	procLookupPrivilegeName   = modadvapi32.NewProc("LookupPrivilegeNameW")
	procLookupPrivilegeValue  = modadvapi32.NewProc("LookupPrivilegeValueW")
	procAdjustTokenPrivileges = modadvapi32.NewProc("AdjustTokenPrivileges")
)

func LookupPrivilegeValue(systemname *uint16, name *uint16, luid *uint64) error {
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

func AdjustTokenPrivileges(token windows.Token, releaseAll bool, input *byte, outputSize uint32, output *byte, requiredSize *uint32) (uint32, error) {
	var _p0 uint32
	if releaseAll {
		_p0 = 1
	} else {
		_p0 = 0
	}

	r0, _, e1 := syscall.Syscall6(procAdjustTokenPrivileges.Addr(),
		6,
		uintptr(token),
		uintptr(_p0),
		uintptr(unsafe.Pointer(input)),
		uintptr(outputSize),
		uintptr(unsafe.Pointer(output)),
		uintptr(unsafe.Pointer(requiredSize)))

	if r0 == 0 {
		if e1 != 0 {
			return 0, errnoErr(e1)
		}
		return 0, syscall.EINVAL

	}
	return uint32(r0), nil
}

func LookupPrivilegeName(systemName *uint16, luid *uint64, buffer *uint16, size *uint32) error {
	r1, _, e1 := syscall.Syscall6(procLookupPrivilegeName.Addr(),
		4,
		uintptr(unsafe.Pointer(systemName)),
		uintptr(unsafe.Pointer(luid)),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(size)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			return error(e1)
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}
