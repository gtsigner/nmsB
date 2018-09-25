package instance

import (
	"../../config"
	"../dispatch"
	"../http"
	"../http/websocket"
)

type ServerInstance struct {
	Config           *config.Config
	HttpServer       *http.HttpServer
	Dispatcher       *dispatch.Dispatcher
	WebSocketManager *websocket.WebSocketManager
}

func NewServerInstance() *ServerInstance {
	return &ServerInstance{}
}
