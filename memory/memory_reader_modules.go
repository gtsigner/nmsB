package memory

import (
	"../win/modules"
	"fmt"
	"golang.org/x/sys/windows"
)

func (reader *MemoryReader) loadModules() error {
	// create the modules for the process
	reader.modules = make(map[string]windows.Handle)
	// load all modules
	mods, err := modules.Get(reader.handle)
	if err != nil {
		return err
	}
	// chekc if modules not nil
	if mods == nil {
		return fmt.Errorf("unable to load modules, because process with pid [ %d ] has no modules", reader.PId)
	}
	// set all modules for the reader
	for _, module := range mods {
		reader.modules[module.Name] = module.Handle
	}
	return nil
}

func (reader *MemoryReader) ModuleBase(name string) uintptr {
	handle, ok := reader.modules[name]
	if !ok {
		return NULL
	}
	return uintptr(handle)
}
