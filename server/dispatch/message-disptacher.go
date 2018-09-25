package dispatch

import (
	"log"

	"../http/websocket"
	"./context"
)

func DispatchMessage(ctx *context.DispatchContext, webSocket *websocket.WebSocket, message string) error {
	log.Printf("message [ %s => server ]: %s\n", webSocket.RemoteAddr(), message)
	return nil
}
