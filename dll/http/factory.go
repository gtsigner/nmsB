package http

import "../../config"

func CreateClient(cfg *config.Config) *Client {
	return &Client{
		config:    cfg,
		running:   false,
		connected: false,
		// internal
		closing:    make(chan bool),
		writing:    make(chan string),
		connecting: make(chan bool),
		// external
		OnClose:   make(chan bool),
		OnError:   make(chan error),
		OnOpen:    make(chan bool),
		OnMessage: make(chan string),
	}
}
