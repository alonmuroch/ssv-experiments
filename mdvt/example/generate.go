package example

//go:generate rm -f ./storage_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path storage.go  --exclude-objs Storage,ClusterData

//go:generate rm -f ./types_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path types.go
