package qbft

//go:generate rm -f ./process_message_spectest_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path process_message_spectest.go --include ../../../../spec/asgard/types/state.go,../../../../spec/asgard/types/qbft_signed_message.go,../../../../spec/asgard/types/consensus_data.go,../../../../spec/asgard/types/duty.go,../../../../spec/asgard/types/partial_signature_message.go
