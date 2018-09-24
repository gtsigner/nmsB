package config

import (
	"fmt"

	"path/filepath"

	"../utils"
)

var (
	fileEndings = []string{"json", "yml", "yaml"}
	fileNames   = []string{"config", "config.dev", "config.local"}
)

func inspectConfigDirectory(directory string) ([]string, error) {
	// check if the directory exists
	exists, err := utils.FileExists(directory)
	if err != nil {
		return nil, err
	}

	// check if the directory exists
	if !exists {
		return nil, nil
	}

	var configFiles []string
	for _, fileName := range fileNames {
		// find all configFile for the given name
		fileNameConfigs, err := inspectConfigFile(directory, fileName)
		if err != nil {
			return nil, err
		}

		// append found config files
		if fileNameConfigs != nil && len(fileNameConfigs) > 0 {
			configFiles = append(configFiles, fileNameConfigs...)
		}
	}

	return configFiles, nil
}

func inspectConfigFile(directory string, fileName string) ([]string, error) {
	var configFiles []string
	for _, ending := range fileEndings {
		// create the fincig file for the ending
		configFile, err := toConfigFile(directory, fileName, ending)
		if err != nil {
			return nil, err
		}
		// check if the config file exists
		configExists, err := utils.FileExists(configFile)
		if err != nil {
			return nil, err
		}
		// add the config if exists
		if configExists {
			configFiles = append(configFiles, configFile)
		}
	}
	return configFiles, nil
}

func toConfigFile(directory string, fileName string, ending string) (string, error) {
	configFileName := fmt.Sprintf("%s.%s", fileName, ending)
	configFile := filepath.Join(directory, configFileName)
	absConfigPath, err := filepath.Abs(configFile)
	return absConfigPath, err
}

func configFiles() ([]string, error) {
	// serach for all config directires
	directories, err := configDirectories()
	if err != nil {
		return nil, err
	}

	// check if some directories found
	if directories == nil || len(directories) > 0 {
		return nil, nil
	}

	var configFiles []string
	// inspect all directories
	for _, directory := range directories {
		// inspect and get all configs for directory
		directoryConfigFiles, err := inspectConfigDirectory(directory)
		if err != nil {
			return nil, err
		}
		// append if config files available
		if directoryConfigFiles != nil && len(directoryConfigFiles) > 0 {
			configFiles = append(configFiles, directoryConfigFiles...)
		}
	}

	return configFiles, nil
}
