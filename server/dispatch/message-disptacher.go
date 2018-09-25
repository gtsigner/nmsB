package dispatch

import (
	"fmt"
	"log"

	"../../message"
	"../../message/json"
	"../http/websocket"
	"./context"
	"./handler"
)

func DispatchMessage(ctx *context.DispatchContext, webSocket *websocket.WebSocket, data string) error {
	log.Printf("message [ %s => server ]: %s\n", webSocket.RemoteAddr(), data)

	msg, err := json.DecodeMessage(data)
	if err != nil {
		return err
	}

	if msg.Type == nil {
		return fmt.Errorf("unable to dispatch message [ %s ], besause type is nil", data)
	}

	switch *msg.Type {
	case message.DllHandshake:
		err := handler.DllHandshakeHandler(ctx, webSocket, msg, data)
		return err
	case message.ClientHandshake:
		err := handler.ClientHandshakeHandler(ctx, webSocket, msg, data)
		return err
	}
	// return unable to dispatch
	return fmt.Errorf("unable to dispatch message [ %s ], because unknown message type [ %s ]", data, *msg.Type)
}
