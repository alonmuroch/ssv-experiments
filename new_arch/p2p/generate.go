package p2p

//go:generate rm -f ./message_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path message.go  --exclude-objs Identifier,MsgType
