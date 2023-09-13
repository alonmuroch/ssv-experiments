package qbft

import (
	"bytes"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func UponRoundChange(msg *types.QBFTSignedMessage) error {
	panic("implement")
}

func CreateRoundChangeMessage(state *types.QBFT) (*types.QBFTMessage, error) {
	panic("implement")
}

func RoundChangeQuorum(state *types.QBFT, share *types.Share) bool {
	all := RoundAndType(state, state.Round, types.RoundChangeMessageType)
	return UniqueSignerQuorum(share.Quorum, all)
}

func RoundChangePartialQuorum(state *types.QBFT, share *types.Share) bool {
	all := RoundAndType(state, state.Round, types.RoundChangeMessageType)
	return UniqueSignerQuorum(share.PartialQuorum, all)
}

// validRoundChangeForData returns nil if round change message is valid
func validRoundChangeForData(
	state *types.QBFT,
	share *types.Share,
	signedMessage *types.QBFTSignedMessage,
	round uint64,
	fullData []byte,
) error {
	if signedMessage.Message.MsgType != types.RoundChangeMessageType {
		return errors.New("round change msg type is wrong")
	}
	if signedMessage.Message.Round != round {
		return errors.New("wrong msg round")
	}
	if len(signedMessage.Signers) != 1 {
		return errors.New("msg allows 1 signer")
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

	if err := signedMessage.Validate(); err != nil {
		return errors.Wrap(err, "message invalid")
	}

	// Addition to formal spec
	// We add this extra tests on the msg itself to filter round change msgs with invalid justifications, before they are inserted into msg containers
	if signedMessage.Message.RoundChangePrepared() {
		r, err := HashDataRoot(fullData)
		if err != nil {
			return errors.Wrap(err, "could not hash input data")
		}

		// validate prepare message justifications
		prepareMsgs, _ := signedMessage.Message.GetRoundChangeJustifications() // no need to check error, checked on signedMsg.Message.Validate()
		for _, pm := range prepareMsgs {
			if err := validSignedPrepareForHeightRoundAndRoot(
				share,
				pm,
				state.Height,
				signedMessage.Message.DataRound,
				signedMessage.Message.Root,
			); err != nil {
				return errors.Wrap(err, "round change justification invalid")
			}
		}

		if !bytes.Equal(r[:], signedMessage.Message.Root[:]) {
			return errors.New("H(data) != root")
		}

		// check quorum
		if !UniqueSignerQuorum(share.Quorum, prepareMsgs) {
			return errors.New("no justifications quorum")
		}

		if signedMessage.Message.DataRound > round {
			return errors.New("prepared round > round")
		}

		return nil
	}
	return nil
}

func highestPrepared(roundChangeMessages []*types.QBFTSignedMessage) (*types.QBFTSignedMessage, error) {
	var ret *types.QBFTSignedMessage
	for _, rc := range roundChangeMessages {
		if !rc.Message.RoundChangePrepared() {
			continue
		}

		if ret == nil {
			ret = rc
		} else {
			if ret.Message.DataRound < rc.Message.DataRound {
				ret = rc
			}
		}
	}
	return ret, nil
}
