package config

import (
	"log"

	"../utils"
)


func Load() (*Config, error) {
	// find the config file
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}
	log.Printf("loading configuration from file [ %s ]", configFile)

	// load default config
	defaultConfig := DefaultConfig(configFile)

	// load the config from file system
	config, err := readConfig(defaultConfig, configFile)
	if err != nil {
		return nil, err
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
