package runner

import (
	"bytes"
	ssz "github.com/ferranbt/fastssz"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
	types "ssv-experiments/new_arch_2/spec/asgard/types"
)

// Runner executes a single duty. It receives a RunnerDuty
type Runner struct {
	State *types.State
	// share is the share for the runner with which messages are verified and signed
	Share *types.Share
}

func ProcessMessage(state *types.State, share *types.Share, message *types.SignedPartialSignatureMessages) error {
	switch message.Message.Type {
	case types.PostConsensusPartialSig:
		return ProcessPostConsensus(state, share, message)
	case types.RandaoPartialSig:
		return ProcessRandao(state, share, message)
	case types.SelectionProofPartialSig:
		return errors.New("not supported")
	case types.ContributionProofs:
		return errors.New("not supported")
	case types.ValidatorRegistrationPartialSig:
		return errors.New("not supported")
	default:
		return errors.New("unknown message type")
	}
}

func ProcessPostConsensus(state *types.State, share *types.Share, message *types.SignedPartialSignatureMessages) error {
	// not processing messages if finished
	if Finished(state, share) {
		return errors.New("state finished")
	}

	cd := DecidedConsensusData(state)
	// can't process post consensus if consensus didn't decide
	if cd == nil {
		return errors.New("not decided")
	}

	// verify expected roots
	roots, err := ExpectedPostConsensusRoots(state)
	if err != nil {
		return err
	}
	if err := verifyExpectedRoots(message, roots, cd.Duty.DomainData); err != nil {
		return err
	}

	// verify signatures
	if err := verifyPartialSigMsgForSlot(share, message, DecidedConsensusData(state).Duty.Slot); err != nil {
		return err
	}

	// add message
	state.PartialSignatures = append(state.PartialSignatures, message)

	return nil
}

func ExpectedPostConsensusRoots(state *types.State) ([]ssz.HashRoot, error) {
	switch Role(state) {
	case types.BeaconRoleAttester:
		return AttesterExpectedPostConsensusRoots(state)
	case types.BeaconRoleProposer:
		return ProposerExpectedPostConsensusRoots(state)
	default:
		return nil, errors.New("unknown role")
	}
}

func HasPreConsensusQuorum(state *types.State, share *types.Share) bool {
	all := AllPreConsensus(state)
	return len(all) >= int(share.Quorum)
}

func HasPostConsensusQuorum(state *types.State, share *types.Share) bool {
	all := AllPostConsensus(state)
	return len(all) >= int(share.Quorum)
}

// Finished returns true if finished post consensus phase
func Finished(state *types.State, share *types.Share) bool {
	if DecidedConsensusData(state) == nil {
		return false
	}

	return HasPostConsensusQuorum(state, share)
}

// verifyPartialSigMsgForSlot verifies the message and beacon signatures, returns nil if valid
func verifyPartialSigMsgForSlot(
	share *types.Share,
	signedMsg *types.SignedPartialSignatureMessages,
	slot uint64,
) error {
	if err := signedMsg.Validate(); err != nil {
		return errors.Wrap(err, "SignedPartialSignatureMessage invalid")
	}

	if signedMsg.Message.Slot != slot {
		return errors.New("invalid partial sig slot")
	}

	if err := types.VerifyObjectSignature(signedMsg.Signature, signedMsg, share.Domain, types.PartialSignatureType, share.Cluster); err != nil {
		return errors.Wrap(err, "failed to verify PartialSignature")
	}

	for _, msg := range signedMsg.Message.Signatures {
		if err := verifyBeaconPartialSignature(signedMsg.Signer, share, msg); err != nil {
			return errors.Wrap(err, "could not verify Beacon partial Signature")
		}
	}

	return nil
}

func verifyBeaconPartialSignature(signer uint64, share *types.Share, msg *types.PartialSignatureMessage) error {
	signature := msg.Signature
	root := msg.Root

	for _, n := range share.Cluster {
		if n.Signer == signer {
			pk := &bls.PublicKey{}
			if err := pk.Deserialize(n.PubKey); err != nil {
				return errors.Wrap(err, "could not deserialized pk")
			}
			sig := &bls.Sign{}
			if err := sig.Deserialize(signature[:]); err != nil {
				return errors.Wrap(err, "could not deserialized Signature")
			}

			// verify
			if !sig.VerifyByte(pk, root[:]) {
				return errors.New("wrong signature")
			}
			return nil
		}
	}
	return errors.New("unknown signer")
}

// verifyExpectedRoots verifies signed roots equal to expected roots, requires order
func verifyExpectedRoots(
	signedMessage *types.SignedPartialSignatureMessages,
	expectedRootObjs []ssz.HashRoot,
	domainData [32]byte,
) error {
	if len(expectedRootObjs) != len(signedMessage.Message.Signatures) {
		return errors.New("wrong expected roots count")
	}

	// convert expected roots to map and mark unique roots when verified
	expectedRoots, err := func(expectedRootObjs []ssz.HashRoot) ([][32]byte, error) {
		ret := make([][32]byte, 0)
		for _, rootI := range expectedRootObjs {
			r, err := types.ComputeETHSigningRoot(rootI, domainData)
			if err != nil {
				return nil, errors.Wrap(err, "could not compute ETH signing root")
			}
			ret = append(ret, r)
		}
		return ret, nil
	}(expectedRootObjs)
	if err != nil {
		return err
	}

	roots := func(msgs types.PartialSignatureMessages) [][32]byte {
		ret := make([][32]byte, 0)
		for _, msg := range msgs.Signatures {
			ret = append(ret, msg.Root)
		}
		return ret
	}(signedMessage.Message)

	// verify roots
	for i, r := range roots {
		if !bytes.Equal(expectedRoots[i][:], r[:]) {
			return errors.New("wrong signing root")
		}
	}
	return nil
}
