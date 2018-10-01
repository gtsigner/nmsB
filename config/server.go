package config

type ServerConfig struct {
	InitProcName *string `yaml:"init-proc-name" json:"init-proc-name"`
	ProcessName  *string `yaml:"process-name" json:"process-name"`
	DllPath      *string `yaml:"dll-path" json:"dll-path"`
}
