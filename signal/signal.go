package signal

import (
	"os"
	"os/signal"
)

type Signal struct {
	interruptSignal chan os.Signal
}

func NewSignal() *Signal {
	return &Signal{
		interruptSignal: make(chan os.Signal, 1),
	}
}

func (s *Signal) Init() {
	signal.Notify(s.interruptSignal, os.Interrupt)
}

func (signal *Signal) Await() {
	<-signal.interruptSignal
}
