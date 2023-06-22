package qbft

type InputData struct {
	// Full data max value is ConsensusData max value ~= 2^22 + 2^16
	Data []byte `ssz-max:"4259840"`

	validateF func(data []byte) error
}

func (input *InputData) Validate() error {
	return input.validateF(input.Data)
}
