package types

const (
	OP_V0 = 0x0
)

const (
	// OP_Sys are operations generated by chain validators, are not signed and aim to carry out system operations (deposits, update cluster instances, etc.)
	OP_Sys = 0x0
	// OP_User are operations generated by users, signed and aim to carry user controlled operations (registrations, withdrawals, etc.)
	OP_User = 0x1
)

const (
	OP_Module    = 0x0
	OP_Cluster   = 0x1
	OP_Operator  = 0x2
	OP_Account   = 0x3
	OP_Stake     = 0x4
	OP_Validator = 0x5
)

const (
	OP_Add       = 0x0
	OP_Remove    = 0x1
	OP_Modify    = 0x2
	OP_ChangeFee = 0x3
	OP_Deposit   = 0x4
	OP_Withdraw  = 0x5
	OP_Lock      = 0x6
	OP_Release   = 0x7
	OP_Delegate  = 0x8
)

type Operation struct {
	// Type of transaction: byte[0] tx origin (system, user), byte[1] operation type (module, cluster, etc), byte[2] sub-operation type (add, remove, etc), byte[3] version
	Type          [4]byte `ssz-size:"4"`
	OperationData []byte  `ssz-max:"2048"`
}

type Transaction struct {
	// Address sending the transaction
	Address    []byte `ssz-max:"128"`
	Nonce      uint64
	MaxGas     uint64
	GasPrice   uint64       // in VGBit
	Operations []*Operation `ssz-max:"128"`
}

type SignedTransaction struct {
	Signature []byte `ssz-size:"1024"`
	// Signer is the L1 account that signed the transaction
	Transaction Transaction
}
