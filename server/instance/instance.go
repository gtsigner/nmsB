package instance

import (
	"../../config"
	"../dispatch"
	"../dispatch/context"
	"../http"
)

type ServerInstance struct {
	Version        string
	Release        string
	Config         *config.Config
	HttpServer     *http.HttpServer
	Dispatcher     *dispatch.Dispatcher
	DispactContext *context.DispatchContext
}

func NewServerInstance(version string, release string) *ServerInstance {
	return &ServerInstance{
		Version: version,
		Release: release,
	}
}
