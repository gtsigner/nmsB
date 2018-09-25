package json

import (
	"encoding/json"

	"../../message"
)

func Encode(msg *message.Message) (string, error) {
	// convert the message to json
	bytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
