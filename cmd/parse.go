package cmd

import (
	"flag"
)

func Parse() (*Arguments, error) {
	help := flag.Bool("help", false, "Print help")
	pointer := flag.Bool("pointer", false, "Pointer Mode")
	reader := flag.Bool("reader", false, "Reader Mode")
	server := flag.Bool("server", false, "Server Mode")
	address := flag.String("address", "", "Pointer as HEX")
	processId := flag.Int("process-id", -1, "Process Id")

	arguments := &Arguments{
		Help:      help,
		Pointer:   pointer,
		Server:    server,
		Reader:    reader,
		Address:   address,
		ProcessId: processId,
	}

	flag.Parse()

	return arguments, nil
}
