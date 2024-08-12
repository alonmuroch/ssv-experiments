package ssv_chain

const (
	V0 = 0x0
	V1 = 0x1
	V2 = 0x2
	V3 = 0x3
	V4 = 0x4
	V5 = 0x5
)

const (
	Module   = 0x0
	Cluster  = 0x1
	Operator = 0x2
	Account  = 0x3
)

const (
	Add       = 0x0
	Remove    = 0x1
	ChangeFee = 0x2
	Deposit   = 0x3
	Withdraw  = 0x4
)

type Transaction struct {
	// Type of transaction: byte[0] empty, byte[1] version, byte[2] operation type (module, cluster, etc), byte[3] sub-operation type (add, remove, etc)
	Type          [4]byte
	OperationData []byte
}

type SignedTransaction struct {
	Signature   []byte
	Transaction Transaction
}

// StatelessCheckTransaction returns nil if transaction passes stateless validation (e.g. types, parsing, etc)
func StatelessCheckTransaction(tx *Transaction) error {
	panic("implement")
}
