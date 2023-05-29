package types

type ConsensusData struct {
	Duty    Duty
	DataSSZ []byte `ssz-max:"4194304"` // 2^22
}
