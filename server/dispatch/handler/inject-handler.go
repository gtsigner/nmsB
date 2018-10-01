package handler

import (
	"../../../message"
	"../../http/websocket"
	"../context"
	"./response"
)

func InjectHandler(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.Message, data string) error {
	err := response.PushDebug(ctx, "injecting dll and calling remote initliztaion")
	if err != nil {
		return err
	}

	// inject the dll into target process
	err = ctx.Injector.Inject()
	if err != nil {
		return err
	}

	err = response.PushDebug(ctx, "successful injected dll")
	if err != nil {
		return err
	}

	err = ctx.Injector.CallInitProc()
	if err != nil {
		return err
	}

	err = response.PushDebug(ctx, "successful called remote initliztaion")
	if err != nil {
		return err
	}

	err = response.PushInfo(ctx, "successful injected and executed remote initliztaion")
	return err
}
