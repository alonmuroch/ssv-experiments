package common

//go:generate rm -f ./crypto_key_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path crypto_key.go
