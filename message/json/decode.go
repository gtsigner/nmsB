package json

import (
	"encoding/json"

	"../../message"
)

func Decode(data string, v interface{}) error {
	bytes := []byte(data)
	err := json.Unmarshal(bytes, v)
	return err
}

func DecodeMessage(data string) (*message.Message, error) {
	var msg message.Message
	err := Decode(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
