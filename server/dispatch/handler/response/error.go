package response

import (
	"../../../../message"
	"../../context"
)

func PushError(ctx *context.DispatchContext, err error) error {
	baseMessage, err := CreateMessage(message.Error, message.ServerToClients)
	if err != nil {
		return err
	}

	errorMessage := err.Error()
	msg := &message.ErrorMessage{
		Message: *baseMessage,
		Error:   &errorMessage,
	}

	err = ctx.ConnectionManager.WriteToClients(msg)
	return err
}
