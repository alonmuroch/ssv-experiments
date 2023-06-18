package process

//go:generate rm -f ./test_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path test.go --include ../../state.go,../../input_data.go,../../container.go,../../messages.go,../../../p2p/message.go,../../../types/share.go --exclude-objs Identifier
