package connection

import (
	"fmt"
	"log"
	"sync"

	"../../message/json"
	"../http/websocket"
)

type ConnectionManager struct {
	lock    sync.RWMutex
	dll     *Connection
	clients map[string]*Connection
}

func (manager *ConnectionManager) RegisterDll(webSocket *websocket.WebSocket) {
	// lock for read and write
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// create a new connection
	connection := NewConnection(webSocket)

	// unregister previous dll connection
	manager.unregisterDll()
	// set the given connection as connection
	manager.dll = connection

	log.Printf("successful register websocket [ %s ] as dll connection", webSocket.Id)
}

func (manager *ConnectionManager) RegisterClient(webSocket *websocket.WebSocket) {
	// lock for read and write
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// create a new connection
	connection := NewConnection(webSocket)
	// register connection as client
	manager.clients[webSocket.Id] = connection

	log.Printf("successful register websocket [ %s ] as client connection", webSocket.Id)
}

func (manager *ConnectionManager) Unregister(webSocket *websocket.WebSocket) {
	// lock for read and write
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// check if the webcoket is a client
	connection, ok := manager.clients[webSocket.Id]
	if ok {
		// close the connection
		manager.close(connection)
		// remove the client from clients
		delete(manager.clients, webSocket.Id)

		log.Printf("successful unregister websocket [ %s ] from client connection", webSocket.Id)
		return
	}

	// check if a dll is connected
	if manager.dll == nil {
		return
	}

	// check if the websocket is the dll
	isDll := manager.dll.IsWebSocket(webSocket)
	if isDll {
		manager.unregisterDll()
	}
}

func (manager *ConnectionManager) unregisterDll() {
	// check if dll connection available
	if manager.dll != nil {
		// close the connection
		manager.close(manager.dll)

		log.Printf("successful unregister websocket [ %s ] from dll connection", manager.dll.webSocket.Id)
	}
	// set dll to nil
	manager.dll = nil
}

func (manager *ConnectionManager) close(connection *Connection) {
	go connection.webSocket.Close()
}

func (manager *ConnectionManager) WriteToClients(v interface{}) error {
	// convert message to string
	data, err := json.Encode(v)
	if err != nil {
		return err
	}

	// write the string to all clients
	manager.WriteStringToClients(data)

	return nil
}

func (manager *ConnectionManager) WriteStringToClients(data string) {
	// lock for read
	manager.lock.RLock()
	defer manager.lock.RUnlock()

	// write the message to all clients
	for _, connection := range manager.clients {
		connection.WriteString(data)
	}
}

func (manager *ConnectionManager) WriteToDll(v interface{}) error {
	// convert message to string
	data, err := json.Encode(v)
	if err != nil {
		return err
	}

	// write the message to dll
	err = manager.WriteStringToDll(data)
	return err
}

func (manager *ConnectionManager) WriteStringToDll(data string) error {
	// lock for read
	manager.lock.RLock()
	defer manager.lock.RUnlock()

	if manager.dll == nil {
		return fmt.Errorf("unable to write message to dll, because no dll connected")
	}

	// write the data to dll
	manager.dll.WriteString(data)
	return nil
}

func (manager *ConnectionManager) WriteToClient(id string, v interface{}) error {
	// convert message to string
	data, err := json.Encode(v)
	if err != nil {
		return err
	}

	err = manager.WriteStringToClient(id, data)
	return err
}

func (manager *ConnectionManager) WriteStringToClient(id string, data string) error {
	// lock for read
	manager.lock.RLock()
	defer manager.lock.RUnlock()

	// try to find the client
	connection, ok := manager.clients[id]
	if !ok {
		return fmt.Errorf("unable to write message to client [ %s ], because client not found", id)
	}

	// write the message to the found connection
	connection.WriteString(data)
	return nil
}

func (manager *ConnectionManager) IsDllConnection(id string) bool {
	// lock for read
	manager.lock.RLock()
	defer manager.lock.RUnlock()

	// check if a connection exists
	if manager.dll == nil {
		return false
	}

	// check if id is equal of the dll connection
	webSocket := manager.dll.webSocket
	return webSocket.Id == id
}

func (manager *ConnectionManager) IsDllConnected() bool {
	// lock for read
	manager.lock.RLock()
	defer manager.lock.RUnlock()

	return manager.dll != nil
}

func (manager *ConnectionManager) ConnectedClients() int {
	// lock for read
	manager.lock.RLock()
	defer manager.lock.RUnlock()

	count := len(manager.clients)
	return count
}
