package ssv

//go:generate rm -f ./partial_sig_container_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path partial_sig_container.go --include ../types/partial_signature_message.go

//go:generate rm -f ./runner_state_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path runner_state.go --include ./partial_sig_container.go,../types/duty.go,../types/consensus_data.go,../types/partial_signature_message.go --exclude-objs PartialSignatureContainer
