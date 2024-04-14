package module

// Operator represents a registered operator to a module's cluster
type Operator struct {
	// ID for the operator
	ID uint64
	// PubKey for the operator's share
	PubKey []byte `ssz-max:"1024"`
}

type Cluster []*Operator
