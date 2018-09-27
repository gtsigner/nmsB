package handler

import (
	"../../../message"
	"../../http/websocket"
	"../context"
	"./response"
)

func ClientHandshakeHandler(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.Message, data string) error {
	// register successful client connection
	ctx.ConnectionManager.RegisterClient(webSocket)

	// acknowledge the handshake
	err := response.PushHandshakeACK(ctx, webSocket, msg)
	if err != nil {
		return err
	}

	// notify about changed server
	err = response.PushServerStatus(ctx)
	return err
}
