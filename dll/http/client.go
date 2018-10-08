package http

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"../../config"
	"../../message/json"
	"./websocket"
)

type Client struct {
	running    bool
	connected  bool
	config     *config.Config
	closing    chan bool
	writing    chan string
	connecting chan bool
	webSocket  *websocket.WebSocket

	// public
	OnMessage chan string
	OnOpen    chan bool
	OnClose   chan bool
	OnError   chan error
}

func (client *Client) serverUrl() (*url.URL, error) {
	if client.config == nil {
		return nil, fmt.Errorf("unable to initialize client, because config is nil")
	}
	cfg := client.config

	if cfg.Http == nil {
		return nil, fmt.Errorf("unable to initialize client, because http config is nil")
	}

	if cfg.Http.Port == nil {
		return nil, fmt.Errorf("unable to initialize client, because port is missing in http config")
	}

	host := fmt.Sprintf("%s:%d", "127.0.0.1", *cfg.Http.Port)

	u := &url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   "/ws",
	}
	return u, nil
}

func (client *Client) Init() error {
	// start the forwarder
	go client.forwarder()

	// start connecting
	client.connecting <- true

	return nil
}

func (client *Client) Write(v interface{}) error {
	// encode the message to json
	data, err := json.Encode(v)
	if err != nil {
		return err
	}

	// forwarde the data to the writer
	client.writing <- data

	return nil
}

func (client *Client) Close() {
	client.closing <- true
}

func (client *Client) forwarder() {
	// defer the exit hook
	defer client.dispose()
	// set the client to running
	client.running = true

	// loop until cleitn running
	for client.running {
		// check if the client is connected
		if client.connected {
			// forward all events
			ok := client.forwarderAll()
			// check if all events ok
			if ok {
				continue
			}
			// exit if event closed
			return
		}

		// dispatch events if not connected
		select {
		case <-client.connecting:
			client.doConnect()
		case data := <-client.writing:
			client.doWrite(data)
		case <-client.closing:
			client.onClose()
			return
		}
	}
}

func (client *Client) forwarderAll() bool {
	select {
	// websocket
	case e := <-client.webSocket.InBoundMessage:
		client.OnMessage <- e
	case e := <-client.webSocket.OutBoundMessage:
		client.onOutBoundMessage(e)
	case e := <-client.webSocket.OnError:
		client.OnError <- e
	case <-client.webSocket.OnClose:
		client.onWebSocketClosed()
	// internal
	case <-client.connecting:
		client.doConnect()
	case data := <-client.writing:
		client.doWrite(data)
	case <-client.closing:
		client.onClose()
		return false
	}

	return true
}

func (client *Client) connect() error {
	// create the server url
	u, err := client.serverUrl()
	if err != nil {
		return err
	}

	// create a new websocket
	client.webSocket = websocket.NewWebSocket()

	log.Printf("connecting to server at [ %s ]", u.String())
	// connect to the server
	err = client.webSocket.Connect(u)
	if err != nil {
		return err
	}

	client.connected = true

	// notify about success
	client.OnOpen <- true

	return nil
}

func (client *Client) doWrite(data string) {
	// check if connected
	if client.webSocket == nil || !client.connected {
		return
	}

	// write the message
	client.webSocket.Write(data)
}

func (client *Client) doConnect() {
	// check if already connected
	if client.connected {
		return
	}

	// executue connect
	err := client.connect()
	if err == nil {
		return
	}

	// get the server url
	u, e := client.serverUrl()
	if e != nil {
		// forwared the error
		client.OnError <- e
		log.Printf("fail to get server endpoint url, because %s", e.Error())
		return
	}

	// forwared the error
	client.OnError <- err
	log.Printf("failed to connect to server endpoint [ %s ], because %s", u.String(), err.Error())
	// trigger reconnect
	go client.triggerReconnect()
}

func (client *Client) onOutBoundMessage(message string) {

}

func (client *Client) onWebSocketClosed() {
	// check if cleint already disconnected
	if !client.connected {
		return
	}

	// disconnect and mark as not connected
	client.webSocket = nil
	client.connected = false
	log.Println("underlying socket closed, trigger reconnected")
	// execute the reconnect
	go client.triggerReconnect()
}

func (client *Client) triggerReconnect() {
	log.Println("trigger reconnecting...")
	time.Sleep(time.Second * 5)
	client.connecting <- true
}

func (client *Client) dispose() {
	client.OnClose <- true
}

func (client *Client) onClose() {
	log.Println("closing client...")
	client.running = false
	// check if cleint connected
	if client.webSocket != nil && client.connected {
		client.webSocket.Close()
	}
}
