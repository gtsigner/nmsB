package asm

import (
	"bytes"
	"encoding/binary"
	"log"
	"unsafe"

	"../api"
	"../process"
	"golang.org/x/sys/windows"
)

const (
	NOP_VALUE    = byte(0x90)
	RETURN_VALUE = byte(0xC3)
	CALL_VALUE   = byte(0xE8)
	JUMP_VALUE   = byte(0xE9)
)

func writeProtection(process windows.Handle, address uintptr, size uint32) (uint32, error) {
	var oldProtect uint32
	err := api.VirtualProtectEx(process, address, size, windows.PAGE_EXECUTE_READWRITE, &oldProtect)
	if err != nil {
		return 0, err
	}
	return oldProtect, nil
}

func resetProtection(process windows.Handle, protect uint32, address uintptr, size uint32) error {
	err := api.VirtualProtectEx(process, address, size, protect, &protect)
	return err
}

func writeByte(address uintptr, b byte) error {
	err := writeBytes(address, []byte{b})
	return err
}

func writeBytes(address uintptr, bytes []byte) error {
	handle, err := process.OpenCurrent()
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)

	size := uint32(unsafe.Sizeof(bytes[0]) * uintptr(len(bytes)))
	protect, err := writeProtection(handle, address, size)
	if err != nil {
		return err
	}

	_, err = api.WriteProcessMemory(handle, address, unsafe.Pointer(&bytes[0]), size)
	if err != nil {
		return err
	}

	err = resetProtection(handle, protect, address, size)
	return err
}

func Nop(address uintptr) error {
	err := writeByte(address, NOP_VALUE)
	return err
}

func Return(address uintptr) error {
	err := writeByte(address, RETURN_VALUE)
	return err
}

func Call(address uintptr, funcPtr uintptr) error {

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint64(funcPtr))
	if err != nil {
		return err
	}

	b := append([]byte{byte(0xFF)}, buf.Bytes()...)
	err = writeBytes(address, b)

	return err
}

func MovEAX(address uintptr, value uint64) (int, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		return 0, err
	}

	b := append([]byte{byte(0x8F), byte(0xC0 | 0)}, buf.Bytes()...)
	err = writeBytes(address, b)
	return len(b), err
}

func Jump(address uintptr, jmpPtr uintptr) error {
	buf := new(bytes.Buffer)
	size := unsafe.Sizeof(jmpPtr)
	log.Println(size)
	err := binary.Write(buf, binary.BigEndian, uint64(jmpPtr))
	if err != nil {
		return err
	}

	b := append([]byte{byte(0xFF), byte(0x25)}, buf.Bytes()...)
	err = writeBytes(address, b)
	return err
}
