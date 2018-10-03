package dispatch

import (
	"../../../message"
	"../context"
)

func createHandshakeMessage(context *context.DispatchContext) *message.DllHandshakeMessage {
	baseMessage := context.RequestManager.CreateMessage(message.DllHandshake, message.DllToServer)

	msg := &message.DllHandshakeMessage{
		Message: baseMessage,
		Version: &context.Version,
		Release: &context.Release,
	}

	return msg
}

func dispatchWebSocketOpen(context *context.DispatchContext) error {
	// create the handshake message
	handshakeMessage := createHandshakeMessage(context)

	// send handshake request and wait for response
	var response message.Message
	err := context.RequestManager.RequestEncode(*handshakeMessage.RequestId, handshakeMessage, &response)
	if err != nil {
		return err
	}

	// check if client id responsed
	if response.ClientId != nil {
		context.RequestManager.SetClientId(*response.ClientId)
	}

	return nil
}
