package fixtures

import "github.com/attestantio/go-eth2-client/spec/phase0"

var Validator = [48]byte{}

var AttestationData = &phase0.AttestationData{
	Slot:            12,
	Index:           3,
	BeaconBlockRoot: phase0.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	Source: &phase0.Checkpoint{
		Epoch: 0,
		Root:  phase0.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	},
	Target: &phase0.Checkpoint{
		Epoch: 1,
		Root:  phase0.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	},
}
var AttestationDataBytes = func() []byte {
	ret, _ := AttestationData.MarshalSSZ()
	return ret
}()
