package context

import (
	"../../../config"
	"../../http"
	"../request"
)

type DispatchContext struct {
	Version        string
	Release        string
	Config         *config.Config
	Client         *http.Client
	RequestManager *request.RequestManager
	Shutdown       chan bool
}

func CreateDispatchContext(version string, release string, config *config.Config, client *http.Client) *DispatchContext {
	requestManager := request.CreateRequestManager(client)
	return &DispatchContext{
		Version:        version,
		Release:        release,
		Config:         config,
		Client:         client,
		RequestManager: requestManager,
		Shutdown:       make(chan bool),
	}
}
