package handler

import (
	"fmt"

	"../../../message"
	"../../../message/json"
	"../../http/websocket"
	"../context"
	"./response"
)

func decodeDllHandshake(data string) (*message.DllHandshakeMessage, error) {
	var msg message.DllHandshakeMessage
	err := json.Decode(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func verifyIntegrity(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.DllHandshakeMessage) error {
	// check if development activated
	if ctx.Config.Development != nil && *ctx.Config.Development {
		return nil
	}

	// check if version and release inside the handshake
	if msg.Version == nil || msg.Release == nil {
		return fmt.Errorf("unable to verify integrity for dll connection on websocket [ %s ], because version or release nil", webSocket.Id)
	}

	// check if the version and release are equal
	if *msg.Version != ctx.Version || *msg.Release != ctx.Release {
		return fmt.Errorf("unable to verify integrity for dll connection on websocket [ %s ], because version [ %s <-> %s ] or release [ %s <-> %s ] not equal",
			webSocket.Id, *msg.Version, ctx.Version, *msg.Release, ctx.Release)
	}

	return nil
}

func DllHandshakeHandler(ctx *context.DispatchContext, webSocket *websocket.WebSocket, msg *message.Message, data string) error {
	// decode the handshake message from dll
	handshake, err := decodeDllHandshake(data)
	if err != nil {
		return err
	}

	// verify the integrity of the connected dll
	err = verifyIntegrity(ctx, webSocket, handshake)
	if err != nil {
		return err
	}

	// register successful dll connection
	ctx.ConnectionManager.RegisterDll(webSocket)

	// acknowledge the handshake
	err = response.PushHandshakeACKToDll(ctx, webSocket, msg)
	if err != nil {
		return err
	}

	// notify about changed server
	err = response.PushServerStatus(ctx)
	return err
}
