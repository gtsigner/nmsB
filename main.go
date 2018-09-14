package main

import (
	"log"

	"./cmd"
	"./cmd/execute"
)

var (
	VERSION string
	RELEASE string
)

func main2() {
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
