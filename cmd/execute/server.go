package execute

import (
	"../../server"
	"../../server/websocket"
	"../../signal"
)

func Server() error {
	websocketManager := websocket.NewWebSocketManager()
	srv, err := server.RunServer(websocketManager)
	if err != nil {
		return nil
	}

	sig := signal.CreateSignal()
	sig.Await()

	err = srv.Shutdown()
	return err
}
