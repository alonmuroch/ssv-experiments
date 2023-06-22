package fixtures

import "ssv-experiments/new_arch/types"

const Slot = 123

var AttesterDuty = &types.Duty{
	Role:        types.BeaconRoleAttester,
	ValidatorPK: Validator,
	Slot:        123,
}
