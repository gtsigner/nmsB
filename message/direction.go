package message

/*MessageDirection defines the direction of the message flow*/
type MessageDirection string

const (
	/*ClientToDll is a message direction from client to dll */
	ClientToDll MessageDirection = "c2d"
	/*DllToClient is a message direction from dll to client */
	DllToClient MessageDirection = "d2c"
	/*DllToClients is a message direction from dll to all clients */
	DllToClients MessageDirection = "d2cs"
)