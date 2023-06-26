package types

import (
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

type ConsensusData struct {
	// Duty max size is
	// 			8 + 48 + 6*8 + 13*8 = 208 ~= 2^8
	Duty        *Duty
	DataVersion uint64
	// PreConsensusJustifications max size is
	//			13*SignedPartialSignatureMessage(2^16) ~= 2^20
	PreConsensusJustification []*SignedPartialSignatureMessages `ssz-max:"13"`
	// DataSSZ has max size as following
	// Biggest object is a full beacon block
	// BeaconBlock is 2*32+2*8 + BeaconBlockBody
	// BeaconBlockBody is
	//			96 + ETH1Data(2*32+8) + 32 +
	//			16*ProposerSlashing(2*SignedBeaconBlockHeader(96 + 3*32 + 2*8)) +
	//			2*AttesterSlashing(2*IndexedAttestation(2048*8 + 96 + AttestationData(2*8 + 32 + 2*(8+32)))) +
	//			128*Attestation(2048*8 + 96 + AttestationData(2*8 + 32 + 2*(8+32))) +
	//			16*Deposit(33*32 + 48 + 32 + 8 + 96) +
	//			16*SignedVoluntaryExit(96 + 2*8) +
	//			SyncAggregate(64 + 96) +
	//			ExecutionPayload(32 + 20 + 2*32 + 256 + 32 + 4*8 + 3*32 + 1048576*1073741824)
	// = 2^21(everything but transactions) + 2^50 (transaction list)
	// We do not need to support such a big DataSSZ size as 2^50 represents 1000X the actual block gas limit
	// Current 30M gas limit produces 30M / 16 (call data cost) = 1,875,000 bytes (https://eips.ethereum.org/EIPS/eip-4488)
	// we can upper limit transactions to 2^21, together with the rest of the data 2*2^21 = 2^22 = 4,194,304 bytes
	// Exmplanation on why transaction sizes are so big https://github.com/ethereum/consensus-specs/pull/2686
	DataSSZ []byte `ssz-max:"4194304"` // 2^22
}

func (cd *ConsensusData) GetSigningRoot() ([32]byte, error) {
	switch cd.Duty.Role {
	case BeaconRoleAttester:
		// TODO version support
		ret := &phase0.AttestationData{}
		if err := ret.UnmarshalSSZ(cd.DataSSZ); err != nil {
			return [32]byte{}, err
		}
		return ret.HashTreeRoot()
	case BeaconRoleProposer:
		// TODO version support
		ret := &capella.BeaconBlock{}
		if err := ret.UnmarshalSSZ(cd.DataSSZ); err != nil {
			return [32]byte{}, err
		}
		return ret.HashTreeRoot()
	default:
		return [32]byte{}, errors.New("unknown duty role")
	}
}
