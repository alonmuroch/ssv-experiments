package fixtures

import (
	"ssv-experiments/new_arch/types"
)

// PartialSignatureMessage returns a signed partial signature
func PartialSignatureMessage(signer, slot uint64, msgType types.PartialSigMsgType) *types.SignedPartialSignatureMessages {
	return &types.SignedPartialSignatureMessages{
		Message: types.PartialSignatureMessages{
			Type: msgType,
			Slot: slot,
		},
		Signer: signer,
	}
}
