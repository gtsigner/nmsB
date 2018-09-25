package dispatch

import (
	"./context"
)

func CreateDispacther(ctx *context.DispatchContext) *Dispatcher {
	// create the disptacher
	disptacher := &Dispatcher{
		connectionManager: connectionManager,
		webSocketManager:  webSocketManager,
		closeing:          make(chan bool),
	}

	// start the disptach loop
	go disptacher.dispatch()

	return disptacher
}
