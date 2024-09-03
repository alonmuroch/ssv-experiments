package types

type State struct {
	/*  Chain */
	// Domain represents: byte[0] empty, byte[1] empty, byte[2] network ID, byte[3] fork ID
	Domain                [4]byte `ssz-size:"4"`
	BlockHeight           uint64
	LatestBlockHeaderHash []byte       `ssz-size:"32"`
	Validators            []*Validator `ssz-max:"128"`
	/*  Registry */
	Accounts  []*Account  `ssz-max:"65536"`   // 2^16
	Clusters  []*Cluster  `ssz-max:"1048576"` // 2^20
	Operators []*Operator `ssz-max:"65536"`   // 2^16
	Modules   []*Module   `ssz-max:"65536"`   // 2^16
}

// AccountByAddress returns account by address, if not found nil
func (s *State) AccountByAddress(address []byte) *Account {
	panic("implement")
}

// CreateAccountForAddress creates and adds account for address if not found
func (s *State) CreateAccountForAddress(address []byte) *Account {
	panic("implement")
}

// ModuleByID returns module by ID or nil if not found
func (s *State) ModuleByID(id uint64) *Module {
	panic("implement")
}

// ClusterByID returns cluster by ID or nil if not found
func (s *State) ClusterByID(id uint64) *Cluster {
	panic("implement")
}

// ValidatorByAddress returns validator by address or nil if not found
func (s *State) ValidatorByAddress(address []byte) *Validator {
	panic("implement")
}

// ValidatorByID returns validator by ID or nil if not found
func (s *State) ValidatorByID(id uint64) *Validator {
	panic("implement")
}
