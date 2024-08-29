package types

import "bytes"

// Configure is a struct holding all configurations for an ssv node
type Configure struct {
	SupportedNetworks        [][]byte `ssz-max:"12,4"`
	SSVTokenAddressByNetwork [][]byte `ssz-max:"12,128"`
}

func (c *Configure) ValidSSVTokenAddress(network [4]byte, address []byte) bool {
	for i, n := range c.SupportedNetworks {
		if bytes.Equal(n, network[:]) {
			return bytes.Equal(c.SSVTokenAddressByNetwork[i], address)
		}
	}
	return false
}
