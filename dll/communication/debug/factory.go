package debug

import (
	"../../http"
	"../request"
)

func CreateDebugManager(requestManager *request.RequestManager, client *http.Client) *DebugManager {
	return &DebugManager{
		requestManager: requestManager,
		client:         client,
	}
}
