package qbft

import (
	"bytes"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// UponProposal process proposal message
// Assumes proposal message is valid!
func UponProposal(state *types.QBFT, signedMessage *types.QBFTSignedMessage) error {
	AddMessage(state, signedMessage)
	newRound := signedMessage.Message.Round
	state.ProposalAcceptedForCurrentRound = signedMessage
	state.Round = newRound
	return nil
}

func CreateProposalMessage(state *types.QBFT) (*types.QBFTMessage, error) {
	// TODO implement
	return &types.QBFTMessage{
		Round:   state.Round,
		MsgType: types.ProposalMessageType,
	}, nil
}

// isValidProposal returns nil if message is valid for state
func isValidProposal(state *types.QBFT, share *types.Share, signedMessage *types.QBFTSignedMessage) error {
	if signedMessage.Message.MsgType != types.ProposalMessageType {
		return errors.New("msg type is not proposal")
	}
	if signedMessage.Message.Height != state.Height {
		return errors.New("wrong msg height")
	}
	if len(signedMessage.Signers) != 1 {
		return errors.New("proposal allows 1 signer")
	}
	if err := types.VerifyObjectSignature(
		signedMessage.Signature,
		signedMessage,
		share.Domain,
		types.QBFTSignatureType,
		share.Cluster,
	); err != nil {
		return err
	}
	if signedMessage.Signers[0] != proposerForRound(signedMessage.Message.Round) {
		return errors.New("proposal leader invalid")
	}
	if err := signedMessage.Validate(); err != nil {
		return errors.Wrap(err, "proposal invalid")
	}

	// verify full data integrity
	r, err := HashDataRoot(signedMessage.FullData)
	if err != nil {
		return errors.Wrap(err, "could not hash input data")
	}
	if !bytes.Equal(signedMessage.Message.Root[:], r[:]) {
		return errors.New("H(data) != root")
	}

	// get justifications
	roundChangeJustification, _ := signedMessage.Message.GetRoundChangeJustifications() // no need to check error, checked on signedProposal.Validate()
	prepareJustification, _ := signedMessage.Message.GetPrepareJustifications()         // no need to check error, checked on signedProposal.Validate()

	if err := isProposalJustification(
		state,
		share,
		roundChangeJustification,
		prepareJustification,
		state.Height,
		signedMessage.Message.Round,
		signedMessage.FullData,
	); err != nil {
		return errors.Wrap(err, "proposal not justified")
	}

	if (state.ProposalAcceptedForCurrentRound == nil && signedMessage.Message.Round == state.Round) ||
		signedMessage.Message.Round > state.Round {
		return nil
	}
	return errors.New("proposal is not valid with current state")
}

func isProposalJustification(
	state *types.QBFT,
	share *types.Share,
	roundChangeMessages, prepareMessages []*types.QBFTSignedMessage,
	height, round uint64,
	fullData []byte) error {

	if round == types.FirstRound {
		return nil
	} else {
		// check all round changes are valid for height and round
		// no quorum, duplicate signers,  invalid still has quorum, invalid no quorum
		// prepared
		for _, rc := range roundChangeMessages {
			if err := validRoundChangeForData(state, share, rc, height, round, fullData); err != nil {
				return errors.Wrap(err, "change round msg not valid")
			}
		}

		// check there is a quorum
		if !UniqueSignerQuorum(share.Quorum, roundChangeMessages) {
			return errors.New("change round has no quorum")
		}

		// previouslyPreparedF returns true if any on the round change messages have a prepared round and fullData
		previouslyPrepared, err := func(rcMsgs []*types.QBFTSignedMessage) (bool, error) {
			for _, rc := range rcMsgs {
				if rc.Message.RoundChangePrepared() {
					return true, nil
				}
			}
			return false, nil
		}(roundChangeMessages)
		if err != nil {
			return errors.Wrap(err, "could not calculate if previously prepared")
		}

		if !previouslyPrepared {
			return nil
		} else {

			// check prepare quorum
			if !UniqueSignerQuorum(share.Quorum, prepareMessages) {
				return errors.New("prepares has no quorum")
			}

			// get a round change data for which there is a justification for the highest previously prepared round
			rcm, err := highestPrepared(roundChangeMessages)
			if err != nil {
				return errors.Wrap(err, "could not get highest prepared")
			}
			if rcm == nil {
				return errors.New("no highest prepared")
			}

			// proposed fullData must equal highest prepared fullData
			r, err := HashDataRoot(fullData)
			if err != nil {
				return errors.Wrap(err, "could not hash input data")
			}
			if !bytes.Equal(r[:], rcm.Message.Root[:]) {
				return errors.New("proposed data doesn't match highest prepared")
			}

			// validate each prepare message against the highest previously prepared fullData and round
			for _, pm := range prepareMessages {
				if err := validSignedPrepareForHeightRoundAndRoot(
					share,
					pm,
					height,
					rcm.Message.DataRound,
					rcm.Message.Root,
				); err != nil {
					return errors.New("signed prepare not valid")
				}
			}
			return nil
		}
	}
}
