package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"../../config"
)

type HttpServer struct {
	Serv       *http.Server
	Handler    *http.ServeMux
	HttpConfig *config.HttpConfig
}

func NewHttpServer(httpConfig *config.HttpConfig, handler *http.ServeMux) *HttpServer {
	return &HttpServer{
		Handler:    handler,
		HttpConfig: httpConfig,
	}
}

func (server *HttpServer) Init() {
	address := fmt.Sprintf("%s:%d", *server.HttpConfig.Address, *server.HttpConfig.Port)

	timeouts := server.HttpConfig.Timeouts
	// create http server
	server.Serv = &http.Server{
		Addr:    address,
		Handler: server.Handler,
		// Good practice to set timeouts to avoid Slowloris attacks.
		ReadTimeout:  *timeouts.ReadTimeout,
		IdleTimeout:  *timeouts.IdleTimeout,
		WriteTimeout: *timeouts.WriteTimeout,
	}
}

func (server *HttpServer) Serve() {
	go func() {
		log.Printf("Listening on %v", server.Serv.Addr)
		err := server.Serv.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()
}

func (server *HttpServer) Shutdown() error {
	log.Println("shutdown http server")
	// create shutdown context for server
	wait := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// shutdown the server
	err := server.Serv.Shutdown(ctx)
	return err
}
