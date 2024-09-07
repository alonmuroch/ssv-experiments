package cluster

import (
	"fmt"
	"ssv-experiments/ssv_chain/operations"
	"ssv-experiments/ssv_chain/operations/gas"
	"ssv-experiments/ssv_chain/types"
)

type addClusterV0 struct {
	ModuleID uint64
	// Operators id's belonging to the cluster
	Operators []uint64 `ssz-max:"13"`
	// FaultyNodes represents the number of faulty nodes a cluster can sustain (fault tolerance)
	FaultyNodes uint64
	// Instances represents cluster instances for a specific module and account (e.g. staking module)
	Instances []*types.ClusterInstance `ssz-max:"500"`
}

type modifyClusterV0 struct {
	ClusterID uint64
	// Instances represents cluster instances for a specific module and account (e.g. staking module)
	InstancesToAdd    []*types.ClusterInstance `ssz-max:"500"`
	InstancesToRemove []*types.ClusterInstance `ssz-max:"500"`
}

func processV0Operation(ctx *operations.Context, op byte, raw []byte) error {
	switch op {
	// OP_Add is a user event for creating a new cluster
	case types.OP_Add:
		opObj := &addClusterV0{}
		if err := opObj.UnmarshalSSZ(raw); err != nil {
			return err
		}

		// calculate and consume gas
		estimatedGas := uint64(gas.ClusterAdd)
		estimatedGasCost := ctx.GasCost(estimatedGas)
		if err := gas.ConsumeGas(ctx, estimatedGasCost); err != nil {
			return err
		}
		ctx.GasConsumed += estimatedGas

		// Verify module exists, if not fail and consume gas
		if ctx.State.ModuleByID(opObj.ModuleID) == nil {
			return fmt.Errorf("module not found")
		}

		// verify operators exist
		for _, indx := range opObj.Operators {
			if ctx.State.Operators[indx] == nil {
				return fmt.Errorf("operator not found")
			}
		}

		cluster := &types.Cluster{
			ID:          uint64(len(ctx.State.Clusters)),
			ModuleID:    opObj.ModuleID,
			Address:     ctx.Account.Address,
			Operators:   opObj.Operators,
			FaultyNodes: opObj.FaultyNodes,
			Active:      true,
			Instances:   make([]*types.ClusterInstance, len(opObj.Instances)),
		}

		if err := addInstancesToCluster(ctx, cluster, opObj.Instances); err != nil {
			return err
		}

		// update cluster
		ctx.State.Clusters = append(ctx.State.Clusters, cluster)
		return nil
		// OP_Add is a user event for creating a new cluster
	case types.OP_Modify:
		opObj := &modifyClusterV0{}
		if err := opObj.UnmarshalSSZ(raw); err != nil {
			return err
		}

		// calculate and consume gas
		estimatedGas := uint64(gas.ClusterModify)
		estimatedGasCost := ctx.GasCost(estimatedGas)
		if err := gas.ConsumeGas(ctx, estimatedGasCost); err != nil {
			return err
		}
		ctx.GasConsumed += estimatedGas

		// Verify cluster exists, if not fail and consume gas
		cluster := ctx.State.ClusterByID(opObj.ClusterID)
		if cluster == nil {
			return fmt.Errorf("cluster not found")
		}

		// remove instances
		for _, inst := range opObj.InstancesToAdd {
			// calculate and consume gas
			estimatedGas = uint64(gas.ClusterInstanceRemove)
			estimatedGasCost = ctx.GasCost(estimatedGas)
			if err := gas.ConsumeGas(ctx, estimatedGasCost); err != nil {
				return err
			}
			ctx.GasConsumed += estimatedGas

			// remove
			removeClusterInstance(inst)
		}

		// add instances to cluster
		if err := addInstancesToCluster(ctx, cluster, opObj.InstancesToAdd); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unknown operation")
	}
}

func estimateAddClusterInstanceGas(ci *types.ClusterInstance) uint64 {
	return uint64(gas.ClusterInstanceAdd + len(ci.Metadata)*gas.ByteData)
}

// RemoveInstance removes cluster instance if found
func removeClusterInstance(instance *types.ClusterInstance) {
	panic("implement")
}

func addInstancesToCluster(ctx *operations.Context, cluster *types.Cluster, instances []*types.ClusterInstance) error {
	for i, inst := range instances {
		if len(inst.Keys) != len(cluster.Operators) {
			fmt.Errorf("invalid operator set")
		}

		// update gas
		estimatedGas := estimateAddClusterInstanceGas(inst)
		estimatedGasCost := ctx.GasCost(estimatedGas)
		if err := gas.ConsumeGas(ctx, estimatedGasCost); err != nil {
			return err
		}
		ctx.GasConsumed += estimatedGas

		// validate and add to operator price tiers
		for i, opIdx := range cluster.Operators {
			op := ctx.State.Operators[opIdx] // operator existence should be validated previously
			if uint64(len(op.Tiers)) < inst.PriceTierIndexes[i] {
				return fmt.Errorf("invalid price tier")
			}
			tier := op.Tiers[inst.PriceTierIndexes[i]]
			if !tier.CanRegister(ctx.Account.Address) {
				return fmt.Errorf("invalid register")
			}
			tier.Registered++
		}

		// add instance
		cluster.Instances[i] = inst
	}
	return nil
}
