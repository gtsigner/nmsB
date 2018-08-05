package memory

import (
	"../api"
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/windows"
	"log"
	"unsafe"
)

func ReadProcessMemoryInt32(handle windows.Handle, address uintptr) (int, error) {
	var i int32
	size := unsafe.Sizeof(i)

	log.Println(size)

	data, err := api.ReadProcessMemory(handle, address, int32(size))
	if err != nil {
		return 0, err
	}

	log.Printf("%v\n", data)

	buf := bytes.NewBuffer(data)
	err = binary.Read(buf, binary.BigEndian, &i)
	if err != nil {
		return 0, err
	}

	log.Printf("%v\n", i)

	return int(i), nil
}
