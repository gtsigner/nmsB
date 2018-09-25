package handler

import (
	"../../../message"
	"../../http/websocket"
	"../context"
)

func ClientHandshakeHandler(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.Message, data string) error {
	// register successful client connection
	ctx.ConnectionManager.RegisterClient(webSocket)

	return nil
}
