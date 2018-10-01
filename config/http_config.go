package config

import (
	"time"
)

type HttpConfig struct {
	Port            *uint16      `yaml:"port" json:"port"`
	Address         *string      `yaml:"address" json:"address"`
	PublicDirectory *string      `yaml:"public-directory" json:"public-directory"`
	Timeouts        *HttpTimeout `yaml:"timeouts" json:"timeouts"`
}

type HttpTimeout struct {
	WriteTimeout *time.Duration `yaml:"write" json:"write"`
	ReadTimeout  *time.Duration `yaml:"read" json:"read"`
	IdleTimeout  *time.Duration `yaml:"idle" json:"idle"`
}
