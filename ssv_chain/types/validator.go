package types

import "ssv-experiments/ssv_chain/common"

type Validator struct {
	// Account that controls the validator
	Account uint64
	// VotingPower represents the validator's voting power, per cometBFT.ValidatorUpdate struct
	VotingPower int64
	// PublicKey the validator uses to vote
	PublicKey *common.CryptoKey
}
