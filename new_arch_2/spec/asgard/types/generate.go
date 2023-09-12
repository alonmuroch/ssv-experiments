package types

//go:generate rm -f ./p2p_message_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path p2p_message.go  --exclude-objs Identifier,MsgType

//go:generate rm -f ./consensus_data_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path consensus_data.go --include ./beacon.go,./duty.go,./partial_signature_message.go

//go:generate rm -f ./share_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path share.go --include ./signatures.go,./forks.go

//go:generate rm -f ./duty_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path duty.go --include ./beacon.go

//go:generate rm -f ./partial_signature_message_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path partial_signature_message.go --exclude-objs PartialSigMsgType

//go:generate rm -f ./qbft_signed_message_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path qbft_signed_message.go --exclude-objs PartialSigMsgType

//go:generate rm -f ./state_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path state.go --include ./beacon.go,./qbft_signed_message.go,./partial_signature_message.go,./duty.go,./consensus_data.go,./partial_signature_message.go --exclude-objs PartialSignatureContainer
