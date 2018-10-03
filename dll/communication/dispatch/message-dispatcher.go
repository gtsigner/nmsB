package dispatch

import (
	"fmt"

	"../../../message/json"
	"../context"
)

func dispatchWebSocketMessage(context *context.DispatchContext, data string) error {
	// forward message to request manager for pending request completion
	success, err := context.RequestManager.Dispatch(data)
	if err != nil {
		return err
	}

	// message already dispatched
	if success {
		return nil
	}

	// decode the message
	msg, err := json.DecodeMessage(data)
	if err != nil {
		return err
	}

	return fmt.Errorf("unable to dispatch message [ %s ], because unknown message type [ %s ] or wrong direction [ %s ]",
		data, *msg.Type, *msg.Direction)
}
