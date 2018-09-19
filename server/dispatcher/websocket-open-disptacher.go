package dispatcher

import (
	"log"

	"../nms"
	"../server/websocket"
)

func dispatchWebSocketOpen(instance *nms.Instance, webSocket *websocket.WebSocket) error {
	log.Printf("websocket [ %s ] established to remote [ %s ]\n", webSocket.Id, webSocket.RemoteAddr())

	return nil
}
