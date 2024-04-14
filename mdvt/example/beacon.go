package example

import (
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/herumi/bls-eth-go-binary/bls"
)

type BeaconAPI interface {
	GetBlock(slot uint64, randaoSig []byte) (*deneb.BeaconBlock, error)
	SubmitBlock(block *deneb.BeaconBlock, signature []byte) error
}

var (
	DomainProposer                    = [4]byte{0x00, 0x00, 0x00, 0x00}
	DomainAttester                    = [4]byte{0x01, 0x00, 0x00, 0x00}
	DomainRandao                      = [4]byte{0x02, 0x00, 0x00, 0x00}
	DomainDeposit                     = [4]byte{0x03, 0x00, 0x00, 0x00}
	DomainVoluntaryExit               = [4]byte{0x04, 0x00, 0x00, 0x00}
	DomainSelectionProof              = [4]byte{0x05, 0x00, 0x00, 0x00}
	DomainAggregateAndProof           = [4]byte{0x06, 0x00, 0x00, 0x00}
	DomainSyncCommittee               = [4]byte{0x07, 0x00, 0x00, 0x00}
	DomainSyncCommitteeSelectionProof = [4]byte{0x08, 0x00, 0x00, 0x00}
	DomainContributionAndProof        = [4]byte{0x09, 0x00, 0x00, 0x00}
	DomainApplicationBuilder          = [4]byte{0x00, 0x00, 0x00, 0x01}
)

func SlotToEpoch(slot uint64) uint64 {
	panic("implement")
}

func ReconstructPartialSignatures(in [][]byte) (*bls.Sign, error) {
	panic("implement")
}
