package context

import (
	"../../connection"
	"../../http/websocket"
)

type DispatchContext struct {
	Version           string
	Release           string
	WebSocketManager  *websocket.WebSocketManager
	ConnectionManager *connection.ConnectionManager
}

func CreateDispatchContext(version string, release string) *DispatchContext {
	webSocketManager := websocket.NewWebSocketManager()
	connectionManager := connection.NewConnectionManager()
	return &DispatchContext{
		Version:           version,
		Release:           release,
		WebSocketManager:  webSocketManager,
		ConnectionManager: connectionManager,
	}
}
