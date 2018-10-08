package signature

import (
	"testing"
	"unsafe"
)

var (
	bytes = []byte{0xF1, 0x3, 0x54, 0xFC, 0x34, 0xA2, 0xB1, 0x0A, 0x12}
)

func assertOffset(t *testing.T, sigString string, offset int64) {
	p, err := FromString(sigString)
	if err != nil {
		t.Fatalf(err.Error())
	}

	length := int64(len(bytes))
	start := uintptr(unsafe.Pointer(&bytes[0]))
	foundOffset := Scan(start, length, p)
	if offset != foundOffset {
		t.Fatalf("scan returned wrong offset [ %d != %d ]", offset, foundOffset)
	}
}

func TestSimpleScan(t *testing.T) {
	assertOffset(t, "54 ?? 34", 2)
}

func TestSimpleScan2(t *testing.T) {
	assertOffset(t, "?? 0A", 6)
}

func TestEndBytes(t *testing.T) {
	assertOffset(t, "12", 8)
}

func TestEndBytes2(t *testing.T) {
	assertOffset(t, "?? 12", 7)
}

func TestSTartBytes(t *testing.T) {
	assertOffset(t, "?? 03", 0)
}
