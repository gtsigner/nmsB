package connection

import (
	"../../message/json"
	"../http/websocket"
)

type Connection struct {
	webSocket *websocket.WebSocket
}

func (connection *Connection) IsWebSocket(webSocket *websocket.WebSocket) bool {
	return connection.webSocket.Id == webSocket.Id
}

func (connection *Connection) Write(v interface{}) error {
	data, err := json.Encode(v)
	if err != nil {
		return err
	}

	connection.WriteString(data)
	return nil
}

func (connection *Connection) WriteString(data string) {
	// write the json string to the websocket
	connection.webSocket.Write(data)
}
