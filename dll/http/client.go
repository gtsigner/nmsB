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
	i := 0
	for {
		i = i + 1
		log.Println("next", i)
		if client.webSocket != nil {
			client.forwarderAll()
			log.Println("done", i)
			continue
		}

		select {
		// only for init
		case <-client.connecting:
			client.doConnect()
			log.Println("doConnect2")
		case <-client.closing:
			client.onClose()
			log.Println("onClose2")
		}
		log.Println("done", i)
	}
}

func (client *Client) forwarderAll() {
	select {
	// websocket
	case e := <-client.webSocket.InBoundMessage:
		client.OnMessage <- e
		log.Println("InBoundMessage")
	case e := <-client.webSocket.OutBoundMessage:
		client.onOutBoundMessage(e)
		log.Println("OutBoundMessage")
	case e := <-client.webSocket.OnError:
		client.OnError <- e
		log.Println("OnError")
	case <-client.webSocket.OnClose:
		client.onWebSocketClosed()
	// internal
	case <-client.connecting:
		client.doConnect()
		log.Println("doConnect")
	case data := <-client.writing:
		client.doWrite(data)
		log.Println("doWrite")
	case <-client.closing:
		client.onClose()
		log.Println("onClose")
	}
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
	if client.webSocket != nil {
		// write the message
		client.webSocket.Write(data)
	}
}

func (client *Client) doConnect() {
	// check if already connected
	if client.connected {
		return
	}
	// trigger connected
	err := client.connect()
	if err != nil {
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
}

func (client *Client) onOutBoundMessage(message string) {

}

func (client *Client) onWebSocketClosed() {
	log.Println("ttt")
	if !client.connected {
		return
	}
	client.webSocket = nil
	client.connected = false
	log.Println("underlying socket closed, trigger reconnected")
	go client.triggerReconnect()
}

func (client *Client) triggerReconnect() {
	log.Println("trigger reconnecting...")
	time.Sleep(time.Second * 5)
	log.Println("before")
	client.connecting <- true

	log.Println("after")
}

func (client *Client) onClose() {
	log.Println("closing client...")
	client.OnClose <- true
	client.connected = false
	if client.webSocket != nil {
		client.webSocket.Close()
	}
}
