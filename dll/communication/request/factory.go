package request

import (
	"sync"

	"../../http"
)

func CreateRequestManager(client *http.Client) *RequestManager {
	return &RequestManager{
		counter: 0,
		client:  client,
		lock:    sync.Mutex{},
		pending: make(map[string]chan string),
	}
}
