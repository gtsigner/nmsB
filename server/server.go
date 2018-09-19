package server

import (
	"log"

	"../sys/signal"
	"./http"
	"./instance"
)

func start(instance *instance.ServerInstance) error {
	// starting http server
	httpServer, err := http.RunHttpServer(nil, nil)
	if err != nil {
		return err
	}
	instance.HttpServer = httpServer

	return nil
}

func await() {
	// create the interupe signal
	sig := signal.CreateSignal()
	sig.Await()
}

func shutdown(instance *instance.ServerInstance) error {
	// shutdown the http server
	if instance.HttpServer != nil {
		err := instance.HttpServer.Shutdown()
		if err != nil {
			return err
		}
	}

	return nil
}

func Run(version string, release string) error {
	log.Printf("starting server [ version: %s, release: %s ]", version, release)
	// create the server instance
	instance := instance.NewServerInstance()

	// start the server
	err := start(instance)
	if err != nil {
		return err
	}

	// wait for shutdown
	await()

	log.Printf("Shutdown server...")

	// shutdown the server
	err = shutdown(instance)
	if err != nil {
		return err
	}

	log.Printf("Done.")

	return nil
}
