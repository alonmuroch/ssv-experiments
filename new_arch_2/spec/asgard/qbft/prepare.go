package qbft

import (
	"bytes"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func UponPrepare(state *types.QBFT, share *types.Share, signedMessage *types.QBFTSignedMessage) error {
	if !uniqueSingerRound(state, signedMessage) {
		return errors.New("duplicate message")
	}

	AddMessage(state, signedMessage)

	if PrepareQuorum(state, share) {
		state.PreparedRound = state.Round
	}

	return nil
}

// CreatePrepareMessage returns unsigned prepare message
func CreatePrepareMessage(state *types.QBFT) (*types.QBFTMessage, error) {
	// TODO implement
	return &types.QBFTMessage{
		Round:   state.Round,
		MsgType: types.PrepareMessageType,
	}, nil
}

func PrepareQuorum(state *types.QBFT, share *types.Share) bool {
	all := RoundAndType(state, state.Round, types.PrepareMessageType)
	return UniqueSignerQuorum(share.Quorum, all)
}

func validSignedPrepareForHeightRoundAndRoot(
	share *types.Share,
	signedMessage *types.QBFTSignedMessage,
	height, round uint64,
	root [32]byte,
) error {
	if signedMessage.Message.MsgType != types.PrepareMessageType {
		return errors.New("prepare msg type is wrong")
	}
	if signedMessage.Message.Height != height {
		return errors.New("wrong msg height")
	}
	if signedMessage.Message.Round != round {
		return errors.New("wrong msg round")
	}

	if err := signedMessage.Validate(); err != nil {
		return errors.Wrap(err, "prepareData invalid")
	}

	if !bytes.Equal(signedMessage.Message.Root[:], root[:]) {
		return errors.New("proposed data mistmatch")
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

	return nil
}
