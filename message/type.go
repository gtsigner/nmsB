package message

type MessageType string

const (
	DllHandshake    MessageType = "DllHandshake"
	ClientHandshake MessageType = "ClientHandshake"
	ServerStatus    MessageType = "ServerStatus"
	Error           MessageType = "Error"
)
