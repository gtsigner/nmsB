package dll

import (
	"testing"

	"../process"

	"golang.org/x/sys/windows"
)

func TestGetRemoteProcAddress(t *testing.T) {
	user32, err := windows.LoadLibrary("user32.dll")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	messageBox, err := windows.GetProcAddress(user32, "MessageBoxW")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	handle, err := process.OpenCurrent()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	procAddress, err := GetRemoteProcAddress(handle, "user32.dll", "MessageBoxW")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if procAddress == 0 {
		t.Fatalf("procAddress is NIL")
		return
	}

	if procAddress != messageBox {
		t.Fatalf("procAddress not equal [ 0x%X != 0x%X ]", messageBox, procAddress)
	}
}

func TestGetProcAddressOffsetFrom(t *testing.T) {
	user32, err := windows.LoadLibrary("user32.dll")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	messageBox, err := windows.GetProcAddress(user32, "MessageBoxW")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	offset := uintptr(messageBox) - uintptr(user32)

	offset2, err := GetProcAddressOffsetFrom("user32.dll", "MessageBoxW")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if offset != offset2 {
		t.Fatalf("offsets not equal [ 0x%X != 0x%X ]", offset, offset2)
	}

}
