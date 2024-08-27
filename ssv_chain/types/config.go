package types

// Configure is a struct holding all configurations for an ssv node
type Configure struct {
	SupportedNetworks        [][]byte `ssz-max:"12,4"`
	SSVTokenAddressByNetwork [][]byte `ssz-max:"12,128"`
}
