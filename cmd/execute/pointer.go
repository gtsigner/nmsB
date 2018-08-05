package execute

import (
	"log"
	"time"
	"bytes"
    "encoding/binary"
	"encoding/hex"
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
		log.Printf("Pointer: 0x%X , Value: %d", &value, value)
		time.Sleep(time.Second * 1)
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

	log.Printf("Value: %d", handle, pointer, value)

	return nil
}

func parseAddress(s string) (uintptr, error) {
    data, err := hex.DecodeString(s)
    if err != nil {
        return uintptr(0), err
    }

	var address uint32
    buf := bytes.NewReader(data)
    err = binary.Read(buf, binary.BigEndian, &address)

    return uintptr(address), err
}
