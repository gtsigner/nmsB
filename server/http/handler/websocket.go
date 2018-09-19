package handler

import (
	"../websocket"
	ws "github.com/gorilla/websocket"
	"log"
	"net/http"
)

func WebSocket(webSocketManager *websocket.WebSocketManager) http.Handler {
	// create the websocket upgrader
	upgrader := ws.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// create new handler func
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// upgrade the request to a websocket
		conn, err := upgrader.Upgrade(response, request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		// reqiester and notify about new websocket
		err = webSocketManager.Open(conn)
		if err != nil {
			log.Printf("fail to open websocket to remote [ %s ], because %s\n", conn.RemoteAddr(), err.Error())
		}
	})
}
