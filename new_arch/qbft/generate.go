package qbft

//go:generate rm -f ./messages_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path messages.go --include ../p2p/message.go --exclude-objs Identifier

//go:generate rm -f ./input_data_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path input_data.go

//go:generate rm -f ./state_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path state.go --include ./container.go,./input_data.go,../p2p/message.go,./messages.go,../types/share.go --exclude-objs Identifier

//go:generate rm -f ./instance_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path instance.go --include ./state.go,../types/share.go,../p2p/message.go,./input_data.go,./container.go,./messages.go
