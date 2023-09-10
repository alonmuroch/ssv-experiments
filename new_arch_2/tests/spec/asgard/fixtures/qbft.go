package fixtures

import (
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

func QBFTSignedMessage(signer, round, msgType uint64) *types.QBFTSignedMessage {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)
	sk := bls.SecretKey{}
	sk.SetByCSPRNG()
	fmt.Printf("sk: %s\n", sk.SerializeToHexStr())

	var fullData []byte
	root := [32]byte{}
	if msgType == types.ProposalMessageType {
		cd := AttesterConsensusData
		fullData, _ = cd.MarshalSSZ()
		root, _ = cd.HashTreeRoot()
	}

	msg := types.QBFTMessage{
		Round:   round,
		MsgType: msgType,
		Root:    root,
	}
	r, _ := msg.HashTreeRoot()
	sigByts := sk.SignByte(r[:])

	sig := [96]byte{}
	copy(sig[:], sigByts.Serialize())

	ss := &bls.Sign{}
	ss.Deserialize(sig[:])

	return &types.QBFTSignedMessage{
		Message: types.QBFTMessage{
			Round:   round,
			MsgType: msgType,
			Root:    root,
		},
		Signature: sig,
		Signers:   []uint64{signer},
		FullData:  fullData,
	}
}
