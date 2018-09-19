package handler

import (
	"../websocket"
	"net/http"
)

func RootRouter(directory string, webSocketManager *websocket.WebSocketManager) (*http.ServeMux, error) {
	// create the file server
	fileServer, err := FileServer(directory)
	if err != nil {
		return nil, err
	}	

	// create the websocket handler
	webSocket := WebSocket(webSocketManager)

	// create server Mux
	serverMux := http.NewServeMux()
	serverMux.Handle("/", AccessLog(fileServer))
	serverMux.Handle("/ws", webSocket)

	return serverMux, nil
}
