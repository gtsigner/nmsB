package handler

import (
	"fmt"

	"../../../message"
	"../../../message/json"
	"../../http/websocket"
	"../context"
)

func decodeDllHandshake(data string) (*message.DllHandshakeMessage, error) {
	var msg message.DllHandshakeMessage
	err := json.Decode(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func DllHandshakeHandler(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.Message, data string) error {
	// decode the handshake message from dll
	handshake, err := decodeDllHandshake(data)
	if err != nil {
		return err
	}

	// check if version and release inside the handshake
	if handshake.Version == nil || handshake.Release == nil {
		return fmt.Errorf("unable to verify integrity for dll connection on websocket [ %s ], because version or release nil", webSocket.Id)
	}

	// check if the version and release are equal
	if *handshake.Version != ctx.Version || *handshake.Release != ctx.Release {
		return fmt.Errorf("unable to verify integrity for dll connection on websocket [ %s ], because version [ %s <-> %s ] or release [ %s <-> %s ] not equal",
			webSocket.Id, *handshake.Version, ctx.Version, *handshake.Release, ctx.Release)
	}

	// register successful dll connection
	ctx.ConnectionManager.RegisterDll(webSocket)

	return nil
}
