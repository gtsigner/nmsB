package config

type Config struct {
	Development *bool         `yaml:"development" json:"development"`
	Http        *HttpConfig   `yaml:"http" json:"http"`
	Server      *ServerConfig `yaml:"server" json:"server"`
}
