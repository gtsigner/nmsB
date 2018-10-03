package response

import (
	"../../../../message"
	"../../context"
)

func PushError(ctx *context.DispatchContext, err error) error {
	e := PushDebugMessage(ctx, err.Error(), message.DebugError)
	return e
}

func PushDebug(ctx *context.DispatchContext, text string) error {
	e := PushDebugMessage(ctx, text, message.DebugDebug)
	return e
}

func PushInfo(ctx *context.DispatchContext, text string) error {
	e := PushDebugMessage(ctx, text, message.DebugInfo)
	return e
}

func PushWarn(ctx *context.DispatchContext, text string) error {
	e := PushDebugMessage(ctx, text, message.DebugWarning)
	return e
}

func PushDebugMessage(ctx *context.DispatchContext, text string, debugType message.DebugType) error {
	baseMessage, err := CreateMessage(message.Debug, message.ServerToClients)
	if err != nil {
		return err
	}

	msg := &message.DebugMessage{
		Message:   baseMessage,
		Text:      &text,
		DebugType: &debugType,
	}

	err = ctx.ConnectionManager.WriteToClients(msg)
	return err
}
