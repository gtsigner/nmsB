package process

import (
	//"log"
	"testing"
)

func TestPs(t *testing.T) {

	processes, err := PS()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if processes == nil {
		t.Fatal("process list nil")
		return
	}

	if len(processes) < 1 {
		t.Fatal("process list is empty")
		return
	}

	/*for _, process := range processes {
		log.Println(process.Id, process.ParentId, process.Name)
	}*/

}

func TestFindProcess(t *testing.T) {
	process, err := FindProcess("go.exe")

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if process == nil {
		t.Fatal("no process with name go.exe")
		return
	}
}

func TestOpenProcess(t *testing.T) {
	process, err := FindProcess("go.exe")

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if process == nil {
		t.Fatal("no process with name go.exe")
		return
	}

	handle, err := Open(process.Id)

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if handle == 0 {
		t.Fatal("fail to open process go.exe")
		return
	}
}
