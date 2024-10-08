package types

//go:generate rm -f ./consensus_data_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path consensus_data.go --include ./duty.go,./partial_signature_message.go,../p2p/message.go

//go:generate rm -f ./share_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path share.go

//go:generate rm -f ./duty_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path duty.go

//go:generate rm -f ./partial_signature_message_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path partial_signature_message.go --include ../p2p/message.go --exclude-objs PartialSigMsgType
