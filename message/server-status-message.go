package message

type ServerStatusMessage struct {
	Message
	Version   *string
	Release   *string
	Connected *bool
	Clients   *int
}
