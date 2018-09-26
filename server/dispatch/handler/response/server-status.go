package response

import (
	"../../../../message"
	"../../context"
)

func PushServerStatus(ctx *context.DispatchContext) error {
	baseMessage, err := CreateMessage(message.ServerStatus, message.ServerToClients)
	if err != nil {
		return err
	}

	version := ctx.Version
	release := ctx.Release
	clients := ctx.ConnectionManager.ConnectedClients()
	connected := ctx.ConnectionManager.IsDllConnected()

	msg := &message.ServerStatusMessage{
		Message:   *baseMessage,
		Version:   &version,
		Release:   &release,
		Connected: &connected,
		Clients:   &clients,
	}

	err = ctx.ConnectionManager.WriteToClients(msg)
	return err
}
