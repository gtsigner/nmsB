package http

import "../../config"

func CreateClient(cfg *config.Config) *Client {
	return &Client{
		config:    cfg,
		closing:   make(chan bool),
		OnClose:   make(chan bool),
		OnError:   make(chan error),
		OnOpen:    make(chan bool),
		OnMessage: make(chan string),
	}
}
