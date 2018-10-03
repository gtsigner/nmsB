package dispatch

import (
	"log"

	"../../../message"
	"../context"
)

func createDebugMessage(context *context.DispatchContext, err error) *message.DebugMessage {
	baseMessage := context.RequestManager.CreateMessage(message.Debug, message.DllToClients)

	text := err.Error()
	debugType := message.DebugError

	msg := &message.DebugMessage{
		Message:   baseMessage,
		DebugType: &debugType,
		Text:      &text,
	}

	return msg
}

func dispatchWebSocketError(context *context.DispatchContext, err error) {
	debugMessage := createDebugMessage(context, err)
	e := context.Client.Write(debugMessage)
	if e != nil {
		log.Println(e)
	}
}
