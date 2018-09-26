package dispatch

import (
	"log"

	"../http/websocket"
	"./context"
	"./handler/response"
)

func dispatchWebSocketError(ctx *context.DispatchContext, webSocket *websocket.WebSocket, err error) {
	log.Printf("error on websocket [ %s, %s ]: %s\n", webSocket.Id, webSocket.RemoteAddr(), err)

	// push the error to all clients
	e := response.PushError(ctx, err)
	// check for errors
	if e != nil {
		log.Printf("unable to push error [ %s ], because %s", err.Error(), e.Error())
	}
}
