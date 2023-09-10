package fixtures

import (
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// SignedPartialSignatureMessage returns a signed partial signature
func SignedPartialSignatureMessage(signer, slot uint64, msgType types.PartialSigMsgType) *types.SignedPartialSignatureMessages {
	return &types.SignedPartialSignatureMessages{
		Message: types.PartialSignatureMessages{
			Type: msgType,
			Slot: slot,
		},
		Signer: signer,
	}
}
