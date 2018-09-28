package inject

import (
	"fmt"
	"path/filepath"

	"golang.org/x/sys/windows"

	"../../config"
	"../../message/json"
	"../../sys/win"
	"../../sys/win/dll"
	"../../sys/win/process"
	"../../utils"
)

func dllPath(serverConfig *config.ServerConfig) (string, error) {

	// check if an dll path configured
	if serverConfig.DllPath == nil {
		return "", fmt.Errorf("unable to inject dll, because missing dll-path in config")
	}

	// make the dll path absolute
	dllPath, err := filepath.Abs(*serverConfig.DllPath)
	if err != nil {
		return "", err
	}

	dllExists, err := utils.FileExists(dllPath)
	if err != nil {
		return "", err
	}

	if !dllExists {
		return "", fmt.Errorf("unable to inject dll [ %s ], because dll not found", dllPath)
	}

	return dllPath, nil
}

func processId(serverConfig *config.ServerConfig) (uint, error) {
	// check if an process name is configured
	if serverConfig.ProcessName == nil {
		return 0, fmt.Errorf("unable to inject dll, because missing process-name in config")
	}

	p, err := process.FindProcess(*serverConfig.ProcessName)
	if err != nil {
		return 0, err
	}

	if p == nil {
		return 0, fmt.Errorf("unable to inject dll, because no process found with name [ %s ]", *serverConfig.ProcessName)
	}

	return p.Id, nil
}

func configToJSON(cfg *config.Config) (string, error) {
	data, err := json.Encode(cfg)
	return data, err
}

func Inject(cfg *config.Config) error {
	if cfg.Server == nil {
		return fmt.Errorf("unable to inject dll, because server config missing")
	}

	// get the path to the dll from the config
	dllPath, err := dllPath(cfg.Server)
	if err != nil {
		return err
	}

	// get the process id of the process-name from the config
	pid, err := processId(cfg.Server)
	if err != nil {
		return err
	}

	// get the debug privilege for inject
	err = win.EnableDebugPrivilege()
	if err != nil {
		return err
	}

	// open the target process
	handle, err := process.Open(pid)
	if err != nil {
		return err
	}

	// inject the dll to the target process
	dllHandle, err := dll.InjectDll(handle, dllPath)
	if err != nil {
		return err
	}

	// close the dll handle
	// TODO store for uninject
	err = windows.CloseHandle(dllHandle)
	if err != nil {
		return err
	}

	// convert the config to json form the injected dll
	parameter, err := configToJSON(cfg)
	if err != nil {
		return err
	}

	// call the remote init method
	// TODO Init Method from config
	err = dll.CallRemoteProc(handle, dllPath, "Init", parameter)
	if err != nil {
		return err
	}

	return nil
}
