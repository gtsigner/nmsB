package api

import (
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// load [kernel32.dll]
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	//load func [kernel32.dll]
	procVirtualQuery       = modkernel32.NewProc("VirtualQuery")
	procVirtualFreeEx      = modkernel32.NewProc("VirtualFreeEx")
	procVirtualQueryEx     = modkernel32.NewProc("VirtualQueryEx")
	procVirtualAllocEx     = modkernel32.NewProc("VirtualAllocEx")
	procVirtualProtectEx   = modkernel32.NewProc("VirtualProtectEx")
	procGetModuleHandleW   = modkernel32.NewProc("GetModuleHandleW")
	procGetExitCodeThread  = modkernel32.NewProc("GetExitCodeThread")
	procReadProcessMemory  = modkernel32.NewProc("ReadProcessMemory")
	procCreateRemoteThread = modkernel32.NewProc("CreateRemoteThread")
	procWriteProcessMemory = modkernel32.NewProc("WriteProcessMemory")
)

func GetLoadLibrary() (uintptr, error) {
	handle, err := GetModuleHandle("kernel32")
	if err != nil {
		return 0, err
	}

	address, err := windows.GetProcAddress(handle, "LoadLibraryW")
	if err != nil {
		return 0, err
	}
	return address, nil
}

func GetFreeLibrary() (uintptr, error) {
	handle, err := GetModuleHandle("kernel32")
	if err != nil {
		return 0, err
	}

	address, err := windows.GetProcAddress(handle, "FreeLibrary")
	if err != nil {
		return 0, err
	}
	return address, nil
}

func GetModuleHandle(name string) (windows.Handle, error) {
	var namePtr *uint16
	if name != "" {
		namePtr = syscall.StringToUTF16Ptr(name)
	}

	r0, _, e1 := syscall.Syscall(procGetModuleHandleW.Addr(),
		1,
		uintptr(unsafe.Pointer(namePtr)),
		0,
		0)
	runtime.KeepAlive(name)

	if r0 == 0 {
		if e1 != 0 {
			return 0, errnoErr(e1)
		}
		return 0, syscall.EINVAL
	}
	return windows.Handle(r0), nil
}

func ReadProcessMemory(handle windows.Handle, address uintptr, size uint64) ([]byte, error) {
	nbr := uintptr(0)
	data := make([]byte, size)

	r0, _, e1 := syscall.Syscall6(procReadProcessMemory.Addr(),
		5,
		uintptr(handle),
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(size),
		uintptr(unsafe.Pointer(&nbr)),
		0)

	if r0 == 0 {
		if e1 != 0 {
			return nil, errnoErr(e1)
		}
		return nil, syscall.EINVAL
	}

	return data, nil
}

func VirtualAllocEx(handle windows.Handle, address uintptr, size uint32, alloctype uint32, protect uint32) (uintptr, error) {
	r0, _, e1 := syscall.Syscall6(procVirtualAllocEx.Addr(),
		5,
		uintptr(handle),
		uintptr(address),
		uintptr(size),
		uintptr(alloctype),
		uintptr(protect),
		0)
	value := uintptr(r0)

	if value == 0 {
		if e1 != 0 {
			return value, errnoErr(e1)
		} else {
			return value, syscall.EINVAL
		}
	}
	return value, nil
}

func VirtualQuery(address uintptr) (*MemoryInformation, error) {
	var memoryInfo MemoryInformation
	r0, _, e1 := syscall.Syscall(procVirtualQuery.Addr(),
		3,
		address,
		uintptr(unsafe.Pointer(&memoryInfo)),
		uintptr(unsafe.Sizeof(memoryInfo)))

	if r0 == 0 {
		if e1 != 0 {
			return nil, errnoErr(e1)
		} else {
			return nil, syscall.EINVAL
		}
	}
	return &memoryInfo, nil
}

func VirtualQueryEx(handle windows.Handle, address uintptr) (*MemoryInformation, error) {
	var memoryInfo MemoryInformation
	r0, _, e1 := syscall.Syscall6(procVirtualQueryEx.Addr(),
		4,
		uintptr(handle),
		address,
		uintptr(unsafe.Pointer(&memoryInfo)),
		uintptr(unsafe.Sizeof(memoryInfo)),
		0,
		0)
	if r0 == 0 {
		if e1 != 0 {
			return nil, errnoErr(e1)
		} else {
			return nil, syscall.EINVAL
		}
	}
	return &memoryInfo, nil
}

func VirtualFreeEx(handle windows.Handle, address uintptr, size uint32, freeType uint32) error {
	r0, _, e1 := syscall.Syscall6(procVirtualAllocEx.Addr(),
		4,
		uintptr(handle),
		address,
		uintptr(size),
		uintptr(freeType),
		0,
		0)

	if r0 == 0 {
		if e1 != 0 {
			return errnoErr(e1)
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func WriteProcessMemory(process windows.Handle, address uintptr, buffer unsafe.Pointer, size uint32) (uint32, error) {
	var nLength uint32
	r0, _, e1 := syscall.Syscall6(procWriteProcessMemory.Addr(),
		5,
		uintptr(process),
		address,
		uintptr(buffer),
		uintptr(size),
		uintptr(unsafe.Pointer(&nLength)),
		0)

	if r0 == 0 {
		if e1 != 0 {
			return 0, errnoErr(e1)
		}
		return 0, syscall.EINVAL
	}

	return nLength, nil
}

func CreateRemoteThread(process windows.Handle, sa *windows.SecurityAttributes, stackSize uint32, startAddress uintptr, parameter uintptr, creationFlags uint32) (windows.Handle, uint32, error) {
	var threadID uint32
	r0, _, e1 := syscall.Syscall9(procCreateRemoteThread.Addr(),
		7,
		uintptr(process),
		uintptr(unsafe.Pointer(sa)),
		uintptr(stackSize),
		startAddress,
		parameter,
		uintptr(creationFlags),
		uintptr(unsafe.Pointer(&threadID)),
		0,
		0)

	runtime.KeepAlive(sa)

	if r0 == 0 {
		if e1 != 0 {
			return 0, 0, errnoErr(e1)
		}
		return 0, 0, syscall.EINVAL
	}

	return windows.Handle(r0), threadID, nil
}

func GetExitCodeThread(thread windows.Handle) (uint32, error) {
	var exitCode uint32
	r0, _, e1 := syscall.Syscall(procGetExitCodeThread.Addr(),
		2,
		uintptr(thread),
		uintptr(unsafe.Pointer(&exitCode)),
		0)
	if r0 == 0 {
		if e1 != 0 {
			return 0, errnoErr(e1)
		}
		return 0, syscall.EINVAL
	}

	return exitCode, nil
}
