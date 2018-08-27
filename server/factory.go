package server

import (
	"./handler"
	"./websocket"
)

func CreateServer(websocketManager *websocket.WebSocketManager) (*Server, error) {
	serverMux, err := handler.RootRouter("./public", websocketManager)
	if err != nil {
		return nil, err
	}

	server := NewServer(4000, serverMux)
	return server, nil
}

func RunServer(websocketManager *websocket.WebSocketManager) (*Server, error) {
	server, err := CreateServer(websocketManager)
	if err != nil {
		return nil, err
	}

	server.Init()
	server.Serve()

	return server, nil
}
