package response

import (
	"../../../../message"
	"../../../http/websocket"
	"../../context"
)

func PushHandshakeACK(ctx *context.DispatchContext, webSocket *websocket.WebSocket, request *message.Message) error {
	id := webSocket.Id

	msg := CreateMessageFor(message.HandshakeACK, message.ServerToClients, *request.RequestId)
	msg.ClientId = &id

	err := ctx.ConnectionManager.WriteToClient(id, msg)
	return err
}
