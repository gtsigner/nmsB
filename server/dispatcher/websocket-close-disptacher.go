package dispatcher

import (
	"log"

	"../nms"
	"../server/websocket"
)

func dispatchWebSocketClose(instance *nms.Instance, webSocket *websocket.WebSocket) error {
	log.Printf("websocket [ %s ] to remote [ %s ] terminated\n", webSocket.Id, webSocket.RemoteAddr())
	return nil
}
