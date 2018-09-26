package handler

import (
	"fmt"

	"../../../message"
	"../../http/websocket"
	"../context"
)

func Dll2ClientsHandler(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.Message, data string) error {
	id := webSocket.Id
	// verify if webSocket is the dll connection
	isDll := ctx.ConnectionManager.IsDllConnection(id)
	if !isDll {
		return fmt.Errorf("unable to forward message [ %s ] to clients, because webSocket [ %s ] is not dll connection", data, id)
	}
	// forward the message from dll to the cleints
	ctx.ConnectionManager.WriteStringToClients(data)
	return nil
}
