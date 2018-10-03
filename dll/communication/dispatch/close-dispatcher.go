package dispatch

import (
	"../context"
)

func dispatchWebSocketClose(context *context.DispatchContext) error {
	context.Shutdown <- true
	return nil
}
