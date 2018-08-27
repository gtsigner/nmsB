package signal

func CreateSignal() *Signal {
	signal := NewSignal()
	signal.Init()
	return signal
}
