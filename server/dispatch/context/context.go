package context

import (
	"../../../config"
	"../../connection"
	"../../http/websocket"
)

type DispatchContext struct {
	Version           string
	Release           string
	Config            *config.Config
	WebSocketManager  *websocket.WebSocketManager
	ConnectionManager *connection.ConnectionManager
}

func CreateDispatchContext(version string, release string, config *config.Config) *DispatchContext {
	webSocketManager := websocket.NewWebSocketManager()
	connectionManager := connection.NewConnectionManager()
	return &DispatchContext{
		Config:            config,
		Version:           version,
		Release:           release,
		WebSocketManager:  webSocketManager,
		ConnectionManager: connectionManager,
	}
}
