package json

import (
	"encoding/json"
)

func Encode(v interface{}) (string, error) {
	// convert the message to json
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
