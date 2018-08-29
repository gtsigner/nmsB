package nms

import (
	"../memory"
	cfg "./config"
)

type Instance struct {
	config *cfg.Config
	reader *memory.MemoryReader
}

func OpenInstance(config *cfg.Config) (*Instance, error) {
	reader := memory.NewMemoryReader()
	err := reader.OpenByName(*config.ProcessName)
	if err != nil {
		return nil, err
	}

	return &Instance{
		config: config,
		reader: reader,
	}, nil
}
