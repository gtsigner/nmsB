package instance

import (
	"../../config"
	"../http"
)

type ServerInstance struct {
	Config     *config.Config
	HttpServer *http.HttpServer
}

func NewServerInstance() *ServerInstance {
	return &ServerInstance{}
}
