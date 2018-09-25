package dispatch

import (
	"log"

	"../http/websocket"
	"./context"
)

func dispatchWebSocketError(ctx *context.DispatchContext, webSocket *websocket.WebSocket, err error) {
	log.Printf("error on websocket [ %s, %s ]: %s\n", webSocket.Id, webSocket.RemoteAddr(), err)
	// TODO handle error with broadcast message
}
