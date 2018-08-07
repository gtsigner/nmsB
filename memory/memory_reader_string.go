package memory

import (
	"../win/memory"
)

func (reader *MemoryReader) ReadString(address uintptr, length int) (string, error) {
	s, err := memory.ReadProcessMemoryString(reader.handle, address, uint64(length))
	return s, err
}
