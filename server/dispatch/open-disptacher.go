package dispatch

import (
	"log"

	"../http/websocket"
	"./context"
)

func dispatchWebSocketOpen(ctx *context.DispatchContext, webSocket *websocket.WebSocket) error {
	log.Printf("websocket [ %s ] established to remote [ %s ]\n", webSocket.Id, webSocket.RemoteAddr())

	return nil
}
