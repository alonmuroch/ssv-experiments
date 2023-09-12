package types

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

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

	DomainError = [4]byte{0x99, 0x99, 0x99, 0x99}
)

const (
	BeaconRoleAttester = iota
	BeaconRoleAggregator
	BeaconRoleProposer
)

// Available networks.
var (
	// PraterNetwork represents the Prater test network.
	PraterNetwork BeaconNetwork = []byte("prater")

	// MainNetwork represents the main network.
	MainNetwork BeaconNetwork = []byte("mainnet")

	// BeaconTestNetwork is a simple test network with a custom genesis time
	BeaconTestNetwork BeaconNetwork = []byte("now_test_network")
)

// BeaconNetwork represents the network.
type BeaconNetwork []byte

// EstimatedEpochAtSlot estimates epoch at the given slot
func (n BeaconNetwork) EstimatedEpochAtSlot(slot uint64) uint64 {
	return slot / n.SlotsPerEpoch()
}

// SlotsPerEpoch returns number of slots per one epoch
func (n BeaconNetwork) SlotsPerEpoch() uint64 {
	return 32
}

func ComputeETHSigningRoot(obj ssz.HashRoot, domain [32]byte) ([32]byte, error) {
	root, err := obj.HashTreeRoot()
	if err != nil {
		return [32]byte{}, err
	}
	signingContainer := phase0.SigningData{
		ObjectRoot: root,
		Domain:     domain,
	}
	return signingContainer.HashTreeRoot()
}
