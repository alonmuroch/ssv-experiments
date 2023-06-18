package runner

//go:generate rm -f ./test_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path test.go --include ../../runner_state.go,../../container.go,../../../p2p/message.go,../../../types/share.go,../../../types/partial_signature_message.go,../../../types/duty.go,../../../types/consensus_data.go --exclude-objs PartialSignatureContainer,Identifier
