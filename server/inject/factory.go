package inject

import (
	"sync"

	"../../config"
)

func CreateInjector(cfg *config.Config) *Injector {
	injector := &Injector{
		config: cfg,
		lock:   sync.Mutex{},
	}

	return injector
}
