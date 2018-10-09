package signature

import (
	"log"
	"reflect"
	"testing"
)

var (
	data = []byte{0x34, 0xFA, 0xCF, 0x1F}
)

func _TestDoIt() {
	x := float64(12)
	y := float64(42)
	z := x + y
	log.Printf("%f", z)
}

func TestScanModule(t *testing.T) {
	sigString := "F2 0F 10 05 ?? ?? ?? ?? F2 0F 11 44 24 ?? 0F 57"
	ptr, err := ScanModule(sigString, "")
	if err != nil {
		t.Error(err)
		return
	}

	fnPtr := uintptr(reflect.ValueOf(_TestDoIt).Pointer())
	if (fnPtr + uintptr(0x28)) != ptr {
		t.Fatalf("unable to find sig, because [ 0x%X != 0x%X + 0x28 ]", ptr, fnPtr)
	}
}
