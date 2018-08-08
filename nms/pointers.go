package nms

import (
	"../memory"
	"./config"
	"fmt"
)

func (instance *Insatnce) Pointer(name string) (uintptr, error) {
	configPtr, ok := instance.config.Pointers[name]
	if !ok {
		return memory.NULL, fmt.Errorf("unable to find config pointer for name [ %s ]", name)
	}

	if (configPtr.Offsets == nil || len(configPtr.Offsets) < 1) && configPtr.Module == nil {
		return memory.NULL, fmt.Errorf("unable process config pointer with name [ %s ], because no module or offsets given", name)
	}

	if configPtr.Module != nil {
		ptr, err := instance.reader.PointerAt(configPtr.Module, configPtr.Offsets...)
		return ptr, err
	}

	ptr, err := instance.reader.PointerAt(configPtr.Offsets...)
	return ptr, err
}
