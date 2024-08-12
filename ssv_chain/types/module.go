package types

// Module represents a DVT specific application
type Module struct {
	// Account that controls the Module
	Account uint64
	ID      uint64
	Name    []byte `ssz-max:"1024"`
}
