package config

type Config struct {
	Http   *HttpConfig   `yaml:"http" json:"http"`
	Server *ServerConfig `yaml:"server" json:"server"`
}
