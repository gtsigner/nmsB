package memory

import (
	"../win/process"
	"fmt"
	"golang.org/x/sys/windows"
)

type MemoryReader struct {
	PId     uint
	handle  windows.Handle
	modules map[string]windows.Handle
}

func NewMemoryReader() *MemoryReader {
	return &MemoryReader{}
}

func (reader *MemoryReader) Open(pid uint) error {
	// set the process id
	reader.PId = pid
	// open the process
	handle, err := process.Open(pid)
	if err != nil {
		return err
	}
	// set the handle to the process
	reader.handle = handle

	// load the modules
	err = reader.loadModules()
	return err
}

func (reader *MemoryReader) OpenByName(name string) error {
	p, err := process.FindProcess(name)
	if err != nil {
		return err
	}

	if p == nil {
		return fmt.Errorf("unable to find process with name [ %s ]", name)
	}
	// open the found process
	err = reader.Open(p.Id)
	return err
}

func (reader *MemoryReader) Close() error {
	err := process.Close(reader.handle)
	return err
}
