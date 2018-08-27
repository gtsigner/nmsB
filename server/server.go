package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Port    int
	Serv    *http.Server
	Handler *http.ServeMux
}

func NewServer(port int, handler *http.ServeMux) *Server {
	return &Server{
		Port:    port,
		Handler: handler,
	}
}

func (server *Server) Init() {
	address := fmt.Sprintf(":%d", server.Port)

	// create http server
	server.Serv = &http.Server{
		Addr:    address,
		Handler: server.Handler,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

func (server *Server) Serve() {
	go func() {
		log.Printf("Listening on %v", server.Serv.Addr)
		err := server.Serv.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()
}

func (server *Server) Shutdown() error {
	log.Println("shutdown http server")
	// create shutdown context for server
	wait := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// shutdown the server
	err := server.Serv.Shutdown(ctx)
	return err
}
