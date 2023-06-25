package qbft

//go:generate rm -f ./messages_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path messages.go --include ../types/duty.go,../types/partial_signature_message.go,../types/consensus_data.go,../p2p/message.go --exclude-objs Identifier

//go:generate rm -f ./state_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path state.go --include ../types/duty.go,../types/partial_signature_message.go,../types/consensus_data.go,./container.go,../p2p/message.go,./messages.go,../types/share.go --exclude-objs Identifier

//go:generate rm -f ./instance_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path instance.go --include ./state.go,../types/duty.go,../types/partial_signature_message.go,../types/consensus_data.go,../types/share.go,../p2p/message.go,./container.go,./messages.go
