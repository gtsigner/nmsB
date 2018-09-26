package handler

import (
	"fmt"

	"../../../message"
	"../../http/websocket"
	"../context"
)

func Dll2ClientHandler(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.Message, data string) error {
	id := webSocket.Id
	// verify if webSocket is the dll connection
	isDll := ctx.ConnectionManager.IsDllConnection(id)
	if !isDll {
		return fmt.Errorf("unable to forward message [ %s ] to client, because webSocket [ %s ] is not dll connection", data, id)
	}

	// verify if client id given
	if msg.ClientId == nil {
		return fmt.Errorf("unable to forward message [ %s ] to client, because cleintId is nil", data)
	}

	// forward the message to the client
	err := ctx.ConnectionManager.WriteStringToClient(*msg.ClientId, data)
	return err
}
