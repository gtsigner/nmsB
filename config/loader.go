package config

import (
	"fmt"
)

func Load() (*Config, error) {
	configFile, err := findConfigFile()
	if err != nil {
		return nil, err
	}

	if configFile == nil {
		return nil, fmt.Errorf("unable to find config file [ %s ] in [CURDIR] and [USER_HOME]", CONFIG_FILE)
	}

	return nil, nil
}
