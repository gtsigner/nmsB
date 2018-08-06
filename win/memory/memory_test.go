package memory

import (
	"../process"
	"testing"	
	"os"
	"golang.org/x/sys/windows"
)

func OpenProcess(t *testing.T)(windows.Handle,bool) {
	pid := uint(os.Getpid())
	handle, err := process.Open(pid)
	if err != nil {
		t.Errorf(err.Error())
		return 0, false
	}

	if handle == 0 {
		t.Fatalf("fail to open process with pid %d\n", pid)
		return 0, false
	}

	return handle, true
}