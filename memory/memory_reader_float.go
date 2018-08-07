package memory

import (
	"../win/memory"
)

func (reader *MemoryReader) ReadFloat32(address uintptr) (float32, error) {
	f, err := memory.ReadProcessMemoryFloat32(reader.handle, address)
	return f, err
}

func (reader *MemoryReader) ReadFloat64(address uintptr) (float64, error) {
	f, err := memory.ReadProcessMemoryFloat64(reader.handle, address)
	return f, err
}
