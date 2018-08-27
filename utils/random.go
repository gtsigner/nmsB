package utils

import (
	"encoding/base64"
	"math/rand"
	"time"
)

var random *rand.Rand = createRand()

func createRand() *rand.Rand {
	now := time.Now()
	source := rand.NewSource(now.UnixNano())
	rand := rand.New(source)
	return rand
}

func RandBytes(size int64) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := random.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func RandString(size int64) (string, error) {
	bytes := make([]byte, size)
	_, err := random.Read(bytes)
	if err != nil {
		return "", err
	}
	str := base64.StdEncoding.EncodeToString(bytes)
	value := str[:size]
	return value, nil
}
