package response

import (
	"../../../../message"
	"../../../http/websocket"
	"../../context"
)

func PushHandshakeACKToClient(ctx *context.DispatchContext, webSocket *websocket.WebSocket, request *message.Message) error {
	ackMessage := CreateHandshakeACK(webSocket, request, message.ServerToClient)
	err := ctx.ConnectionManager.WriteToClient(webSocket.Id, ackMessage)
	return err
}

func PushHandshakeACKToDll(ctx *context.DispatchContext, webSocket *websocket.WebSocket, request *message.Message) error {
	ackMessage := CreateHandshakeACK(webSocket, request, message.ServerToDll)
	err := ctx.ConnectionManager.WriteToDll(ackMessage)
	return err
}

func CreateHandshakeACK(webSocket *websocket.WebSocket, request *message.Message, direction message.MessageDirection) *message.Message {
	msg := CreateMessageFor(message.HandshakeACK, direction, *request.RequestId)
	msg.ClientId = &webSocket.Id
	return msg
}
