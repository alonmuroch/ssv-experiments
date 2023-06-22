package process

//go:generate rm -f ./test_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path test.go --include ../../../qbft/instance.go,../../../qbft/state.go,../../../qbft/input_data.go,../../../qbft/container.go,../../../qbft/messages.go,../../../p2p/message.go,../../../types/share.go --exclude-objs Identifier
