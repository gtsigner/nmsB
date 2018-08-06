package execute

import (
	"log"
	"time"
	"unsafe"
	"strconv"
	"../../win/memory"
	"../../win/process"
)

func Pointer(processId *int, address *string) error {
	if *processId < 0 || *address == "" {
		err := pointerLoop()
		return err
	}

	err := readAddress(processId, address)
	return err
}

func pointerLoop() error {
	value := int32(0)
	for {
		value += 1
		pointer := uintptr(unsafe.Pointer(&value))
		log.Printf("Pointer: 0x%X , Value: %d", pointer, value)
		time.Sleep(time.Second * 10)
	}
}

func readAddress(processId *int, address *string) error {
	// open given process
	handle, err := process.Open(uint(*processId))
	if err != nil {
		return err
	}

	pointer, err := parseAddress(*address)

	if err != nil {
		return err
	}

	log.Printf("Handler: %d, Pointer 0x%X", handle, pointer)
	
	value, err := memory.ReadProcessMemoryInt32(handle, pointer)
	if err != nil {
		return err
	}

	log.Printf("Value: %d, Handle: %d, Value: %d", handle, pointer, value)

	return nil
}

func parseAddress(s string) (uintptr, error) {
	address, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return uintptr(0), err
	}
    return uintptr(address), nil
}
