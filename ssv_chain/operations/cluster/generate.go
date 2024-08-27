package cluster

//go:generate rm -f ./v0_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path v0.go --include ../../types/cluster.go,../../common/crypto.go
