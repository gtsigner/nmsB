package memory

import (
	"testing"
	"unsafe"
)

func TestReadProcessMemoryString(t *testing.T) {
	handle, ok := OpenProcess(t)
	if !ok {
		return
	}

	value := []byte("Hello World")
	pointer := uintptr(unsafe.Pointer(&value[0]))
	value2, err := ReadProcessMemoryString(handle, pointer, uint64(len(value)))
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Logf("handle: %d, value: %s, pointer: %x, value2: %s\n", handle, value, pointer, value2)

	if string(value) == value2 {
		t.Fatalf("fail to read value string, because %s != %s \n", value, value2)
		return
	}

}
