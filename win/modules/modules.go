package modules

import (
	"../api"
	"golang.org/x/sys/windows"
)

type Module struct {
	Handle windows.Handle
	Name   string
}

func Get(handle windows.Handle) ([]Module, error) {
	// get the amount of modules
	n, err := api.EnumProcessModules(handle, nil)
	if err != nil {
		return nil, err
	}

	// get all the modules
	moduleHandles := make([]windows.Handle, n)
	_, err = api.EnumProcessModules(handle, moduleHandles)
	if err != nil {
		return nil, err
	}

	modules := make([]Module, 0)
	// buffer for length of the module name
	buffer := make([]uint16, 255)
	for _, module := range moduleHandles {
		// get the module name
		n, err := api.GetModuleBaseName(handle, module, &buffer[0], uint32(len(buffer)))
		if err != nil {
			return nil, err
		}
		name := windows.UTF16ToString(buffer[:n])
		// append the module
		modules = append(modules, Module{
			Handle: module,
			Name:   name,
		})
	}

	return modules, nil
}

func Find(handle windows.Handle, name string)(*Module, error){
	modules, err := Get(handle)
	if err != nil {
		return nil, err
	}
	
	for _, module := range modules {
		if module.Name == name {
			return &module, nil
		}
	}

	return nil, nil
}
