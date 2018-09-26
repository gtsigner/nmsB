package handler

import (
	"../../../message"
	"../../http/websocket"
	"../context"
)

func Client2DllHandler(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.Message, data string) error {
	// write the recived message to dll
	err := ctx.ConnectionManager.WriteStringToDll(data)
	return err
}
