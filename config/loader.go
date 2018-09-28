package config

import (
	"log"

	"../utils"
)

func Load() (*Config, error) {
	// create default config
	defaultConfig := DefaultConfig()

	// overwrite default with all config files
	config, err := readConfigFiles(defaultConfig)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func readConfigFiles(config *Config) (*Config, error) {
	allConfigFiles, err := configFiles()
	if err != nil {
		return nil, err
	}

	// check if some config files exists
	if allConfigFiles == nil || len(allConfigFiles) < 1 {
		return config, nil
	}

	for index, configFile := range allConfigFiles {
		log.Printf("loading config file [ %s ] at index [ %d ]", configFile, index)
		// loading config and owerwrite
		updatedConfig, err := readConfig(config, configFile)
		if err != nil {
			return nil, err
		}
		config = updatedConfig
	}
	return config, nil
}

func readConfig(config *Config, configFile string) (*Config, error) {
	err := utils.ReadObject(configFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
