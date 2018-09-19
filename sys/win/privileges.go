package win

import (
	"bytes"
	"encoding/binary"
	"syscall"

	"./api"
	"golang.org/x/sys/windows"
)

type Privilege struct {
	Name       string
	Luid       uint64
	Attributes uint32
}

func LookupPrivilegeName(systemName string, luid uint64) (string, error) {
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

func LookupPrivilegeValue(systemName string, name string) (uint64, error) {
	systemNamePtr, err := syscall.UTF16PtrFromString(systemName)
	if err != nil {
		return 0, nil
	}

	namePtr, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, nil
	}

	var luid uint64
	err = api.LookupPrivilegeValue(systemNamePtr, namePtr, &luid)
	if err != nil {
		return 0, err
	}

	return luid, err
}

func GetPrivilege(token windows.Token, name string) (*Privilege, error) {
	privileges, err := GetPrivileges(token)
	if err != nil {
		return nil, err
	}

	for _, privilege := range privileges {
		if privilege.Name == name {
			return &privilege, nil
		}
	}

	return nil, nil
}

func GetCurrentProcessPrivilege(name string) (*Privilege, error) {
	token, err := windows.OpenCurrentProcessToken()
	if err != nil {
		return nil, err
	}

	privilege, err := GetPrivilege(token, name)
	return privilege, err
}

func GetCurrentProcessPrivileges() ([]Privilege, error) {
	token, err := windows.OpenCurrentProcessToken()
	if err != nil {
		return nil, err
	}

	privileges, err := GetPrivileges(token)
	if err != nil {
		return nil, err
	}

	err = token.Close()
	if err != nil {
		return nil, err
	}

	return privileges, err
}

func GetPrivileges(token windows.Token) ([]Privilege, error) {
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

	privileges := make([]Privilege, privilegeCount)

	for i := 0; i < int(privilegeCount); i++ {
		var luid uint64
		err = binary.Read(buffer, binary.LittleEndian, &luid)
		if err != nil {
			return nil, err
		}

		var attributes uint32
		err = binary.Read(buffer, binary.LittleEndian, &attributes)
		if err != nil {
			return nil, err
		}

		name, err := LookupPrivilegeName("", luid)
		if err != nil {
			return nil, err
		}

		privileges[i] = Privilege{
			Luid:       luid,
			Name:       name,
			Attributes: attributes,
		}
	}

	return privileges, nil
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
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, uint32(len(privileges)))

	for _, privilege := range privileges {
		luid, err := LookupPrivilegeValue("", privilege)
		if err != nil {
			return err
		}

		err = windows.GetLastError()
		if err != nil {
			return err
		}

		binary.Write(&buffer, binary.LittleEndian, luid)

		if enabled {
			binary.Write(&buffer, binary.LittleEndian, uint32(api.SE_PRIVILEGE_ENABLED))
		} else {
			binary.Write(&buffer, binary.LittleEndian, uint32(0))
		}
	}

	_, err := api.AdjustTokenPrivileges(token, false, &buffer.Bytes()[0], uint32(buffer.Len()), nil, nil)
	if err != nil {
		return err
	}

	err = windows.GetLastError()
	if err != nil {
		return err
	}

	return err
}
