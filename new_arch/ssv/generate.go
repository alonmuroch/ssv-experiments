package ssv

//go:generate rm -f ./runner_state_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path runner_state.go --include ./container.go,../types/duty.go,../types/consensus_data.go,../types/partial_signature_message.go --exclude-objs PartialSignatureContainer

//go:generate rm -f ./runner_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path runner.go --include ../p2p/message.go,../types/share.go,./runner_state.go,./container.go,../types/duty.go,../types/consensus_data.go,../types/partial_signature_message.go --exclude-objs PartialSignatureContainer
