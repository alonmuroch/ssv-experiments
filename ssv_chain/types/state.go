package types

type State struct {
	/*  Chain */
	// Domain represents: byte[0] empty, byte[1] empty, byte[2] network ID, byte[3] fork ID
	Domain                [4]byte `ssz-size:"4"`
	BlockHeight           uint64
	LatestBlockHeaderHash [32]byte     `ssz-size:"32"`
	Validators            []*Validator `ssz-max:"128"`
	/*  Registry */
	Accounts  []*Account  `ssz-max:"65536"`   // 2^16
	Clusters  []*Cluster  `ssz-max:"1048576"` // 2^20
	Operators []*Operator `ssz-max:"65536"`   // 2^16
	Modules   []*Module   `ssz-max:"65536"`   // 2^16
}

func (s *State) AccountByAddress(address []byte) *Account {
	panic("implement")
}

func (s *State) ModuleByID(id uint64) *Module {
	panic("implement")
}
