package dsp

type Network interface {
	Broadcast(payload []byte) error
	Receive() <-chan []byte
}

type Store interface {
	Get(key [32]byte) ([]byte, error)
	Set(key [32]byte, value []byte) error
}

type Consensus interface {
	Decide(inputValue []byte, valueCheck func(value []byte) error) (decidedValue []byte, proof interface{}, err error)
}

type TSS interface {
	SignAndAggregate(inputValue []byte) (signature []byte, signers []uint64, err error)
}

type SecureSigner interface {
	Sign(inputValue []byte) ([]byte, error)
}

type DVT interface {
	DecideAndSign(
		inputValue []byte,
		valueCheck func(value []byte) error,
	) (decidedValue []byte, proof interface{}, signature []byte, signers []uint64, err error)
}

type API struct {
	// core methods
	Network
	Store
	Consensus

	// signing methods
	TSS
	SecureSigner
	DVT
}
