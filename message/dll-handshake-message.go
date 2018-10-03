package message

type DllHandshakeMessage struct {
	*Message
	Version *string
	Release *string
}
