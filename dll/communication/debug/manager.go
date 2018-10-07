package debug

import (
	"log"

	"../../../message"
	"../../http"
	"../request"
)

type DebugManager struct {
	requestManager *request.RequestManager
	client         *http.Client
}

func (debugManager *DebugManager) pushDebugMessage(debugType message.DebugType, text string) error {
	// create the base message
	baseMessage := debugManager.requestManager.CreateMessage(message.Debug, message.DllToClients)

	// create the debug message
	msg := &message.DebugMessage{
		Message:   baseMessage,
		DebugType: &debugType,
		Text:      &text,
	}
	// write the debug message
	err := debugManager.client.Write(msg)
	return err
}

func (debugManager *DebugManager) handleError(err error) {
	e := debugManager.pushDebugMessage(message.DebugError, err.Error())
	if e != nil {
		log.Println(e)
	}
}

func (debugManager *DebugManager) Debug(text string) {
	err := debugManager.pushDebugMessage(message.DebugDebug, text)
	if err != nil {
		debugManager.handleError(err)
	}
}

func (debugManager *DebugManager) Info(text string) {
	err := debugManager.pushDebugMessage(message.DebugInfo, text)
	if err != nil {
		debugManager.handleError(err)
	}
}

func (debugManager *DebugManager) Warn(text string) {
	err := debugManager.pushDebugMessage(message.DebugWarning, text)
	if err != nil {
		debugManager.handleError(err)
	}
}

func (debugManager *DebugManager) Error(text string) {
	err := debugManager.pushDebugMessage(message.DebugError, text)
	if err != nil {
		debugManager.handleError(err)
	}
}
