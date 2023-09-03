package types

type QBFT struct {
	Round  uint64
	Height uint64

	PreparedRound uint64

	// Messages is a unified (to all message type) container slice, simple and easy to serialize.
	// All messages in the container are verified and authenticated
	Messages []*QBFTSignedMessage `ssz-max:"256"`

	// StartValue used to propose (when leader) unless already prepared previously
	StartValue *ConsensusData

	// Stopped when true, can't process any messages
	Stopped bool
}

func (qbft *QBFT) DecidedValue() *ConsensusData {
	panic("implement")
}

type State struct {
	// PartialSignatures holds partial BLS signatures
	PartialSignatures []*SignedPartialSignatureMessages `ssz-max:"256"`
	// DecidedValue holds the decided value set after consensus phase
	QBFT         *QBFT
	StartingDuty *Duty
}

func NewState(duty *Duty) *State {
	return &State{
		StartingDuty:      duty,
		PartialSignatures: []*SignedPartialSignatureMessages{},
	}
}
