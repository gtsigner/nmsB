package request

import (
	"strconv"
	"sync"

	"../../../message"
	"../../../message/json"
	"../../http"
)

type RequestManager struct {
	lock     sync.Mutex
	counter  uint64
	pending  map[string]chan string
	client   *http.Client
	clientId string
}

func (manager *RequestManager) CreateMessage(messageType message.MessageType, direction message.MessageDirection) *message.Message {
	requestId := manager.NextRequestID()

	// lock the manager for dispatch
	manager.lock.Lock()
	defer manager.lock.Unlock()

	msg := &message.Message{
		RequestId: &requestId,
		Direction: &direction,
		Type:      &messageType,
		ClientId:  &manager.clientId,
	}
	return msg
}

func (manager *RequestManager) SetClientId(clientId string) {
	// lock the manager for dispatch
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.clientId = clientId
}

func (manager *RequestManager) Dispatch(data string) (bool, error) {
	// decode the message
	msg, err := json.DecodeMessage(data)
	if err != nil {
		return false, err
	}

	// check if the message has a request id
	if msg.RequestId == nil {
		return false, nil
	}

	// lock the manager for dispatch
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// find the callback
	callback, ok := manager.pending[*msg.RequestId]
	if !ok {
		return false, nil
	}
	// remove the pending request
	delete(manager.pending, *msg.RequestId)

	// push the data back to the callback
	callback <- data

	return true, nil
}

func (manager *RequestManager) NextRequestID() string {
	// lock the manager for new request id
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// convert the id to string
	id := strconv.FormatUint(manager.counter, 10)
	// increment to next id
	manager.counter = manager.counter + 1
	return id
}

func (manager *RequestManager) Request(requestId string, v interface{}) (string, error) {
	// lock the manager for Request
	manager.lock.Lock()
	// unlock the manager
	defer manager.lock.Unlock()

	// request the callback
	callback := make(chan string, 1)
	manager.pending[requestId] = callback

	// send the message
	err := manager.client.Write(v)
	if err != nil {
		return "", err
	}

	// unlock the manager to wait for the response
	manager.lock.Unlock()
	// wait for the response
	responseDate := <-callback
	// lock the manager to handle the response
	manager.lock.Lock()

	// remove the pending request
	delete(manager.pending, requestId)

	return responseDate, err
}

func (manager *RequestManager) RequestEncode(requestId string, v interface{}, r interface{}) error {
	// execute the request
	response, err := manager.Request(requestId, v)
	if err != nil {
		return err
	}
	// decode the message
	err = json.Decode(response, r)
	return err
}
