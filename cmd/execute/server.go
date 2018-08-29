package execute

import (
	"../../dispatcher"
	"../../nms"
	"../../server"
	"../../server/websocket"
	"../../signal"
)

func Server() error {
	instance := &nms.Instance{}
	websocketManager := websocket.NewWebSocketManager()

	srv, err := server.RunServer(websocketManager)
	if err != nil {
		return nil
	}

	websocketDisptacher := dispatcher.NewWebSocketDispatcher(instance, websocketManager)

	sig := signal.CreateSignal()
	sig.Await()

	websocketDisptacher.Close()

	err = srv.Shutdown()
	return err
}
