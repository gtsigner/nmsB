package server

import (
	"log"

	"../config"
	"../sys/signal"
	"./dispatch"
	"./dispatch/context"
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

	// create the dispatch context
	dispatchContext := context.CreateDispatchContext(instance.Version, instance.Release)

	// create the dispatcher
	instance.Dispatcher = dispatch.CreateDispacther(dispatchContext)

	// starting http server
	httpServer, err := http.RunHttpServer(cfg.Http, dispatchContext.WebSocketManager)
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

	// shutdown the dispatcher
	if instance.Dispatcher != nil {
		instance.Dispatcher.Close()
	}

	return nil
}

func Run(version string, release string) error {
	log.Printf("starting server [ version: %s, release: %s ]", version, release)
	// create the server instance
	instance := instance.NewServerInstance(version, release)

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
