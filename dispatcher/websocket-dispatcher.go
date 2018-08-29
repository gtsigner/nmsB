package dispatcher

import (
	"log"

	"../nms"
	"../server/websocket"
	"./message"
)

type WebSocketDispatcher struct {
	closeing         chan bool
	instance         *nms.Instance
	webSocketManager *websocket.WebSocketManager
}

func NewWebSocketDispatcher(instance *nms.Instance, webSocketManager *websocket.WebSocketManager) *WebSocketDispatcher {
	// create the disptacher
	disptacher := &WebSocketDispatcher{
		instance:         instance,
		webSocketManager: webSocketManager,
		closeing:         make(chan bool),
	}

	// start the disptach loop
	go disptacher.dispatch()

	return disptacher
}

func (disptacher *WebSocketDispatcher) dispatch() {
	log.Println("starting websocket disptacher...")
	for {
		select {
		case e := <-disptacher.webSocketManager.OnOpen:
			disptacher.onOpen(e)
		case e := <-disptacher.webSocketManager.OnError:
			disptacher.onError(e)
		case e := <-disptacher.webSocketManager.OnClose:
			disptacher.onClose(e)
		case e := <-disptacher.webSocketManager.InBoundMessage:
			disptacher.onMessage(e)
		case e := <-disptacher.webSocketManager.OutBoundMessage:
			disptacher.onOutMessage(e)
		case <-disptacher.closeing:
			log.Println("closing websocket disptacher...")
			return
		}
	}
}

func (disptacher *WebSocketDispatcher) onOpen(event websocket.OpenEvent) {
	err := dispatchWebSocketOpen(disptacher.instance, event.WebSocket)
	if err != nil {
		dispatchWebSocketError(event.WebSocket, err)
	}
}

func (disptacher *WebSocketDispatcher) onError(event websocket.ErrorEvent) {
	dispatchWebSocketError(event.WebSocket, event.Error)
}

func (disptacher *WebSocketDispatcher) onClose(event websocket.CloseEvent) {
	err := dispatchWebSocketClose(disptacher.instance, event.WebSocket)
	if err != nil {
		dispatchWebSocketError(event.WebSocket, err)
	}
}

func (disptacher *WebSocketDispatcher) onMessage(event websocket.InBoundMessageEvent) {
	err := message.DispatchMessage(disptacher.instance, event.WebSocket, event.Message)
	if err != nil {
		dispatchWebSocketError(event.WebSocket, err)
	}
}

func (disptacher *WebSocketDispatcher) onOutMessage(event websocket.OutBoundMessageEvent) {
	log.Printf("message [ server => %s ]: %s\n", event.WebSocket.RemoteAddr(), event.Message)
}

func (disptacher *WebSocketDispatcher) Close() {
	go func() {
		disptacher.closeing <- true
	}()
}
