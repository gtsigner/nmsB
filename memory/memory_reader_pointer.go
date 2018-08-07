package memory

import (
	"fmt"
	"log"
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
		log.Printf("address: 0x%X, value: %v\n", address, value)
		switch v := value.(type) {
		// if type is string, find the module
		case string:
			moduleBase := reader.ModuleBase(v)
			if moduleBase == NULL {
				return NULL, fmt.Errorf("unable to find module base [ %s ] in process with id [ %d ]", v, reader.PId)
			}
			log.Printf("string: %s, moduleBase: 0x%X", v, moduleBase)
			address = moduleBase
		// if type uintptr add the offset and read ptr
		case uintptr:
			if address == NULL {
				address = v
				log.Printf("uintptr: 0x%X", v)
			} else {
				log.Printf("ptr: 0x%X", address+v)
				ptr, err := reader.ReadPtr(address + v)

				if err != nil {
					return NULL, err
				}
				i, _ := reader.ReadInt64(ptr)
				log.Printf("uintptr: 0x%X, ptr: 0x%X, value: %d, i: %d", v, ptr, ptr, i)
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
			log.Println("default")
		}

	}
	return address, nil
}
