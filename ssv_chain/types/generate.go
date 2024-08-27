package types

//go:generate rm -f ./config_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path config.go

//go:generate rm -f ./balance_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path balance.go

//go:generate rm -f ./transaction_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path transaction.go

//go:generate rm -f ./account_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path account.go  --include ./balance.go

//go:generate rm -f ./cluster_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path cluster.go --include ../common/crypto.go

//go:generate rm -f ./module_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path module.go

//go:generate rm -f ./operator_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path operator.go --include ../common/crypto.go

//go:generate rm -f ./state_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path state.go --include ./validator.go,./account.go,./cluster.go,./operator.go,./module.go,../common/crypto.go,./balance.go

//go:generate rm -f ./validator_encoding.go
//go:generate go run github.com/ferranbt/fastssz/sszgen --path validator.go --include ../common/crypto.go
