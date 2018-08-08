package memory

import (
	"fmt"
)

func (reader *MemoryReader) ReadPtr(address uintptr) (uintptr, error) {
	value, err := reader.ReadInt64(address)
	if err != nil {
		return NULL, err
	}
	return uintptr(value), nil
}

func (reader *MemoryReader) PointerAt(a ...interface{}) (uintptr, error) {
	address := NULL
	for _, value := range a {
		// check the type
		switch v := value.(type) {
		// if type is string, find the module
		case string:
			moduleBase := reader.ModuleBase(v)
			if moduleBase == NULL {
				return NULL, fmt.Errorf("unable to find module base [ %s ] in process with id [ %d ]", v, reader.PId)
			}
			address = moduleBase
		// if type uintptr add the offset and read ptr
		case uintptr:
			if address == NULL {
				address = v
			} else {
				ptr, err := reader.ReadPtr(address + v)
				if err != nil {
					return NULL, err
				}
				address = ptr
			}
		// if type uint64 add the offset
		case uint64:
			if address == NULL {
				address = uintptr(v)
			} else {
				address = address + uintptr(v)
			}
		default:
			return NULL, fmt.Errorf("unknown pointer at type for value [ %v ]", v)
		}

	}
	return address, nil
}
