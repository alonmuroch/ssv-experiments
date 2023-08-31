package types

type PartialSigMsgType uint64

func (t PartialSigMsgType) IsPreConsensusType() bool {
	return t == RandaoPartialSig || t == SelectionProofPartialSig || t == ContributionProofs || t == ValidatorRegistrationPartialSig
}

func (t PartialSigMsgType) IsPostConsensusType() bool {
	return t == PostConsensusPartialSig
}

const (
	// PostConsensusPartialSig is a partial signature over a decided duty (attestation data, block, etc)
	PostConsensusPartialSig PartialSigMsgType = iota
	// RandaoPartialSig is a partial signature over randao reveal
	RandaoPartialSig
	// SelectionProofPartialSig is a partial signature for aggregator selection proof
	SelectionProofPartialSig
	// ContributionProofs is the partial selection proofs for sync committee contributions (it's an array of sigs)
	ContributionProofs
	// ValidatorRegistrationPartialSig is a partial signature over a ValidatorRegistration object
	ValidatorRegistrationPartialSig
)

type PartialSignatureMessage struct {
	Signature [96]byte `ssz-size:"96"`
	Root      [32]byte `ssz-size:"32"`
}

type PartialSignatureMessages struct {
	Type       PartialSigMsgType
	Slot       uint64
	Signatures []*PartialSignatureMessage `ssz-max:"13"`
}

type SignedPartialSignatureMessages struct {
	// Message is at the top for quick identifier look (see docs)
	Message   PartialSignatureMessages
	Signature [96]byte `ssz-size:"96"`
	Signer    uint64
}
