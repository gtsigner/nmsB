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

	// veriofy if message type given
	if msg.Type == nil {
		return fmt.Errorf("unable to dispatch message [ %s ], besause type is nil", data)
	}

	// check the message type
	switch *msg.Type {
	case message.DllHandshake:
		err := handler.DllHandshakeHandler(ctx, webSocket, msg, data)
		return err
	case message.ClientHandshake:
		err := handler.ClientHandshakeHandler(ctx, webSocket, msg, data)
		return err
	case message.Inject:
		err := handler.InjectHandler(ctx, webSocket, msg, data)
		return err
	}

	// verify if message direction given
	if msg.Direction == nil {
		return fmt.Errorf("unable to dispatch message [ %s ], besause direction is nil", data)
	}

	// forword message based on the message direction
	switch *msg.Direction {
	case message.ClientToDll:
		err := handler.Client2DllHandler(ctx, webSocket, msg, data)
		return err
	case message.DllToClient:
		err := handler.Dll2ClientHandler(ctx, webSocket, msg, data)
		return err
	case message.DllToClients:
		err := handler.Dll2ClientsHandler(ctx, webSocket, msg, data)
		return err
	}

	// return unable to dispatch
	return fmt.Errorf("unable to dispatch message [ %s ], because unknown message type [ %s ] or wrong direction [ %s ]",
		data, *msg.Type, *msg.Direction)
}
