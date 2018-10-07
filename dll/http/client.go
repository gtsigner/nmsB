package http

import (
	"fmt"
	"log"
	"net/url"

	"../../config"
	"../../message/json"
	"./websocket"
)

type Client struct {
	config    *config.Config
	closing   chan bool
	webSocket *websocket.WebSocket

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

func (client *Client) Write(v interface{}) error {
	if client.webSocket == nil {
		return nil
	}

	// encode the message to json
	data, err := json.Encode(v)
	if err != nil {
		return err
	}
	// write the message
	client.webSocket.Write(data)

	return nil
}

func (client *Client) Init() error {
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

	go client.forwarder()

	// notify about success
	client.OnOpen <- true

	return nil
}

func (client *Client) forwarder() {
	defer client.Close()
	for {
		select {
		case e := <-client.webSocket.InBoundMessage:
			client.OnMessage <- e
		case e := <-client.webSocket.OutBoundMessage:
			client.onOutBoundMessage(e)
		case e := <-client.webSocket.OnError:
			client.OnError <- e
		case e := <-client.webSocket.OnClose:
			client.OnClose <- e
		case <-client.closing:
			return
		}
	}
}

func (client *Client) onOutBoundMessage(message string) {

}

func (client *Client) Close() {
	if client.webSocket != nil {
		client.webSocket.Close()
	}
}
