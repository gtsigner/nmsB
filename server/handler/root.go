package handler

import (
	"../websocket"
	"net/http"
)

func RootRouter(directory string, webSocketManager *websocket.WebSocketManager) (*http.ServeMux, error) {
	// create the file server
	fileServer := AccessLog(
		FileServer(directory),
	)

	// create the websocket handler
	webSocket := WebSocket(webSocketManager)

	// create server Mux
	serverMux := http.NewServeMux()
	serverMux.Handle("/", fileServer)
	serverMux.Handle("/ws", webSocket)

	return serverMux, nil
}
