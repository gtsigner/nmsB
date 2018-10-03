package message

type DebugType string

const (
	DebugError   DebugType = "error"
	DebugWarning DebugType = "warn"
	DebugInfo    DebugType = "info"
	DebugDebug   DebugType = "debug"
)

type DebugMessage struct {
	*Message
	DebugType *DebugType
	Text      *string
}
