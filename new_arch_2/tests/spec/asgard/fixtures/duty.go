package fixtures

import "ssv-experiments/new_arch_2/spec/asgard/types"

const (
	Slot   = 123
	Height = 100
)

var AttesterDuty = &types.Duty{
	Role:        types.BeaconRoleAttester,
	ValidatorPK: ValidatorPK,
	Slot:        Slot,
}
