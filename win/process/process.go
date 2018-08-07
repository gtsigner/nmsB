package process

import (
	"../api"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type WinProcess struct {
	Id             uint
	ParentId       uint
	Threads        uint
	ThreadPriority uint
	Name           string
}

func PS() ([]WinProcess, error) {
	// open the snapshot handler
	hSnapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}

	// create an empty ProcessEntry32
	var pEntry32 windows.ProcessEntry32
	pEntry32.Size = uint32(unsafe.Sizeof(pEntry32))

	// get the first process
	err = windows.Process32First(hSnapshot, &pEntry32)
	if err != nil {
		return nil, err
	}

	// create the list with processes
	processes := make([]WinProcess, 0)

	for {
		// append a new process
		processes = append(processes, WinProcess{
			Id:             uint(pEntry32.ProcessID),
			ParentId:       uint(pEntry32.ParentProcessID),
			Threads:        uint(pEntry32.Threads),
			ThreadPriority: uint(pEntry32.PriClassBase),
			Name:           syscall.UTF16ToString(pEntry32.ExeFile[:windows.MAX_PATH]),
		})

		// try to get the next process
		err = windows.Process32Next(hSnapshot, &pEntry32)
		if err != nil {
			// check if no more processes available
			if err == windows.ERROR_NO_MORE_FILES {
				return processes, nil
			}
			return nil, err
		}
	}
}

func FindProcess(name string) (*WinProcess, error) {
	processes, err := PS()
	if err != nil {
		return nil, err
	}

	for _, process := range processes {
		if process.Name == name {
			return &process, nil
		}
	}
	return nil, nil
}

func Open(id uint) (windows.Handle, error) {
	handle, err := windows.OpenProcess(api.PROCESS_ALL_ACCESS, false, uint32(id))
	return handle, err
}

func Close(handle windows.Handle) error {
	err := windows.CloseHandle(handle)
	return err
}
