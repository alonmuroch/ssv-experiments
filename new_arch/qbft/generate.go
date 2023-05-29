package qbft

//go:generate rm -f ./state_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path state.go --include ../p2p/message.go,./messages.go,../types/share.go --exclude-objs Identifier

//go:generate rm -f ./messages_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path messages.go --include ../p2p/message.go --exclude-objs Identifier
