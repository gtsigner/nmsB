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
	log.Println("loading configuration files")
	// load configuration
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	instance.Config = cfg

	log.Println("initiating server dispatcher")
	// create the dispatch context
	dispatchContext := context.CreateDispatchContext(instance.Version, instance.Release, cfg)
	// create the dispatcher
	instance.Dispatcher = dispatch.CreateDispacther(dispatchContext)

	log.Println("starting http web-server")
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
		log.Println("shutdown http web-server")
		err := instance.HttpServer.Shutdown()
		if err != nil {
			return err
		}
	}

	// shutdown the dispatcher
	if instance.Dispatcher != nil {
		log.Println("closing server dispatcher")
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
