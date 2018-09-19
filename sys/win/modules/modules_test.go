package modules

import (
	"../process"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
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

	modules, err := Get(handle)

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if modules == nil {
		t.Fatal("modules for the process go.exe are nil")
		return
	}

	if len(modules) < 1 {
		t.Fatal("modules for the process go.exe are empty")
		return
	}

	x := 2
	value := 2 + x

	for _, module := range modules {
		log.Printf("%s: 0x%X, %X, %d", module.Name, module.Handle, &value, uint32(module.Handle)-uint32(uintptr(value)))
	}

}
