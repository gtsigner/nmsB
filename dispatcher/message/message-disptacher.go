package message

import (
	"log"

	"../../nms"
	"../../server/websocket"
)

func DispatchMessage(instance *nms.Instance, webSocket *websocket.WebSocket, message string) error {
	log.Printf("message [ %s => server ]: %s\n", webSocket.RemoteAddr(), message)
	return nil
}
