package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"../win"
	"../win/dll"
	"../win/process"
)

var (
	DLL_FILE            = "../dist/nmsB-windows-amd64.dll"
	TARGET_PROCESS_NAME = "notepad.exe"
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

func TestDllAsInject(t *testing.T) {
	dllPath, err := filepath.Abs(DLL_FILE)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	exists, err := Exists(dllPath)
	if !exists {
		t.Fatalf("unable to find dll [ %s ], because file not exists", dllPath)
		return
	}

	err = win.EnableDebugPrivilege()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	p, err := process.FindProcess(TARGET_PROCESS_NAME)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if p == nil {
		t.Fatalf("unable to find target process with name [ %s ]", TARGET_PROCESS_NAME)
		return
	}

	handle, err := process.Open(p.Id)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	dllHandle, err := dll.InjectDll(handle, dllPath)
	if err != nil {
		t.Error(err)
		return
	}

	log.Printf("0x%X", dllHandle)

	now := time.Now()
	parameter := fmt.Sprintf("{\"port\":4567, \"host\":\"test\", \"date\": \"%s\"}", now.String())
	err = dll.CallRemoteProc(handle, dllPath, "Init", parameter)
	if err != nil {
		t.Error(err)
		return
	}

	//time.Sleep(time.Second * 100)
}
