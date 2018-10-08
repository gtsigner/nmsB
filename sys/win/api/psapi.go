package api

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// load [psapi.dll]
	modpsapi = windows.NewLazySystemDLL("psapi.dll")

	//load func [psapi.dll]
	procGetModuleBaseName    = modpsapi.NewProc("GetModuleBaseNameW")
	procEnumProcessModules   = modpsapi.NewProc("EnumProcessModules")
	procGetModuleInformation = modpsapi.NewProc("GetModuleInformation")
)

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

func GetModuleInformation(process windows.Handle, module windows.Handle) (*ModuleInfo, error) {
	var moduleInfo ModuleInfo
	r1, _, e1 := syscall.Syscall6(procGetModuleInformation.Addr(),
		4,
		uintptr(process),
		uintptr(module),
		uintptr(unsafe.Pointer(&moduleInfo)),
		uintptr(unsafe.Sizeof(moduleInfo)),
		0,
		0,
	)

	if r1 == 0 {
		if e1 != 0 {
			return nil, errnoErr(e1)
		}
		return nil, syscall.EINVAL
	}

	return &moduleInfo, nil
}
