package response

import (
	"../../../../message"
	"../../../../utils"
)

var (
	DefaultClientId = "server"
)

func CreateMessage(messageType message.MessageType, direction message.MessageDirection) (*message.Message, error) {
	requestId, err := utils.RandString(int64(32))
	if err != nil {
		return nil, err
	}

	msg := CreateMessageFor(messageType, direction, requestId)
	return msg, nil
}

func CreateMessageFor(messageType message.MessageType, direction message.MessageDirection, requestId string) *message.Message {
	return &message.Message{
		RequestId: &requestId,
		Direction: &direction,
		Type:      &messageType,
		ClientId:  &DefaultClientId,
	}
}
