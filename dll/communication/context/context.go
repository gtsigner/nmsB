package context

import (
	"../../../config"
	"../../http"
	"../debug"
	"../request"
)

type DispatchContext struct {
	Version        string
	Release        string
	Config         *config.Config
	Client         *http.Client
	Debug          *debug.DebugManager
	RequestManager *request.RequestManager
	Shutdown       chan bool
	Connected      chan bool
}

func CreateDispatchContext(version string, release string, config *config.Config, client *http.Client) *DispatchContext {
	requestManager := request.CreateRequestManager(client)
	debug := debug.CreateDebugManager(requestManager, client)
	return &DispatchContext{
		Version:        version,
		Release:        release,
		Config:         config,
		Client:         client,
		Debug:          debug,
		RequestManager: requestManager,
		Shutdown:       make(chan bool),
		Connected:      make(chan bool),
	}
}
