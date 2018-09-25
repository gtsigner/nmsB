package dispatch

import (
	"log"

	"../http/websocket"
	"./context"
)

type Dispatcher struct {
	closeing chan bool
	context  *context.DispatchContext
}

func (disptacher *Dispatcher) dispatch() {
	log.Println("starting websocket disptacher...")
	defer func() {
		log.Println("closing websocket disptacher...")
	}()

	webSocketManager := disptacher.context.WebSocketManager
	for {
		select {
		case e := <-webSocketManager.OnOpen:
			disptacher.onOpen(e)
		case e := <-webSocketManager.OnError:
			disptacher.onError(e)
		case e := <-webSocketManager.OnClose:
			disptacher.onClose(e)
		case e := <-webSocketManager.InBoundMessage:
			disptacher.onMessage(e)
		case e := <-webSocketManager.OutBoundMessage:
			disptacher.onOutMessage(e)
		case <-disptacher.closeing:
			return
		}
	}
}

func (disptacher *Dispatcher) onOpen(event websocket.OpenEvent) {
	err := dispatchWebSocketOpen(disptacher.context, event.WebSocket)
	if err != nil {
		dispatchWebSocketError(disptacher.context, event.WebSocket, err)
	}
}

func (disptacher *Dispatcher) onError(event websocket.ErrorEvent) {
	dispatchWebSocketError(disptacher.context, event.WebSocket, event.Error)
}

func (disptacher *Dispatcher) onClose(event websocket.CloseEvent) {
	err := dispatchWebSocketClose(disptacher.context, event.WebSocket)
	if err != nil {
		dispatchWebSocketError(disptacher.context, event.WebSocket, err)
	}
}

func (disptacher *Dispatcher) onMessage(event websocket.InBoundMessageEvent) {
	err := DispatchMessage(disptacher.context, event.WebSocket, event.Message)
	if err != nil {
		dispatchWebSocketError(disptacher.context, event.WebSocket, err)
	}
}

func (disptacher *Dispatcher) onOutMessage(event websocket.OutBoundMessageEvent) {
	log.Printf("message [ server => %s ]: %s\n", event.WebSocket.RemoteAddr(), event.Message)
}

func (disptacher *Dispatcher) Close() {
	go func() {
		disptacher.closeing <- true
	}()
}
