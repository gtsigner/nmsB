package connection

import (
	"sync"

	"../http/websocket"
)

func NewConnection(webSocket *websocket.WebSocket) *Connection {
	return &Connection{
		webSocket: webSocket,
	}
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		clients: make(map[string]*Connection),
		dll:     nil,
		lock:    sync.RWMutex{},
	}
}
