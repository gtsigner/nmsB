package websocket

type OpenEvent struct {
	WebSocket *WebSocket
}

type CloseEvent struct {
	WebSocket *WebSocket
}

type ErrorEvent struct {
	Error     error
	WebSocket *WebSocket
}

type InBoundMessageEvent struct {
	Message   string
	WebSocket *WebSocket
}

type OutBoundMessageEvent struct {
	Message   string
	WebSocket *WebSocket
}
