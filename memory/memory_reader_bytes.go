package memory

import (
	"../win/memory"
)

func (reader *MemoryReader) ReadByte(address uintptr) (byte, error) {
	b, err := memory.ReadProcessMemoryByte(reader.handle, address)
	return b, err
}

func (reader *MemoryReader) ReadBytes(address uintptr, size int) ([]byte, error) {
	bytes, err := memory.ReadProcessMemoryBytes(reader.handle, address, uint64(size))
	return bytes, err
}
