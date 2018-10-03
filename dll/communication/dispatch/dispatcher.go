package dispatch

import (
	"../context"
)

type Dispatcher struct {
	closing chan bool
	context *context.DispatchContext
}

func (dispatcher *Dispatcher) dispatch() {
	client := dispatcher.context.Client
	for {
		select {
		case <-client.OnOpen:
			go dispatcher.onOpen()
		case e := <-client.OnError:
			go dispatcher.onError(e)
		case <-client.OnClose:
			go dispatcher.onClose()
		case e := <-client.OnMessage:
			go dispatcher.onMessage(e)
		case <-dispatcher.closing:
			return
		}
	}
}

func (dispatcher *Dispatcher) onOpen() {
	err := dispatchWebSocketOpen(dispatcher.context)
	if err != nil {
		dispatchWebSocketError(dispatcher.context, err)
	}
}

func (dispatcher *Dispatcher) onError(err error) {
	dispatchWebSocketError(dispatcher.context, err)
}

func (dispatcher *Dispatcher) onClose() {
	err := dispatchWebSocketClose(dispatcher.context)
	if err != nil {
		dispatchWebSocketError(dispatcher.context, err)
	}
}

func (dispatcher *Dispatcher) onMessage(message string) {
	err := dispatchWebSocketMessage(dispatcher.context, message)
	if err != nil {
		dispatchWebSocketError(dispatcher.context, err)
	}
}

func (dispatcher *Dispatcher) Close() {
	dispatcher.closing <- true
}
