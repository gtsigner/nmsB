package memory

import (
	"testing"
	"unsafe"
)

func TestReadProcessMemoryInt32(t *testing.T) {
	handle, ok := OpenProcess(t)
	if !ok {
		return
	}

	value := int32(42)
	pointer := uintptr(unsafe.Pointer(&value))
	value2, err := ReadProcessMemoryInt32(handle, pointer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Logf("handle: %d, value: %d, pointer: %x, value2: %d\n", handle, value, pointer, value2)

	if value != value2 {
		t.Fatalf("fail to read value int32, because %d != %d \n", value, value2)
		return
	}

}

func TestReadProcessMemoryInt64(t *testing.T) {
	handle, ok := OpenProcess(t)
	if !ok {
		return
	}

	value := int64(42)
	pointer := uintptr(unsafe.Pointer(&value))
	value2, err := ReadProcessMemoryInt64(handle, pointer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Logf("handle: %d, value: %d, pointer: %x, value2: %d\n", handle, value, pointer, value2)

	if value != value2 {
		t.Fatalf("fail to read value int64, because %d != %d \n", value, value2)
		return
	}

}
