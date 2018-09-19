package config

import (
	"time"
)

type HttpConfig struct {
	Port            *uint16
	Address         *string
	PublicDirectory *string
	Timeouts        *HttpTimeout
}

type HttpTimeout struct {
	WriteTimeout *time.Duration
	ReadTimeout  *time.Duration
	IdleTimeout  *time.Duration
}
