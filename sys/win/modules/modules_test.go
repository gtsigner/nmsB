package modules

import (
	"log"
	"testing"

	"../process"
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

	for _, module := range modules {
		log.Printf("%s Base: 0x%X Size: %d", module.Name, module.Handle, module.Size)
	}

}

func TestGetProcessModule(t *testing.T) {
	mod, err := GetProcessModule("")
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("%s: 0x%X", mod.Name, mod.Handle)
}
