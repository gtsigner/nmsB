package nms

import (
	"fmt"

	"../memory"
)


func toInterfaces(module *string, array []uintptr) []interface{} {
	indexOffset := 0

	// calculate the length of the interfaces
	length := len(array)
	if module != nil {
		indexOffset = 1
		length = length + 1
	}

	// make the interfaces
	interfaces := make([]interface{}, length)
	// set module as first if given
	if module != nil {
		interfaces[0] = *module
	}

	// set all the pointers
	for i, v := range array {
		interfaces[i+indexOffset] = v
	}
	return interfaces
}

func (instance *Instance) Pointer(name string) (uintptr, error) {
	configPtr, ok := instance.config.Pointers[name]
	if !ok {
		return memory.NULL, fmt.Errorf("unable to find config pointer for name [ %s ]", name)
	}

	if (configPtr.Offsets == nil || len(configPtr.Offsets) < 1) && configPtr.Module == nil {
		return memory.NULL, fmt.Errorf("unable process config pointer with name [ %s ], because no module or offsets given", name)
	}

	ptr, err := instance.reader.PointerAt(toInterfaces(configPtr.Module, configPtr.Offsets)...)
	return ptr, err
}
