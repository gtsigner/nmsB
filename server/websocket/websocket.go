package websocket

import (
	ws "github.com/gorilla/websocket"
	"net"
	"log"
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
	conn *ws.Conn
}

func NewWebSocket(id string, conn *ws.Conn) *WebSocket {
	webSocket := &WebSocket{
		Id:              id,
		OnClose:         make(chan ws.CloseError),
		OnError:         make(chan error),
		InBoundMessage:  make(chan string),
		OutBoundMessage: make(chan string),
		conn:            conn,
	}

	// set the close handler
	conn.SetCloseHandler(func(code int, text string) error {
		log.Println("closedxx")
		// emit the close event
		webSocket.OnClose <- ws.CloseError{
			Code: code,
			Text: text,
		}
		return nil
	})

	go webSocket.inBoundLoop()
	go webSocket.outBoundLoop()

	return webSocket
}

func (webSocket *WebSocket) inBoundLoop() {
	defer webSocket.Close()
	for {
		_, bytes, err := webSocket.conn.ReadMessage()
		if err != nil {			
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure) {
				webSocket.OnError <- err
			}
			return
		}

		message := string(bytes)
		webSocket.InBoundMessage <- message
	}
}

func (webSocket *WebSocket) outBoundLoop() {
	defer webSocket.Close()
	for {
		select {
		case message := <-webSocket.OutBoundMessage:
			bytes := []byte(message)
			err := webSocket.conn.WriteMessage(ws.TextMessage, bytes)
			if err != nil {
				webSocket.OnError <- err
			}
		case <-webSocket.OnClose:
			return
		}
	}
}

func (webSocket *WebSocket) Write(message string) {
	webSocket.OutBoundMessage <- message
}

func (webSocket *WebSocket) Close() {
	webSocket.conn.Close()
	webSocket.OnClose <- ws.CloseError{
		Code: 1000, // normal close
		Text: "Normal Closure",
	}
}

func (webSocket *WebSocket) RemoteAddr() net.Addr {
	return webSocket.conn.RemoteAddr()
}
