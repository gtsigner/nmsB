package nms

import (
	"../memory"
	"./config"
)

type Instance struct {
	config *config.Config
	reader *memory.MemoryReader
}

func OpenInstance(config *config.Config) (*Instance, error) {
	reader := memory.NewMemoryReader()
	err := reader.OpenByName(*config.ProcessName)
	if err != nil {
		return nil, err
	}

	return &Instance{
		config: config,
		reader: reader,
	}
}
