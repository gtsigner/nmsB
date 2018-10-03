package instance

import (
	"../../config"
	"../communication/context"
	"../communication/dispatch"
	"../http"
)

type DllInstance struct {
	Version         string
	Release         string
	Config          *config.Config
	Client          *http.Client
	Dispatcher      *dispatch.Dispatcher
	DispatchContext *context.DispatchContext
}

func NewDllInstance(version string, release string) *DllInstance {
	return &DllInstance{
		Version: version,
		Release: release,
	}
}
