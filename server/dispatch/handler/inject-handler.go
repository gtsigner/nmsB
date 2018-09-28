package handler

import (
	"../../../message"
	"../../http/websocket"
	"../../inject"
	"../context"
	"./response"
)

func InjectHandler(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.Message, data string) error {
	err := response.PushDebug(ctx, "injecting dll")
	if err != nil {
		return err
	}

	err = inject.Inject(ctx.Config)
	if err != nil {
		return err
	}

	err = response.PushInfo(ctx, "successful injected dll")
	return err
}
