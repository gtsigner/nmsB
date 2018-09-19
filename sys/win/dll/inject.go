package dll

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"../api"

	"golang.org/x/sys/windows"
)

func MakeInheritSa() *windows.SecurityAttributes {
	var sa windows.SecurityAttributes
	sa.Length = uint32(unsafe.Sizeof(sa))
	sa.InheritHandle = 1
	return &sa
}

func InjectDll(handle windows.Handle, dllPath string) (windows.Handle, error) {

	name, err := syscall.UTF16FromString(dllPath)
	if err != nil {
		return 0, fmt.Errorf("fail to convert [ %s ] to []uint16, because %s", dllPath, err.Error())
	}

	nameLength := uint32((len(name) + 1)) * uint32(unsafe.Sizeof(name[0]))

	remoteAddress, err := api.VirtualAllocEx(handle, 0, nameLength, windows.MEM_COMMIT, windows.PAGE_READWRITE)
	if err != nil {
		return 0, fmt.Errorf("fail to allocate memory, because %s", err.Error())
	}
	defer api.VirtualFreeEx(handle, remoteAddress, 0, windows.MEM_RELEASE)

	_, err = api.WriteProcessMemory(handle, remoteAddress, unsafe.Pointer(&name[0]), nameLength)
	if err != nil {
		return 0, fmt.Errorf("fail to write memory [ %d bytes ] to [ 0x%X ], because %s", nameLength, remoteAddress, err.Error())
	}

	loadLibraryW, err := api.GetLoadLibrary()
	if err != nil {
		return 0, fmt.Errorf("fail find base address for load-libaray, because %s", err.Error())
	}

	runtime.KeepAlive(name)
	thread, _, err := api.CreateRemoteThread(handle, MakeInheritSa(), 0, loadLibraryW, remoteAddress, 0)
	if err != nil {
		return 0, fmt.Errorf("fail to create remote thread, because %s", err.Error())
	}
	defer windows.CloseHandle(thread)

	wr, err := windows.WaitForSingleObject(thread, windows.INFINITE)
	if err != nil {
		return 0, fmt.Errorf("fail to wait for single object, because %s", err.Error())
	}

	if wr != windows.WAIT_OBJECT_0 {
		return 0, fmt.Errorf("Unexpected wait result %d", wr)
	}

	exitCode, err := api.GetExitCodeThread(thread)
	if err != nil {
		return 0, fmt.Errorf("fail to get exit-code from thread, because %s", err.Error())
	}

	return windows.Handle(exitCode), nil
}

func UnInjectDll(handle windows.Handle, dllHandle windows.Handle) error {
	freeLibrary, err := api.GetFreeLibrary()
	if err != nil {
		return nil
	}

	thread, _, err := api.CreateRemoteThread(handle, nil, 0, freeLibrary, uintptr(dllHandle), 0)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(thread)

	wr, err := windows.WaitForSingleObject(thread, windows.INFINITE)
	if err != nil {
		return err
	}

	if wr != windows.WAIT_OBJECT_0 {
		return fmt.Errorf("Unexpected wait result %d", wr)
	}

	exitCode, err := api.GetExitCodeThread(thread)
	if err != nil {
		return err
	}

	if exitCode == 0 {
		return fmt.Errorf("fail to unload dll, because thread exit with zero")
	}

	return nil
}
