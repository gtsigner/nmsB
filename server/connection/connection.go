package connection

import (
	"../../message"
	"../../message/json"
	"../http/websocket"
)

type Connection struct {
	webSocket *websocket.WebSocket
}

func (connection *Connection) IsWebSocket(webSocket *websocket.WebSocket) bool {
	return connection.webSocket.Id == webSocket.Id
}

func (connection *Connection) Write(msg *message.Message) error {
	data, err := json.Encode(msg)
	if err != nil {
		return err
	}
	// write the json string to the websocket
	connection.webSocket.Write(data)

	return nil
}
