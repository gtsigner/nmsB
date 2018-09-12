package win

import (
	"syscall"

	"./api"
	"golang.org/x/sys/windows"
)

func EnableDebugPrivilege() error {
	handle, err := windows.GetCurrentProcess()
	if err != nil {
		return err
	}

	var tokenHandle windows.Token
	err = windows.OpenProcessToken(handle, windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &tokenHandle)
	if err != nil {
		return err
	}

	err = windows.GetLastError()
	if err != nil {
		return err
	}

	privilegeName, err := syscall.UTF16PtrFromString(api.SE_DEBUG_NAME)
	if err != nil {
		return err
	}

	var tokenPrivileges api.TOKEN_PRIVILEGES
	err = api.LookupPrivilegeValue(nil, privilegeName, &tokenPrivileges.Privileges[0].Luid)
	if err != nil {
		return err
	}

	err = windows.GetLastError()
	if err != nil {
		return err
	}

	tokenPrivileges.PrivilegeCount = 1
	tokenPrivileges.Privileges[0].Attributes = api.SE_PRIVILEGE_ENABLED

	_, err = api.AdjustTokenPrivileges(tokenHandle, false, &tokenPrivileges, 0, nil, nil)
	if err != nil {
		return err
	}

	err = windows.GetLastError()
	if err != nil {
		return err
	}

	err = tokenHandle.Close()
	return err

}
