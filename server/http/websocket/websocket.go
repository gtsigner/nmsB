package websocket

import (
	"net"

	ws "github.com/gorilla/websocket"
)

type WebSocket struct {
	// id
	Id string

	// Events
	OnClose         chan ws.CloseError
	OnError         chan error
	InBoundMessage  chan string
	OutBoundMessage chan string

	// private
	conn    *ws.Conn
	closing chan bool
}

func NewWebSocket(id string, conn *ws.Conn) *WebSocket {
	webSocket := &WebSocket{
		Id:              id,
		OnClose:         make(chan ws.CloseError),
		OnError:         make(chan error),
		InBoundMessage:  make(chan string),
		OutBoundMessage: make(chan string),
		conn:            conn,
		closing:         make(chan bool),
	}

	/*// set the close handler
	conn.SetCloseHandler(func(code int, text string) error {
		// emit the close event
		webSocket.OnClose <- ws.CloseError{
			Code: code,
			Text: text,
		}
		return nil
	})
	//*/

	go webSocket.inBoundLoop()
	go webSocket.outBoundLoop()

	return webSocket
}

func (webSocket *WebSocket) inBoundLoop() {
	defer func() {
		webSocket.Close()
	}()

	for {
		_, bytes, err := webSocket.conn.ReadMessage()
		if err != nil {
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure, ws.CloseNoStatusReceived) {
				webSocket.OnError <- err
			}
			return
		}

		message := string(bytes)
		webSocket.InBoundMessage <- message
	}

}

func (webSocket *WebSocket) outBoundLoop() {
	defer func() {
		webSocket.Close()
	}()

	for {
		select {
		case message := <-webSocket.OutBoundMessage:
			bytes := []byte(message)
			err := webSocket.conn.WriteMessage(ws.TextMessage, bytes)
			if err != nil {
				if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure, ws.CloseNoStatusReceived) {
					webSocket.OnError <- err
				}
			}
		case <-webSocket.closing:
			return
		}
	}
}

func (webSocket *WebSocket) Write(message string) {
	webSocket.OutBoundMessage <- message
}

func (webSocket *WebSocket) Close() {
	// close the web-socket
	webSocket.conn.Close()
	// notify external about close
	webSocket.OnClose <- ws.CloseError{
		Code: 1000, // normal close
		Text: "Normal Closure",
	}
	// notify internal close
	webSocket.closing <- true
}

func (webSocket *WebSocket) RemoteAddr() net.Addr {
	return webSocket.conn.RemoteAddr()
}
