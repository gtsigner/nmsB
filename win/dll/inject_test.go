package dll

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"../../win"
	"../process"
)

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func TestInject(t *testing.T) {
	dllPath := "../../dist/nmsB-windows-amd64.dll"

	absPath, err := filepath.Abs(dllPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	exists, err := Exists(absPath)
	if !exists {
		t.Fatalf("unable to find dll [ %s ], because file not exists", absPath)
		return
	}

	err = win.EnableDebugPrivilege()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	handle, err := process.OpenCurrent()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	dllHandle, err := InjectDll(handle, absPath)
	if err != nil {
		t.Error(err)
		return
	}

	log.Printf("0x%X", dllHandle)

	time.Sleep(time.Second * 100)
}
