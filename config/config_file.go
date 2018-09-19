package config

import (
	"fmt"
	"os"

	"path/filepath"

	"../utils"
)

const (
	CONFIG_FILE = "config.yml"
)

func toConfigPath(configFile string) (*string, error) {
	// check if path is abs
	if !filepath.IsAbs(configFile) {
		configFile, err := filepath.Abs(configFile)
		if err != nil {
			return nil, err
		}
	}
	// check if the file exists
	exists, err := utils.FileExists(configFile)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, nil
	}

	return &configFile, nil
}

func getPWDConfig() (*string, error) {
	configPath := filepath.Join(".", CONFIG_FILE)
	configFile, err := toConfigPath(configPath)
	return configFile, err
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

func getUserHomeConfig() (*string, error) {
	homeDirectory, err := userHome()
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(*homeDirectory, CONFIG_FILE)
	configFile, err := toConfigPath(configPath)
	return configFile, err
}

func findConfigFile() (*string, error) {
	// try to find config in user-home
	userConfigFile, err := getUserHomeConfig()
	if err != nil {
		return nil, err
	}
	if userConfigFile != nil {
		return userConfigFile, nil
	}

	// try to find config in pwd
	pwdConfigFile, err := getPWDConfig()
	if err != nil {
		return nil, err
	}

	if pwdConfigFile != nil {
		return pwdConfigFile, nil
	}

	return nil, nil
}
