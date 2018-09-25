package message

type ServerStatusMessage struct {
	Message
	Version   *string
	Release   *string
	Connected *bool
}

func CreateServerStatusMessage(version string, release string, connected bool) *ServerStatusMessage {

	return &ServerStatusMessage{
		Message:   Message{},
		Version:   &version,
		Release:   &release,
		Connected: &connected,
	}

}
