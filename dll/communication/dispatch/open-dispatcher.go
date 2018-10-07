package dispatch

import (
	"fmt"
	"log"

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

func dispatchWebSocketOpen(ctx *context.DispatchContext) error {
	log.Println("websocket connection to server established")
	// create the handshake message
	handshakeMessage := createHandshakeMessage(ctx)

	// send handshake request and wait for response
	var response message.Message
	err := ctx.RequestManager.RequestEncode(*handshakeMessage.RequestId, handshakeMessage, &response)
	if err != nil {
		return err
	}

	// check if message type responsed
	if response.Type == nil {
		return fmt.Errorf("unable to acknowledged handshake, because server responsed with nil message-type")
	}

	// check if message-type is ack
	if *response.Type != message.HandshakeACK {
		return fmt.Errorf("unable to acknowledged handshake, because server responsed with message-type [ %s ]", *response.Type)
	}

	log.Println("server responsed with acknowledged handshake")

	// check if client id responsed
	if response.ClientId != nil {
		log.Printf("handshake successed with client-id [ %s ]", *response.ClientId)
		ctx.RequestManager.SetClientId(*response.ClientId)
	}

	// notify about connected
	ctx.Connected <- true

	return nil
}
