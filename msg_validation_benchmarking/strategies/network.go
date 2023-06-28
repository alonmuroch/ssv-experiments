package strategies

type INetwork interface {
	Start()
	GetMessagesChannel() chan []byte
	GetStopChannel() chan bool
}
