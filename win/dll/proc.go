package dll

import (
	"fmt"
	"log"
	"path/filepath"
	"syscall"
	"unsafe"

	"../api"
	"../modules"

	"golang.org/x/sys/windows"
)

func CallRemoteProc(handle windows.Handle, dllPath string, procName string, parameter string) error {
	procAddress, err := GetRemoteProcAddress(handle, dllPath, procName)
	if err != nil {
		return err
	}

	data, err := syscall.UTF16FromString(parameter)
	if err != nil {
		return fmt.Errorf("fail to convert [ %s ] to []uint16, because %s", parameter, err.Error())
	}

	dataLength := uint32((len(data) + 1)) * uint32(unsafe.Sizeof(data[0]))

	parameterRemoteAddress, err := api.VirtualAllocEx(handle, 0, dataLength, windows.MEM_COMMIT, windows.PAGE_READWRITE)
	if err != nil {
		return fmt.Errorf("fail to allocate memory, because %s", err.Error())
	}
	//defer api.VirtualFreeEx(handle, parameterRemoteAddress, 0, windows.MEM_RELEASE)
	log.Printf("parameterRemoteAddress: 0x%X", parameterRemoteAddress)

	_, err = api.WriteProcessMemory(handle, parameterRemoteAddress, unsafe.Pointer(&data[0]), dataLength)
	if err != nil {
		return fmt.Errorf("fail to write parameter data [ %d bytes ] to [ 0x%X ], because %s",
			dataLength, parameterRemoteAddress, err.Error())
	}

	thread, _, err := api.CreateRemoteThread(handle, MakeInheritSa(), 0, procAddress, parameterRemoteAddress, 0)
	if err != nil {
		return fmt.Errorf("fail to create remote thread, because %s", err.Error())
	}
	defer windows.CloseHandle(thread)

	return nil
}

func GetRemoteProcAddress(handle windows.Handle, dllPath string, procName string) (uintptr, error) {
	baseName := filepath.Base(dllPath)
	module, err := modules.Find(handle, baseName)
	if err != nil {
		return 0, err
	}

	if module == nil {
		return 0, fmt.Errorf("unable to find module [ %s ] in process with handle [ 0x%X ]", baseName, handle)
	}

	procOffset, err := GetProcAddressOffsetFrom(dllPath, procName)
	if err != nil {
		return 0, err
	}

	procAddress := uintptr(module.Handle) + procOffset
	return procAddress, nil
}

func GetProcAddressOffsetFrom(dllPath string, procName string) (uintptr, error) {
	handle, err := windows.LoadLibrary(dllPath)
	if err != nil {
		return 0, err
	}

	procPtr, err := windows.GetProcAddress(handle, procName)
	if err != nil {
		log.Println(dllPath, "error")
		return 0, err
	}

	offset := procPtr - uintptr(handle)

	err = windows.FreeLibrary(handle)
	if err != nil {
		return 0, err
	}

	return offset, nil
}
