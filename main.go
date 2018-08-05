package main

import (
	"./cmd"
	"./cmd/execute"
	"log"
)

var (
	VERSION string
	RELEASE string
)

func main() {
	// parse the CMD
	args, err := cmd.Parse()
	if err != nil {
		log.Panicln(err)
	}

	// check if help given
	if *args.Help {
		cmd.Help()
		return
	}

	err = execute.Execute(args)
	if err != nil {
		log.Panicln(err)
	}

}
