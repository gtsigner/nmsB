package dispatch

import (
	"log"

	"../http/websocket"
	"./context"
)

type Dispatcher struct {
	closing chan bool
	context *context.DispatchContext
}

func (dispatcher *Dispatcher) dispatch() {
	log.Println("starting websocket dispatcher...")
	defer func() {
		log.Println("closing websocket dispatcher...")
	}()

	webSocketManager := dispatcher.context.WebSocketManager
	for {
		select {
		case e := <-webSocketManager.OnOpen:
			go dispatcher.onOpen(e)
		case e := <-webSocketManager.OnError:
			go dispatcher.onError(e)
		case e := <-webSocketManager.OnClose:
			go dispatcher.onClose(e)
		case e := <-webSocketManager.InBoundMessage:
			go dispatcher.onMessage(e)
		case e := <-webSocketManager.OutBoundMessage:
			go dispatcher.onOutMessage(e)
		case <-dispatcher.closing:
			return
		}
	}
}

func (dispatcher *Dispatcher) onOpen(event websocket.OpenEvent) {
	err := dispatchWebSocketOpen(dispatcher.context, event.WebSocket)
	if err != nil {
		dispatchWebSocketError(dispatcher.context, event.WebSocket, err)
	}
}

func (dispatcher *Dispatcher) onError(event websocket.ErrorEvent) {
	dispatchWebSocketError(dispatcher.context, event.WebSocket, event.Error)
}

func (dispatcher *Dispatcher) onClose(event websocket.CloseEvent) {
	err := dispatchWebSocketClose(dispatcher.context, event.WebSocket)
	if err != nil {
		dispatchWebSocketError(dispatcher.context, event.WebSocket, err)
	}
}

func (dispatcher *Dispatcher) onMessage(event websocket.InBoundMessageEvent) {
	err := DispatchMessage(dispatcher.context, event.WebSocket, event.Message)
	if err != nil {
		dispatchWebSocketError(dispatcher.context, event.WebSocket, err)
	}
}

func (dispatcher *Dispatcher) onOutMessage(event websocket.OutBoundMessageEvent) {
	log.Printf("message [ server => %s ]: %s\n", event.WebSocket.RemoteAddr(), event.Message)
}

func (dispatcher *Dispatcher) Close() {
	dispatcher.closing <- true
}
