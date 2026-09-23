package interfaces

type Client interface {
	Updates() []byte
	SendMessage(int64, string)
}
