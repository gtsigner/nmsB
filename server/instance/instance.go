package instance

import (
	"../http"
)

func NewServerInstance() *ServerInstance {
	return &ServerInstance{}
}

type ServerInstance struct {
	HttpServer *http.HttpServer
}
