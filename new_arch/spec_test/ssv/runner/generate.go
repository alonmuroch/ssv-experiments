package runner

//go:generate rm -f ./test_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path test.go --include ../../../ssv/runner.go,../../../ssv/runner_state.go,../../../ssv/container.go,../../../p2p/message.go,../../../types/share.go,../../../types/partial_signature_message.go,../../../types/duty.go,../../../types/consensus_data.go --exclude-objs PartialSignatureContainer,Identifier
