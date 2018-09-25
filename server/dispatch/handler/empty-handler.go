package handler

import (
	"../../../message"
	"../../http/websocket"
	"../context"
)

func EmptyHandler(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.Message, data string) error {
	return nil
}
