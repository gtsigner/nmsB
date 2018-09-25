package message

type Message struct {
	Type      *MessageType
	Direction *MessageDirection
	ClientId  *string
	RequestId *string
}
