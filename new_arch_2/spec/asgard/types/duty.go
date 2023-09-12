package types

type Duty struct {
	BeaconNetwork BeaconNetwork `ssz-max:"48"`
	Role          uint64
	ValidatorPK   [48]byte `ssz-size:"48"`
	Slot          uint64
	DomainData    [32]byte `ssz-size:"32"`

	// ValidatorIndex is the index of the validator that should attest.
	ValidatorIndex uint64
	// CommitteeIndex is the index of the committee in which the attesting validator has been placed.
	CommitteeIndex uint64
	// CommitteeLength is the length of the committee in which the attesting validator has been placed.
	CommitteeLength uint64
	// CommitteesAtSlot is the number of committees in the slot.
	CommitteesAtSlot uint64
	// ValidatorCommitteeIndex is the index of the validator in the list of validators in the committee.
	ValidatorCommitteeIndex uint64
	// ValidatorSyncCommitteeIndices is the index of the validator in the list of validators in the committee.
	ValidatorSyncCommitteeIndices []uint64 `ssz-max:"13"`
}
