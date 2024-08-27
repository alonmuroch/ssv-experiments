package types

// Module represents a DVT specific application
type Module struct {
	// Network for on which the module operates
	Network [4]byte `ssz-size:"4"`
	// Address that controls the Module
	Address []byte `ssz-max:"128"`
	ID      uint64
	Name    []byte `ssz-max:"1024"`
}
