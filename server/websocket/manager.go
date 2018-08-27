package websocket

import (
	"../../utils"
	ws "github.com/gorilla/websocket"
	"log"
	"sync"
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

	log.Printf("websocket [ %s ] established to remote [ %s ]\n", id, conn.RemoteAddr())

	// create a new web-socket
	webSocket := NewWebSocket(id, conn)

	// aquire lock
	webSocketManager.lock.Lock()
	// store the websocket
	webSocketManager.webSockets[id] = webSocket
	// release the lock
	webSocketManager.lock.Unlock()

	// start the forwarder
	go func() {
		log.Println("forwarder3")
		webSocketManager.forwarder(webSocket)
	}()

	return nil
}

func (webSocketManager *WebSocketManager) forwarder(webSocket *WebSocket) {
	
	// close the websocket on end
	defer webSocketManager.close(webSocket)

	log.Println("forwarder2")
	// notify about open
	webSocketManager.OnOpen <- OpenEvent{
		WebSocket: webSocket,
	}

	log.Println("forwarder")

	for {
		select {
		case closeError := <-webSocket.OnClose:
			log.Printf("close event on websocket [ %s ] to remote [ %s ] with [ %d, %s ]\n",
				webSocket.Id,
				webSocket.RemoteAddr(),
				closeError.Code,
				closeError.Text,
			)
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
	log.Printf("closing websocket [ %s ] to remote [ %s ]\n", webSocket.Id, webSocket.RemoteAddr())
	// close the underlying connection
	webSocket.Close()

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
