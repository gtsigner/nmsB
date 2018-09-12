package win

import (
	"bytes"
	"encoding/binary"
	"log"
	"syscall"
	"unsafe"

	"./api"
	"golang.org/x/sys/windows"
)

func LookupPrivilegeName(systemName string, luid api.LUID) (string, error) {
	systemNamePtr, err := syscall.UTF16PtrFromString(systemName)
	if err != nil {
		return "", nil
	}

	buffer := make([]uint16, 256)
	bufferSize := uint32(len(buffer))
	err = api.LookupPrivilegeName(systemNamePtr, &luid, &buffer[0], &bufferSize)
	if err != nil {
		return "", err
	}
	privilegeName := syscall.UTF16ToString(buffer)
	return privilegeName, nil
}

func LookupPrivilegeValue(systemName string, name string) (*api.LUID, error) {
	systemNamePtr, err := syscall.UTF16PtrFromString(systemName)
	if err != nil {
		return nil, nil
	}

	namePtr, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, nil
	}

	var luid api.LUID
	err = api.LookupPrivilegeValue(systemNamePtr, namePtr, &luid)
	if err != nil {
		return nil, err
	}

	return &luid, err
}

func GetCurrentProcessPrivileges() ([]string, error) {
	token, err := windows.OpenCurrentProcessToken()
	if err != nil {
		return nil, err
	}

	names, err := GetPrivileges(token)
	if err != nil {
		return nil, err
	}

	err = token.Close()
	if err != nil {
		return nil, err
	}

	return names, err
}

func GetPrivileges(token windows.Token) ([]string, error) {
	var n uint32
	err := windows.GetTokenInformation(token, windows.TokenPrivileges, nil, uint32(0), &n)
	if err != nil {
		if err != windows.ERROR_INSUFFICIENT_BUFFER {
			return nil, err
		}
	}

	buffer := bytes.NewBuffer(make([]byte, n))
	err = windows.GetTokenInformation(token, windows.TokenPrivileges, &buffer.Bytes()[0], uint32(buffer.Len()), &n)

	if err != nil {
		return nil, err
	}

	var privilegeCount uint32
	err = binary.Read(buffer, binary.LittleEndian, &privilegeCount)
	if err != nil {
		return nil, err
	}

	names := make([]string, privilegeCount)

	for i := 0; i < int(privilegeCount); i++ {
		var luidAndAttr api.LUID_AND_ATTRIBUTES
		err = binary.Read(buffer, binary.LittleEndian, &luidAndAttr.Luid)
		if err != nil {
			return nil, err
		}

		err = binary.Read(buffer, binary.LittleEndian, &luidAndAttr.Attributes)
		if err != nil {
			return nil, err
		}

		name, err := LookupPrivilegeName("", luidAndAttr.Luid)
		if err != nil {
			return nil, err
		}
		log.Println(name, luidAndAttr.Attributes)
		names[i] = name
	}

	return names, nil
}

func EnableDebugPrivilege() error {
	err := SetCurrentProcessPrivilege(true, api.SE_DEBUG_NAME)
	return err
}

func SetCurrentProcessPrivilege(enabled bool, privileges ...string) error {
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

	err = SetPrivilege(tokenHandle, enabled, privileges...)
	if err != nil {
		return err
	}

	err = tokenHandle.Close()
	return err
}

func SetPrivilege(token windows.Token, enabled bool, privileges ...string) error {
	length := len(privileges)
	tokenPrivileges := api.TOKEN_PRIVILEGES{
		PrivilegeCount: uint32(length),
		Privileges:     make([]api.LUID_AND_ATTRIBUTES, length),
	}

	for index, privilege := range privileges {
		luid, err := LookupPrivilegeValue("", privilege)
		if err != nil {
			return err
		}

		err = windows.GetLastError()
		if err != nil {
			return err
		}

		// https://github.com/elastic/beats/blob/master/vendor/github.com/elastic/gosigar/sys/windows/privileges.go

		log.Println(*luid)

		tokenPrivileges.Privileges[index].Luid = *luid

		if enabled {
			tokenPrivileges.Privileges[index].Attributes = api.SE_PRIVILEGE_ENABLED
		} else {
			tokenPrivileges.Privileges[index].Attributes = 0
		}
	}

	size := uint32(unsafe.Sizeof(tokenPrivileges))
	_, err := api.AdjustTokenPrivileges(token, false, &tokenPrivileges, size, nil, nil)
	if err != nil {
		return err
	}

	err = windows.GetLastError()
	if err != nil {
		return err
	}

	return err
}
