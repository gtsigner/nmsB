package memory

import (
	"../win/memory"
)

func (reader *MemoryReader) ReadInt32(address uintptr) (int32, error) {
	i, err := memory.ReadProcessMemoryInt32(reader.handle, address)
	return i, err
}

func (reader *MemoryReader) ReadInt64(address uintptr) (int64, error) {
	i, err := memory.ReadProcessMemoryInt64(reader.handle, address)
	return i, err
}
