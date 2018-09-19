package http

import (
	"../../config"
	"./handler"
	"./websocket"
)

func CreateHttpServer(httpConfig *config.HttpConfig, websocketManager *websocket.WebSocketManager) (*HttpServer, error) {
	publicDirectory := *httpConfig.PublicDirectory
	serverMux, err := handler.RootRouter(publicDirectory, websocketManager)
	if err != nil {
		return nil, err
	}

	server := NewHttpServer(httpConfig, serverMux)
	return server, nil
}

func RunHttpServer(httpConfig *config.HttpConfig, websocketManager *websocket.WebSocketManager) (*HttpServer, error) {
	server, err := CreateHttpServer(httpConfig, websocketManager)
	if err != nil {
		return nil, err
	}

	server.Init()
	server.Serve()

	return server, nil
}
