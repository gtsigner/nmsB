package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	configDirectoryName = "config"
)

func currentWorkingDirectory() (string, error) {
	cwdPath, err := filepath.Abs(".")
	return cwdPath, err
}

func userHome() (*string, error) {
	userProfile := os.Getenv("USERPROFILE")
	if userProfile != "" {
		return &userProfile, nil
	}

	userHome := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
	if userHome != "" {
		return &userHome, nil
	}
	return nil, fmt.Errorf("unable to find user home directory")
}

func configDirectories() ([]string, error) {
	var directories []string

	// append ${pwd}
	cwdPath, err := currentWorkingDirectory()
	if err != nil {
		return nil, err
	}
	directories = append(directories, cwdPath)

	// append ${pwd}/config
	cwdConfigDirectory := filepath.Join(cwdPath, configDirectoryName)
	directories = append(directories, cwdConfigDirectory)

	// append ${home}
	homeDirectory, err := userHome()
	if err != nil {
		return nil, err
	}
	// check if home found
	if homeDirectory != nil {
		directories = append(directories, *homeDirectory)
	}

	return directories, nil
}
