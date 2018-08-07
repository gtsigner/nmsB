package memory

import (
	"testing"
	"unsafe"
)

func TestReadProcessMemoryFloat32(t *testing.T) {
	handle, ok := OpenProcess(t)
	if !ok {
		return
	}

	value := float32(42.42)
	pointer := uintptr(unsafe.Pointer(&value))
	value2, err := ReadProcessMemoryFloat32(handle, pointer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Logf("handle: %d, value: %f, pointer: %x, value2: %f\n", handle, value, pointer, value2)

	if value != value2 {
		t.Fatalf("fail to read value float32, because %f != %f \n", value, value2)
		return
	}

}

func TestReadProcessMemoryFloat64(t *testing.T) {
	handle, ok := OpenProcess(t)
	if !ok {
		return
	}

	value := float64(42.42)
	pointer := uintptr(unsafe.Pointer(&value))
	value2, err := ReadProcessMemoryFloat64(handle, pointer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Logf("handle: %d, value: %f, pointer: %x, value2: %f\n", handle, value, pointer, value2)

	if value != value2 {
		t.Fatalf("fail to read value float64, because %f != %f \n", value, value2)
		return
	}

}
