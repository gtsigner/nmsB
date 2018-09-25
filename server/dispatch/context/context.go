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
