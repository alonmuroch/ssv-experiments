package qbft

type State struct {
	Round  uint64
	Height uint64

	PreparedRound uint64
	PreparedValue *InputData

	// Messages is a unified (to all message type) container slice, simple and easy to serialize
	Messages Container `ssz-max:"256"`
}
