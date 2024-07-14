package dsp

type RunParams struct {
	Key   string
	Value string
}

type Fork struct {
	Version      [4]byte
	StartingSlot uint64
}

type Manifest struct {
	Address       [20]byte
	DockerFileURL string

	Fork *Fork

	// RunParams holds the various params needed when starting the DSP
	RunParams *RunParams

	// SlashingProtection holds wasm code for slashing protection(secure signer), called before every signing
	SlashingProtection []byte
}
