package dispatch

import (
	"log"

	"../http/websocket"
	"./context"
)

type Dispatcher struct {
	closeing chan bool
	ctx      *context.DispatchContext
}

func (disptacher *Dispatcher) dispatch() {
	log.Println("starting websocket disptacher...")
	for {
		select {
		case e := <-disptacher.ctx.WebSocketManager.OnOpen:
			disptacher.onOpen(e)
		case e := <-disptacher.ctx.WebSocketManager.OnError:
			disptacher.onError(e)
		case e := <-disptacher.ctx.WebSocketManager.OnClose:
			disptacher.onClose(e)
		case e := <-disptacher.ctx.WebSocketManager.InBoundMessage:
			disptacher.onMessage(e)
		case e := <-disptacher.ctx.WebSocketManager.OutBoundMessage:
			disptacher.onOutMessage(e)
		case <-disptacher.closeing:
			log.Println("closing websocket disptacher...")
			return
		}
	}
}

func (disptacher *Dispatcher) onOpen(event websocket.OpenEvent) {
	err := dispatchWebSocketOpen(disptacher.connectionManager, event.WebSocket)
	if err != nil {
		dispatchWebSocketError(disptacher.connectionManager, event.WebSocket, err)
	}
}

func (disptacher *Dispatcher) onError(event websocket.ErrorEvent) {
	dispatchWebSocketError(disptacher.connectionManager, event.WebSocket, event.Error)
}

func (disptacher *Dispatcher) onClose(event websocket.CloseEvent) {
	err := dispatchWebSocketClose(disptacher.connectionManager, event.WebSocket)
	if err != nil {
		dispatchWebSocketError(disptacher.connectionManager, event.WebSocket, err)
	}
}

func (disptacher *Dispatcher) onMessage(event websocket.InBoundMessageEvent) {
	err := DispatchMessage(disptacher.connectionManager, event.WebSocket, event.Message)
	if err != nil {
		dispatchWebSocketError(disptacher.connectionManager, event.WebSocket, err)
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
