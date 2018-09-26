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

	return &message.Message{
		RequestId: &requestId,
		Direction: &direction,
		Type:      &messageType,
		ClientId:  &DefaultClientId,
	}, nil
}
