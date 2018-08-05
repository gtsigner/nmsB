package memory

import (
	"../modules"
	"../process"
	"log"
	"testing"
	"unsafe"
)

func TestReadProcessMemoryInt32(t *testing.T) {
	value := int32(42)

	p, err := process.FindProcess("go.exe")

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if p == nil {
		t.Fatal("no process with name go.exe")
		return
	}

	handle, err := process.Open(p.Id)

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if handle == 0 {
		t.Fatal("fail to open process go.exe")
		return
	}

	module, err := modules.Find(handle, "go.exe")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	log.Printf("%d, value: %x, module: %x\n", value, &value, module.Handle)

	v, err := ReadProcessMemoryInt32(handle, uintptr(unsafe.Pointer(&value)))
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	log.Println(value, v)

}
