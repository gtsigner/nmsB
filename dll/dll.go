package dll

import (
	"../config"
	"../message/json"
	"./communication/context"
	"./communication/dispatch"
	"./http"
	"./instance"
)

func loadConfig(parameter string) (*config.Config, error) {
	cfg := &config.Config{}
	err := json.Decode(parameter, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func start(dllInstance *instance.DllInstance) error {
	// create the client
	client := http.CreateClient(dllInstance.Config)
	// create the dispatch context
	dllInstance.DispatchContext = context.CreateDispatchContext(dllInstance.Version, dllInstance.Release, dllInstance.Config, client)
	// create the dispatcher
	dllInstance.Dispatcher = dispatch.CreateDispacther(dllInstance.DispatchContext)

	// start the client
	err := client.Init()
	if err != nil {
		return err
	}

	return nil
}

func shutdown(dllInstance *instance.DllInstance) error {
	// close the dispatcher
	if dllInstance.Dispatcher != nil {
		dllInstance.Dispatcher.Close()
	}

	// close the client
	if dllInstance.Client != nil {
		dllInstance.Client.Close()
	}

	return nil
}

func Run(version string, release string, parameter string) error {
	// create the dll instance
	dllInstance := instance.NewDllInstance(version, release)

	// load the config from parameters
	cfg, err := loadConfig(parameter)
	if err != nil {
		return err
	}
	dllInstance.Config = cfg

	// start the dll
	err = start(dllInstance)
	if err != nil {
		return err
	}

	// wait for the shutdown
	<-dllInstance.DispatchContext.Shutdown

	// shutdown the dll
	err = shutdown(dllInstance)
	return err
}
