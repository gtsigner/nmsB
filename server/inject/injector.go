package inject

import (
	"fmt"
	"log"
	"sync"

	"../../config"
	"../../sys/win/dll"
	"../../sys/win/modules"
	"golang.org/x/sys/windows"
)

type Injector struct {
	lock       sync.Mutex
	config     *config.Config
	procOffset uintptr
}

func (injector *Injector) Init() error {
	// lock the injector
	injector.lock.Lock()
	defer injector.lock.Unlock()

	// check for server config
	if injector.config.Server == nil {
		return fmt.Errorf("unable to initialize injector, because server config missing")
	}

	// caluclate the init proc offset for remote init call
	initProcOffset, err := procOffset(injector.config.Server)
	if err != nil {
		return err
	}
	injector.procOffset = initProcOffset

	log.Printf("successful initialized injector and found procOffset [ 0x%X ]", initProcOffset)

	return nil
}

func (injector *Injector) IsInjected() (bool, error) {
	// lock the injector
	injector.lock.Lock()
	defer injector.lock.Unlock()

	// open the process handle
	handle, err := openProcessHandle(injector.config.Server)
	if err != nil {
		return false, err
	}
	defer windows.CloseHandle(handle)

	// check if dll is loaded
	isLoaded, err := injector.isDllLoaded(handle)
	return isLoaded, err
}

func (injector *Injector) isDllLoaded(handle windows.Handle) (bool, error) {
	module, err := injector.findModule(handle)
	if err != nil {
		return false, err
	}
	return module != nil, nil
}

func (injector *Injector) findModule(handle windows.Handle) (*modules.Module, error) {
	// get the name of the module
	baseName, err := moduleName(injector.config.Server)
	if err != nil {
		return nil, err
	}

	// find the module in the remote process
	module, err := modules.Find(handle, baseName)
	return module, err
}

func (injector *Injector) Inject() error {
	// lock the injector
	injector.lock.Lock()
	defer injector.lock.Unlock()

	// open the process handle
	handle, err := openProcessHandle(injector.config.Server)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)

	// check if dll is loaded
	isLoaded, err := injector.isDllLoaded(handle)
	if err != nil {
		return err
	}

	// check if the dll is already loaded
	if isLoaded {
		return nil
	}

	// get the dll path
	dllPath, err := dllPath(injector.config.Server)
	if err != nil {
		return err
	}

	// inject the dll to the target process
	moduleAddress, err := dll.InjectDll(handle, dllPath)
	if err != nil {
		return err
	}

	log.Printf("successful injected [ %s ] on target process [ 0x%X ] at module-address [ 0x%X ]", dllPath, handle, moduleAddress)

	return nil
}

func (injector *Injector) CallInitProc() error {
	// lock the injector
	injector.lock.Lock()
	defer injector.lock.Unlock()

	// open the process handle
	handle, err := openProcessHandle(injector.config.Server)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)

	// convert the config to json form the injected dll
	parameter, err := configToJSON(injector.config)
	if err != nil {
		return err
	}

	// find the dll module in remote process
	module, err := injector.findModule(handle)
	if err != nil {
		return err
	}

	// check if the module was found
	if module == nil {
		dllPath, err := dllPath(injector.config.Server)
		if err != nil {
			return err
		}
		return fmt.Errorf("unable to find module [ %s ] in target process [ 0x%X ], please inject first the module", dllPath, handle)
	}

	// calculate the init proc address
	initProcAddress := uintptr(module.Handle) + injector.procOffset

	// call the remote proc
	err = dll.CallRemoteProc(handle, initProcAddress, parameter)
	if err != nil {
		return nil
	}

	log.Printf("successful called remote proc at [ 0x%X ] on target process [ 0x%X ]", initProcAddress, handle)
	return nil
}
