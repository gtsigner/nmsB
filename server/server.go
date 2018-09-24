package server

import (
	"log"

	"../config"
	"../sys/signal"
	"./http"
	"./instance"
)

func start(instance *instance.ServerInstance) error {
	// load configuration
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	instance.Config = cfg

	// starting http server
	httpServer, err := http.RunHttpServer(cfg.Http, nil)
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
