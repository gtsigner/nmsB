package websocket

import (
	"sync"

	"../../utils"
	ws "github.com/gorilla/websocket"
)

type WebSocketManager struct {
	lock            sync.RWMutex
	OnOpen          chan OpenEvent
	OnClose         chan CloseEvent
	OnError         chan ErrorEvent
	InBoundMessage  chan InBoundMessageEvent
	OutBoundMessage chan OutBoundMessageEvent
	webSockets      map[string]*WebSocket
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{ // TODO Remove Buffer from Events
		lock:            sync.RWMutex{},
		OnOpen:          make(chan OpenEvent, 1),
		OnClose:         make(chan CloseEvent, 1),
		OnError:         make(chan ErrorEvent, 1),
		InBoundMessage:  make(chan InBoundMessageEvent, 1),
		OutBoundMessage: make(chan OutBoundMessageEvent, 1),
		webSockets:      make(map[string]*WebSocket),
	}
}

func (webSocketManager *WebSocketManager) Open(conn *ws.Conn) error {
	// generate new connection id
	id, err := utils.RandString(64)
	if err != nil {
		return err
	}

	// create a new web-socket
	webSocket := NewWebSocket(id, conn)

	// aquire lock
	webSocketManager.lock.Lock()
	// store the websocket
	webSocketManager.webSockets[id] = webSocket
	// release the lock
	webSocketManager.lock.Unlock()

	// start the forwarder
	go webSocketManager.forwarder(webSocket)

	return nil
}

func (webSocketManager *WebSocketManager) forwarder(webSocket *WebSocket) {
	// notify about open
	webSocketManager.OnOpen <- OpenEvent{
		WebSocket: webSocket,
	}

	for {
		select {
		case <-webSocket.OnClose:
			webSocketManager.close(webSocket)
			return
		case err := <-webSocket.OnError:
			webSocketManager.OnError <- ErrorEvent{
				Error:     err,
				WebSocket: webSocket,
			}
		case message := <-webSocket.InBoundMessage:
			webSocketManager.InBoundMessage <- InBoundMessageEvent{
				Message:   message,
				WebSocket: webSocket,
			}
		case message := <-webSocket.OutBoundMessage:
			webSocketManager.OutBoundMessage <- OutBoundMessageEvent{
				Message:   message,
				WebSocket: webSocket,
			}
		}
	}

}

func (webSocketManager *WebSocketManager) close(webSocket *WebSocket) {
	// notify about close
	webSocketManager.OnClose <- CloseEvent{
		WebSocket: webSocket,
	}

	// aquire lock
	webSocketManager.lock.Lock()
	// delete the websocket
	delete(webSocketManager.webSockets, webSocket.Id)
	// release the lock
	webSocketManager.lock.Unlock()
}

func (webSocketManager *WebSocketManager) Broadcast(message string) {
	// aquire read lock
	webSocketManager.lock.RLock()
	// unlock read lock on end
	defer webSocketManager.lock.RUnlock()

	// send message to all websockets
	for _, webSocket := range webSocketManager.webSockets {
		webSocket.Write(message)
	}
}
