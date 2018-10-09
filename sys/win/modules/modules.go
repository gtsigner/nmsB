package modules

import (
	"../api"
	"golang.org/x/sys/windows"
)

type Module struct {
	Handle     windows.Handle
	Name       string
	Size       uint32
	EntryPoint uintptr
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
	for _, moduleHandle := range moduleHandles {
		mod, err := module(handle, moduleHandle)
		if err != nil {
			return nil, err
		}
		modules = append(modules, *mod)
	}

	return modules, nil
}

func Find(handle windows.Handle, name string) (*Module, error) {
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

func module(handle windows.Handle, module windows.Handle) (*Module, error) {
	buffer := make([]uint16, 255)
	// get the module name
	n, err := api.GetModuleBaseName(handle, module, &buffer[0], uint32(len(buffer)))
	if err != nil {
		return nil, err
	}
	moduleName := windows.UTF16ToString(buffer[:n])

	// get the module information
	info, err := api.GetModuleInformation(handle, module)
	if err != nil {
		return nil, err
	}

	return &Module{
		Handle:     module,
		Name:       moduleName,
		Size:       uint32(info.SizeOfImage),
		EntryPoint: uintptr(info.EntryPoint),
	}, nil
}

func GetModule(handle windows.Handle, name string) (*Module, error) {
	moduleHandle, err := api.GetModuleHandle(name)
	if err != nil {
		return nil, err
	}

	mod, err := module(handle, moduleHandle)
	return mod, err
}

func GetProcessModule(name string) (*Module, error) {
	handle, err := windows.GetCurrentProcess()
	if err != nil {
		return nil, err
	}
	mod, err := GetModule(handle, name)
	return mod, err
}
