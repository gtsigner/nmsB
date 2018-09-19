package dll

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
	"log"
	"path/filepath"
	"unsafe"

	"../api"
	"../modules"

	"golang.org/x/sys/windows"
)

func CallRemoteProc(handle windows.Handle, dllPath string, procName string, parameter string) error {
	// get the rmote address of the given dll
	procAddress, err := GetRemoteProcAddress(handle, dllPath, procName)
	if err != nil {
		return err
	}

	// convert the parameter toa CString
	cParameter := C.CString(parameter)
	// free the CString after usage
	defer C.free(unsafe.Pointer(cParameter))
	// count the length of the cStrCStringing
	cParameterLength := uint32(C.strlen(cParameter))

	// allocate memory in remote process
	parameterRemoteAddress, err := api.VirtualAllocEx(handle, 0, cParameterLength, windows.MEM_COMMIT, windows.PAGE_READWRITE)
	if err != nil {
		return fmt.Errorf("fail to allocate memory, because %s", err.Error())
	}

	// write the parameter to remote address
	_, err = api.WriteProcessMemory(handle, parameterRemoteAddress, unsafe.Pointer(cParameter), cParameterLength)
	if err != nil {
		return fmt.Errorf("fail to write parameter data [ %d bytes ] to [ 0x%X ], because %s",
			cParameterLength, parameterRemoteAddress, err.Error())
	}

	// create the thread for the given remote proc with the parameter
	thread, _, err := api.CreateRemoteThread(handle, MakeInheritSa(), 0, procAddress, parameterRemoteAddress, 0)
	if err != nil {
		return fmt.Errorf("fail to create remote thread, because %s", err.Error())
	}
	// close the thread handle
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
