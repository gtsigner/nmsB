package websocket

import (
	"net/url"

	ws "github.com/gorilla/websocket"
)

type WebSocket struct {
	conn     *ws.Conn
	closing  chan bool
	outbound chan string

	// Events
	OnClose         chan bool
	OnError         chan error
	InBoundMessage  chan string
	OutBoundMessage chan string
}

func NewWebSocket() *WebSocket {
	return &WebSocket{
		OnClose:         make(chan bool),
		OnError:         make(chan error),
		InBoundMessage:  make(chan string),
		OutBoundMessage: make(chan string),
		closing:         make(chan bool),
		outbound:        make(chan string),
	}
}

func (webSocket *WebSocket) Connect(u *url.URL) error {
	c, _, err := ws.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	webSocket.conn = c

	// start the read and write loop
	go webSocket.inBoundLoop()
	go webSocket.outBoundLoop()

	return nil
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
		case message := <-webSocket.outbound:
			webSocket.OutBoundMessage <- message
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

func (webSocket *WebSocket) Close() {
	// close the web-socket
	webSocket.conn.Close()
	// notify external about close
	webSocket.OnClose <- true
	// notify internal close
	webSocket.closing <- true

}

func (webSocket *WebSocket) Write(data string) {
	webSocket.outbound <- data
}
