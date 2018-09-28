package message

type MessageType string

const (
	DllHandshake    MessageType = "DllHandshake"
	ClientHandshake MessageType = "ClientHandshake"
	ServerStatus    MessageType = "ServerStatus"
	Debug           MessageType = "Debug"
	HandshakeACK    MessageType = "HandshakeACK"
	Inject          MessageType = "Inject"
)
