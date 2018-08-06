package memory

import (
	"bytes"
	"testing"
	"unsafe"
	"math/rand"
	"time"
)

func TestReadProcessMemoryByte(t *testing.T) {	
	handle, ok := OpenProcess(t)
	if !ok {
		return
	}

	value := byte(42)
	pointer := uintptr(unsafe.Pointer(&value))
	value2, err := ReadProcessMemoryByte(handle, pointer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Logf("handle: %d, value: %d, pointer: %x, value2: %d\n", handle, value, pointer, value2)

	if value != value2 {
		t.Fatalf("fail to read value byte, because %d != %d \n", value, value2)
		return
	}

}

func TestReadProcessMemoryBytes(t *testing.T) {	
	handle, ok := OpenProcess(t)
	if !ok {
		return
	}

	byteArray := make([]byte, 42)
	rand.Seed(time.Now().UnixNano())
	rand.Read(byteArray)

	pointer := uintptr(unsafe.Pointer(&byteArray[0]))
	byteArray2, err := ReadProcessMemoryBytes(handle, pointer, uint64(len(byteArray)))
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	t.Logf("handle: %d, bytes: %d, pointer: %x, bytes2: %d\n", handle, byteArray, pointer, byteArray2)

	if !bytes.Equal(byteArray, byteArray2) {
		t.Fatalf("fail to read value bytes, because %v != %v \n", byteArray, byteArray2)
		return
	}

}
