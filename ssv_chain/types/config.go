package types

import "bytes"

// Configure is a struct holding all configurations for a ssv node
type Configure struct {
	SupportedNetworks [][]byte `ssz-max:"12,4"`

	// Transactions
	// SystemTxSigner is the address from which system transactions are "signed" with
	SystemTxSigner []byte `ssz-max:"128"`

	// SSV token
	SSVTokenAddressByNetwork [][]byte `ssz-max:"12,128"`
	MainSSVTokenNetwork      [4]byte  `ssz-size:"4"`
	MainSSVTokenAddress      []byte   `ssz-max:"128"`

	// MissedValidationPenalty marks the amount of SSV to penalize a validator if it missed a block vote
	MissedValidationPenalty uint64
}

func (c *Configure) ValidSSVTokenAddress(network [4]byte, address []byte) bool {
	for i, n := range c.SupportedNetworks {
		if bytes.Equal(n[:], network[:]) {
			return bytes.Equal(c.SSVTokenAddressByNetwork[i], address)
		}
	}
	return false
}

func (c *Configure) IsSupportedNetwork(network [4]byte) bool {
	for _, n := range c.SupportedNetworks {
		if bytes.Equal(n[:], network[:]) {
			return true
		}
	}
	return false
}
