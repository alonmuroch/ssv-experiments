package types

const (
	OP_V0 = 0x0
	OP_V1 = 0x1
	OP_V2 = 0x2
	OP_V3 = 0x3
	OP_V4 = 0x4
	OP_V5 = 0x5
)

const (
	OP_Module   = 0x0
	OP_Cluster  = 0x1
	OP_Operator = 0x2
	OP_Account  = 0x3
)

const (
	OP_Add       = 0x0
	OP_Remove    = 0x1
	OP_ChangeFee = 0x2
	OP_Deposit   = 0x3
	OP_Withdraw  = 0x4
)

type Operation struct {
	// Type of transaction: byte[0] empty, byte[1] operation type (module, cluster, etc), byte[2] sub-operation type (add, remove, etc), byte[3] version
	Type          [4]byte `ssz-size:"4"`
	OperationData []byte  `ssz-max:"2048"`
}

type Transaction struct {
	Address    []byte `ssz-max:"128"`
	Nonce      uint64
	MaxGas     uint64
	GasPrice   uint64
	Operations []*Operation `ssz-max:"128"`
}

type SignedTransaction struct {
	Signature   []byte `ssz-size:"1024"`
	Transaction Transaction
}

// StatelessCheckTransaction returns nil if transaction passes stateless validation (e.g. types, parsing, etc)
func StatelessCheckTransaction(tx *Transaction) error {
	panic("implement")
}
