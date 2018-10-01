package context

import (
	"../../../config"
	"../../connection"
	"../../http/websocket"
	"../../inject"
)

type DispatchContext struct {
	Version           string
	Release           string
	Config            *config.Config
	Injector          *inject.Injector
	WebSocketManager  *websocket.WebSocketManager
	ConnectionManager *connection.ConnectionManager
}

func CreateDispatchContext(version string, release string, config *config.Config, injector *inject.Injector) *DispatchContext {
	webSocketManager := websocket.NewWebSocketManager()
	connectionManager := connection.NewConnectionManager()
	return &DispatchContext{
		Config:            config,
		Version:           version,
		Release:           release,
		Injector:          injector,
		WebSocketManager:  webSocketManager,
		ConnectionManager: connectionManager,
	}
}
