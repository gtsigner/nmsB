package dispatcher

import (
	"log"

	"../server/websocket"
)

func dispatchWebSocketError(webSocket *websocket.WebSocket, err error) {
	log.Printf("error on websocket [ %s, %s ]: %s\n", webSocket.Id, webSocket.RemoteAddr(), err)
}
