package inject

import (
	"fmt"
	"path/filepath"

	"../../config"
	"../../message/json"
	"../../sys/win/dll"
	"../../sys/win/process"
	"../../utils"
	"golang.org/x/sys/windows"
)

func procOffset(serverConfig *config.ServerConfig) (uintptr, error) {
	// get the path to the dll
	path, err := dllPath(serverConfig)
	if err != nil {
		return 0, err
	}

	// check if init proc name given
	if serverConfig.InitProcName == nil {
		return 0, fmt.Errorf("missing init proc-name in server config")
	}

	// get the proc name
	procName := *serverConfig.InitProcName
	// calulcate the procOffset
	offset, err := dll.GetProcAddressOffsetFrom(path, procName)
	if err != nil {
		return 0, err
	}
	return offset, nil
}

func moduleName(serverConfig *config.ServerConfig) (string, error) {
	// get the path to the dll
	dllPath, err := dllPath(serverConfig)
	if err != nil {
		return "", err
	}
	baseName := filepath.Base(dllPath)
	return baseName, nil
}

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

func openProcessHandle(serverConfig *config.ServerConfig) (windows.Handle, error) {
	// get the process id of the process-name from the config
	pid, err := processId(serverConfig)
	if err != nil {
		return 0, err
	}

	// open the target process
	handle, err := process.Open(pid)
	if err != nil {
		return 0, err
	}
	return handle, nil
}

func configToJSON(cfg *config.Config) (string, error) {
	data, err := json.Encode(cfg)
	return data, err
}
