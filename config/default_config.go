package config

import (
	"time"
)

func DefaultConfig() *Config {
	httpConfig := DefaultsHttpConfig()
	serverConfig := DefaultServerConfig()
	return &Config{
		Http:   httpConfig,
		Server: serverConfig,
	}
}

func DefaultServerConfig() *ServerConfig {
	dllPath := "./nms.dll"
	initProcName := "Init"
	processName := "nms.exe"
	return &ServerConfig{
		DllPath:      &dllPath,
		ProcessName:  &processName,
		InitProcName: &initProcName,
	}
}

func DefaultsHttpConfig() *HttpConfig {
	port := uint16(4000)
	address := "0.0.0.0"
	publicDirectory := "./public"
	timeouts := DefaultHttpTimeout()
	return &HttpConfig{
		Port:            &port,
		Address:         &address,
		Timeouts:        timeouts,
		PublicDirectory: &publicDirectory,
	}
}

func DefaultHttpTimeout() *HttpTimeout {
	readTimeout := 15 * time.Second
	idleTimeout := 15 * time.Second
	writeTimeout := 15 * time.Second
	return &HttpTimeout{
		IdleTimeout:  &idleTimeout,
		ReadTimeout:  &readTimeout,
		WriteTimeout: &writeTimeout,
	}

}
