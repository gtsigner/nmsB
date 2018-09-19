package memory

import (
	"../api"
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/windows"
	"unsafe"
)

func readBytes(byteArray []byte, data interface{}) error {
	buf := bytes.NewBuffer(byteArray)
	// BigEndian || LittleEndian
	err := binary.Read(buf, binary.LittleEndian, data)
	return err
}

func ReadProcessMemoryInt32(handle windows.Handle, address uintptr) (int32, error) {
	var i int32
	size := unsafe.Sizeof(i)

	data, err := api.ReadProcessMemory(handle, address, uint64(size))
	if err != nil {
		return 0, err
	}

	err = readBytes(data, &i)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ReadProcessMemoryInt64(handle windows.Handle, address uintptr) (int64, error) {
	var i int64
	size := unsafe.Sizeof(i)

	data, err := api.ReadProcessMemory(handle, address, uint64(size))
	if err != nil {
		return 0, err
	}

	err = readBytes(data, &i)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ReadProcessMemoryFloat32(handle windows.Handle, address uintptr) (float32, error) {
	var f float32
	size := unsafe.Sizeof(f)

	data, err := api.ReadProcessMemory(handle, address, uint64(size))
	if err != nil {
		return 0, err
	}

	err = readBytes(data, &f)
	if err != nil {
		return 0.0, err
	}

	return f, nil
}

func ReadProcessMemoryFloat64(handle windows.Handle, address uintptr) (float64, error) {
	var f float64
	size := unsafe.Sizeof(f)

	data, err := api.ReadProcessMemory(handle, address, uint64(size))
	if err != nil {
		return 0, err
	}

	err = readBytes(data, &f)
	if err != nil {
		return 0.0, err
	}

	return f, nil
}

func ReadProcessMemoryByte(handle windows.Handle, address uintptr) (byte, error) {
	var f byte
	size := unsafe.Sizeof(f)

	data, err := api.ReadProcessMemory(handle, address, uint64(size))
	if err != nil {
		return 0, err
	}

	err = readBytes(data, &f)
	if err != nil {
		return 0, err
	}

	return f, nil
}

func ReadProcessMemoryBytes(handle windows.Handle, address uintptr, size uint64) ([]byte, error) {
	data, err := api.ReadProcessMemory(handle, address, size)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadProcessMemoryString(handle windows.Handle, address uintptr, size uint64) (string, error) {
	stringLength := size + 1
	data, err := api.ReadProcessMemory(handle, address, stringLength)
	if err != nil {
		return "", err
	}
	return string(data[:stringLength]), nil
}
