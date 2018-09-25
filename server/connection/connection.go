package connection

import (
	"encoding/json"

	"../../message"
	"../http/websocket"
)

type Connection struct {
	webSocket *websocket.WebSocket
}

func (connection *Connection) IsWebSocket(webSocket *websocket.WebSocket) bool {
	return connection.webSocket.Id == webSocket.Id
}

func (connection *Connection) Write(msg *message.Message) error {
	// convert the message to json
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	// write the json string to the websocket
	connection.webSocket.Write(string(bytes))

	return nil
}
